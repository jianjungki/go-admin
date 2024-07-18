package payment

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	wechatPay "github.com/liyoung1992/wechatpay"
	"github.com/lunny/log"
	uuid "github.com/satori/go.uuid"
	qrcode "github.com/skip2/go-qrcode"
)

var cert = []byte(`-----BEGIN CERTIFICATE-----
MIIEATCCAumgAwIBAgIUMgsohLHRC5kgCm4NeMkgrht0nsQwDQYJKoZIhvcNAQEL
BQAwXjELMAkGA1UEBhMCQ04xEzARBgNVBAoTClRlbnBheS5jb20xHTAbBgNVBAsT
FFRlbnBheS5jb20gQ0EgQ2VudGVyMRswGQYDVQQDExJUZW5wYXkuY29tIFJvb3Qg
Q0EwHhcNMTgxMjE4MTM1MjU1WhcNMjMxMjE3MTM1MjU1WjCBkjETMBEGA1UEAwwK
MTUxNTk0MTk2MTEbMBkGA1UECgwS5b6u5L+h5ZWG5oi357O757ufMSowKAYDVQQL
DCHljJfkuqzkv6HlpKnpgq7np5HmioDmnInpmZDlhazlj7gxCzAJBgNVBAYMAkNO
MRIwEAYDVQQIDAlHdWFuZ0RvbmcxETAPBgNVBAcMCFNoZW5aaGVuMIIBIjANBgkq
hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAy9P9s2TenuwL3dv47t7XrKYNL1b9KqmV
fgkOqNe4R9ljgKvqccqmY16T6Ax+Z095i4d6mP7RMG4bBvbve0Q5WMmgGqVvRZcw
zdE9bySmpDZZoytEBgqtuOmINCnxRctoHzwWcP2BqYzWZ1bRZc4N3PMIbWHuapMh
zB72vLiWDHffLfHzy7bnWh4YzRn+OquX9RP4LLvfeCGzbXc0HdETlomGHASXcLrG
UyQ6u+iMpEcBmXUUcEoHPFBJlElQT2oXm8QCZifUv9FABHiXRVCehupDQaKha8h3
rpwo6n/9SF+3bc0PmXq1vEuqvRDe5fkSLZf660lfPKrs2hwflCwvLwIDAQABo4GB
MH8wCQYDVR0TBAIwADALBgNVHQ8EBAMCBPAwZQYDVR0fBF4wXDBaoFigVoZUaHR0
cDovL2V2Y2EuaXRydXMuY29tLmNuL3B1YmxpYy9pdHJ1c2NybD9DQT0xQkQ0MjIw
RTUwREJDMDRCMDZBRDM5NzU0OTg0NkMwMUMzRThFQkQyMA0GCSqGSIb3DQEBCwUA
A4IBAQCOxXblJRAfiFGrGXLewYh0BJHh8sUGBdgK8DDiPlp89jal0/8s6U5Ll8I8
L72fvW/qzwS+1unMK9WArGa0wf6UhefbXVqpckI+gowEPdWMXUaP9O/xn+9G4YWG
Xv57UhV5VluN232LLBDeTHaHn5Tpz8Y1vVFSx7dekbVbVNCpvz6U0EOBuOjjCnQb
lbD9+GXeHu+uNnBShFV6+SluxdS4uWEykxtpcHUYQgdGtzUwFNF6HPXoMPijEUQZ
JXPmI8drqBi3xFZl8AflYiu2901Gx0DOhWlavHA9wkMTD/Rxv6xfYR+WLF/CSluw
EPWSSHclt6DjDJJCYTslmskMVeEU
-----END CERTIFICATE-----`)

var key = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDL0/2zZN6e7Avd
2/ju3tespg0vVv0qqZV+CQ6o17hH2WOAq+pxyqZjXpPoDH5nT3mLh3qY/tEwbhsG
9u97RDlYyaAapW9FlzDN0T1vJKakNlmjK0QGCq246Yg0KfFFy2gfPBZw/YGpjNZn
VtFlzg3c8whtYe5qkyHMHva8uJYMd98t8fPLtudaHhjNGf46q5f1E/gsu994IbNt
dzQd0ROWiYYcBJdwusZTJDq76IykRwGZdRRwSgc8UEmUSVBPahebxAJmJ9S/0UAE
eJdFUJ6G6kNBoqFryHeunCjqf/1IX7dtzQ+ZerW8S6q9EN7l+RItl/rrSV88quza
HB+ULC8vAgMBAAECggEBALvLn+4PO8W0ueI5cAu1qYTaoT9CqJMMGD49XhJWXuVj
93dv7dvV2KOpWiDskfHhmZziwCQCcv5qE+DA1zj9TNDvQ+IJFpJJkJ+w7ydd32Ng
6AmcLVpbsF+0Wtek5TVnCZHtrMeB6lsq9cfUh62CwecWzGbrWd07nmIRDUgn0ae9
jIM6wUVOERrj9bBDdZ4hi9k3rcuvORj6LtmJHdwiFuN7uMvVR3gQ0jqOASv6jwKX
Mgfv5SCZR9SNOSi7nBqE2gghbSgXcCt2o6lkTQsl2kVGML+0xBTQz2+KffZH8v3y
P9JqDni4upc+sLnLpB1dGal+9P2biTFJ1O/4frB+rDECgYEA8WHH8HPicucop47t
+Ceq5T23OecIOgU3INtf/Q/tVEsiJ5KN8DrAQM1SX+hAn2ysqUcSN36XKn1zHmb8
zBE+VcuONWiL9cJP7tqaChh6xC1dqayK6aEVWH/zQoJF0LNyCmPwevlpwE3wA9E/
cXlOw2TqDNHIcgijYrj1Jgs+X1UCgYEA2CwANS83dUAjM/QD5DfKKLYtWzpK7na1
qYIs0q/4+9Jq+Mljk90fTcG3hON/HIBsNMwKKiEqElXz9l4lKQ4up5l4Q9BVycM7
TF02Qo8oeD0s5Lfk2eoaebtS8EHBvrt0DJS8pu3KlUIogTYhKLzMZhuqBwsg/w6o
GFZlf6gx7HMCgYBLSIvrWrlNzPrJB7MAqp6KUO1Midnu/qFrBgCyFL2YhLZqUTMj
sa80htUiRHWjQ2m2ggwUcs8C2Y2F9ejyXI5sg8gesDS2uZvkixlv/KKNfPixarcI
lszphT6Bh8/uf4kZ5IGkWRW6fJRXHHQFGCiQmHDG/sbdigQrTjRnYSpa4QKBgQCn
Jz7x4RTNAU83ORRZQBp96ICh7i1Kv1gxPotSTugEMOyue2sYRv+RrJ0vkIoVi+gf
2zZw5TwW+YevlZS+bkz9I9qu9UWgwen0Xo86YGA2uXchVGENb8wRHRxA01fxGBng
MJK/sVfYmlOfti9lxLd4rSOIVr6o+rhYOVY6o+AWqQKBgQCcHMvess2fBK7Nt65s
EqsfSSz7t/b73YIRz+8BjjIIIl1UbXNy4zFn2DbSgydpgfu2Gtb7rnyKBgWLerRW
dW4iwxG/XxEikbHBQI289S/6BJHqEq5liKSdvhfh75lEHHnTwq3saD5Yky8JvUUx
9lxxxnnPnj+hXJN/OomqOIr/dw==
-----END PRIVATE KEY-----`)

var WechatAPPID = "wx23aca6b7e1239250"
var WechatMCHID = "1515941961"
var WechatAPIKEY = "rgn6MXP6uMLEBMuw4tFusEnjgLNHm2mU"
var wechatClient *wechatPay.WechatPay

func init() {
	wechatCert := cert
	wechatKey := key
	wechatClient = wechatPay.New(WechatAPPID, WechatMCHID, WechatAPIKEY, wechatCert, wechatKey)

}

func WechatpayPay(c gin.Context) error {
	//fmt.Printf("job %v\n", c.Get("user"))
	outTradeNo := utils.GetMD5(uuid.NewV4().String())

	payObj, err := CreateOrder(c, outTradeNo)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	var payData wechatPay.UnitOrder

	payData.NotifyUrl = "http://wefile.com/wechat/callback"
	payData.TradeType = "NATIVE"
	payData.Body = "充值费用"
	payData.SpbillCreateIp = c.RealIP()
	payData.TotalFee = payObj.PayFee
	payData.OutTradeNo = payObj.OutTradeNo
	result, _ := wechatClient.Pay(payData)

	//if result.ResultCode == "FAIL" || result.ReturnCode == "FAIL" {
	//	return c.JSON(http.StatusOK, "fail")
	//}
	//timeStamp := strconv.Itoa(int(time.Now().Unix()))
	//qrcodeStr := `weixin：//wxpay/bizpayurl?sign=` + result.Sign + `&appid=` + result.AppId + `&mch_id=` + result.MchId + `&product_id=123&time_stamp=` + timeStamp + `&nonce_str=` + result.NonceStr
	png, err := qrcode.Encode(result.CodeUrl, qrcode.Medium, 256)
	//err = qrcode.WriteFile(result.CodeUrl, qrcode.Medium, 256, "qrcode.png")
	//prefix := []byte("data:image/png;base64,")
	retData := map[string]interface{}{
		"order_id": payObj.Id,
		"qrcode":   png,
	}
	return utils.CommJSONRet(c, utils.Success, retData)
}

func WechatpayRefund(c gin.Context) error {
	payObj := GetOrder(c)

	var refundData wechatPay.OrderRefund
	refundData.TotalFee = payObj.PayFee
	refundData.OutTradeNo = payObj.OutTradeNo
	refundData.OutRefundNo = payObj.OutTradeNo
	refundData.RefundFee = payObj.PayFee

	fmt.Printf("%v\n", refundData)
	result, err := wechatClient.Refund(refundData)
	if err != nil {
		return utils.CommJSONRet(c, utils.PayCreateError, err.Error())
	}
	if result.ResultCode == "SUCCESS" {
		return utils.CommJSONRet(c, utils.Success, result)
	}

	return utils.CommJSONRet(c, utils.PayCreateError, result.ReturnMsg)
}

func WechatpayRefundQuery(c gin.Context) error {
	var queryData wechatPay.OrderRefundQuery
	queryData.OutTradeNo = "111122"

	result, _ := wechatClient.RefundQuery(queryData)
	return c.JSON(http.StatusOK, result)
}

func WechatCallBack(c gin.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Error(err, "read notify body error")
		return c.JSON(http.StatusOK, "fail")
	}
	retParams := string(body)
	fmt.Println("body:", retParams)

	var wxNotifyReq wechatPay.PayNotifyResult
	err = xml.Unmarshal(body, &wxNotifyReq)
	if err != nil {
		log.Error(err, "read http body xml failed! err :"+err.Error())
		return c.String(http.StatusOK, "fail")
	}

	if wxNotifyReq.ResultCode == "SUCCESS" {
		ret := UpdateOrderStatus(PAYED, retParams, wxNotifyReq.OutTradeNo)
		if !ret {
			fmt.Println("update order status error, outTradeNo: ", wxNotifyReq.OutTradeNo)
		}
	} else {
		UpdateOrderStatus(FAILD, retParams, wxNotifyReq.OutTradeNo)
	}
	return c.String(http.StatusOK, "success")

}

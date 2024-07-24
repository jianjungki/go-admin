package payment

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/smartwalle/alipay"

	"bytes"
)

var publickey = `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhVS+S0q4TjuyGZnMrSgrOYZDev+92JH/pTtaRSEVD3lpowisP/eUsuwDMQBXXxjX4mUMQE0XrcQ6jZHdBbzxYD7svQAJ1fnZWLMAtnYr9DV1mjCz7nsWRE6P1KjrNPSJOVMA8M3u6gdKdTcgWMidz/VLzhjDWmLI3Aq3oMysKoESHzpPkKjFOIgpg9Q7WAWdbmawAsQbsa0DZ2Ph2ilIN05GD92AJFEGVXa+EH/wtP+0cUz46PS6cd56vs5vpngjgkmVfuC+raCLt/ShZchC5T2hMyuoj7eDmskVg+qPYdVapNw/WN+YeMS9jmXino/7R5CG6FibdTpbyPVgSFGnCQIDAQAB`
var privatekey = `MIIEpAIBAAKCAQEAuAjdAbELz+Hfubh8Vg16pD8fyedGKyGFGj/mX+JS3i6kIE1EO1oRwGAn5G/A2BT5XM+8hqMMjY9uSZCsLbXr49Ru7rQjLwkj/EK4WG+6YCMO6OwkkkvxdHGh5GtNPgCLhv3SUZNMyH2poRlJQA6DN2iCueGLaMulqAdoFEmziFxRyGMooafWp4vkZT/ZfnMaF0ceamjBw89lO6x7HYyX6qFdH3S0+rMp1LI2rQCxgnNJkSE4NJ96JIfO8NeEiBbtmuEh9cL7aw4vQEth8mxePT74GM7uvmf9V0ZUlo09GuicrMP84eA44TqWZshhBL8xyKAfi3vYMIyvwsVF1mTmGwIDAQABAoIBAFowFca/ce91XQOR8nTldCXLvZNfoJaAfc8oY4i3zC6QIuSlmPuRH0Psm4AqeBIs7StbrqM4gm1ZTg4FNSvNjIWBVL/sTZW+YAE+3UKRtn9ojuT83MtJJlIPYewj3Z21/Kqx9wD9M9LJKnkFgE1BD3L256Gac4xuyIb44juXBsGaz6cU1o+nMkKg3DAMDRCGojGI1W7XOk7xtY9HYFNZz9gbCMBYPasCGEU+I1FDIdCv2vuG7GhjsREaQ2pwgRQUOJHS4ZwbSjtf2X2/NJbMe2cn8B+k8DTwr3kvhRs9wiWW3ojjBF8hzVDaeFFEJwDjFtWeu0qLcAMC/VkYX/PjPBECgYEA4KV4mFAnTtO70/pTn7LhuQs7s21qf/L+/1qH0d8eoruP9IN5Rn3QYjDtLM8QxqftX6WGL5r3TllUnw6b9nL0ja6VrSd+WCTj73tQHzM04yCOO8VuEIBc7AwGvU27wF/BlXKbAF0EqyCcdSkR2SRsDb62k5K0IxvU3NtznCkdnOMCgYEA0bhZl7gpIX2vyqZFmBgCxEp0yPF+c4EvX+rYIMr+ll8FionI5XFixAKitB8KH/0gkSNMeS2YYI3zAnG5gaCfM3wybqsT3IYAwwln3nMaOevbskMHmIGc8FwA0XEUhupAir2D8qA+fg11XZMX1DBeDBCQPfNq+log1kZLQPxMz2kCgYBhjRmIMs5O+b0HuR2VLM2+WDVtnibwNuoqovp11YjtK7vAd1MznDDgYtSrb6OeUI/QyAU0KPwZhssrYuGxLIwr9lNC5PKFfkeI2Cib09W0h4+cPwijm8o+Kp7Gl7nfOPXMWv8aMkpzsW6fpdz9SUau1VFdVghpco2F9roGZwhX9wKBgQCd8MQoat0vB8T4TThkTtDmm1xQaO7edhADoI2cRMRu4qj9qL7PbREApSt7k1TAupVoOLlDGX+EV+MHmDYo1ZqP35zrk5OlpJW2TMMY2H6L2IDgQ3LY++QBcIh9MzhOUuTAl/FjLCNkMyaMjlaHytVJKMA4p7WNTYslVFMgYeJG4QKBgQDIS/RU6ruMUXj+TcyBDX39/BquTyr6gmvFJumexc0ZtpCKMO9DXJxa10vdBvnMZ2mgdiuaqhs3kFmq+SWURvTT6a4P3c5BvNFNiEhq9CyXEWxHh5Rh47cuGUDsLbQ+P2cRn4MzQeXyHeFaiPSatwOekcnVxs+LUbLCScRPaDsiYg==`

var aliClient *alipay.AliPay

func init() {
	var appId = "2018092861536722"
	//var err error
	//var partnerId = "2088821267063894"
	//publicKey := publickey
	//privateKey := privatekey
	aliClient = alipay.New(appId, publickey, privatekey, true)
	/*
		if  err!= nil{
			log.Info(err, "alipay create error")
		}
	*/
}

// AlipayPay godoc
// @Summary 支付宝支付
// @Description 支付宝支付
// @Tags api-pay
// @ID api-alipay-pay
// @Accept  json
// @Produce  json
// @Param  pay_params body pay.UserPayOrder true "支付对象"
// @Success 200 {object} pay.AliPayResp
// @Router /alipay/pay [post]
func AlipayPay(c *gin.Context) error {
	outTradeNo := utils.GetMD5(uuid.NewV4().String())
	payObj, err := CreateOrder(c, outTradeNo)
	if payObj.PayFee <= 0 {
		return utils.CommJSONRet(c, utils.PayCreateError, "支付金额必须大于1角")
	}

	if err != nil {
		return utils.CommJSONRet(c, utils.PayCreateError, err.Error())
	}

	var p = alipay.AliPayTradePagePay{}

	var PayAmount float64
	PayAmount = float64(payObj.PayFee) / 100

	p.NotifyURL = "http://wefile.com/alipay/callback"
	p.Subject = "Wefile转换费用"
	p.OutTradeNo = payObj.OutTradeNo
	p.TotalAmount = fmt.Sprintf("%v", PayAmount)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	html, err := aliClient.TradePagePay(p)
	if err != nil {
		return utils.CommJSONRet(c, utils.Success, map[string]interface{}{"error": err.Error()})
	}
	retData := AliPayResp{
		OrderID: payObj.Id,
		URL:     html.String(),
	}
	return utils.CommJSONRet(c, utils.Success, retData)
}

// AlipayRefund godoc
// @Summary 支付宝退款
// @Description 支付宝退款
// @Tags api-pay
// @ID api-alipay-refund
// @Accept  json
// @Produce  json
// @Param order_id path int true "订单id"
// @Success 200 {object} pay.AliPayResp
// @Router /alipay/refund [post]
func AlipayRefund(c *gin.Context) error {
	payObj := GetOrder(c)

	var r = alipay.AliPayTradeRefund{}
	r.OutTradeNo = payObj.OutTradeNo
	r.RefundReason = "用户退款"

	var PayAmount float64
	PayAmount = float64(payObj.PayFee) / 100
	r.RefundAmount = strconv.FormatFloat(PayAmount, 'f', -1, 32)
	//fmt.Printf("refundObj: %s", r.RefundAmount)
	var res, _ = aliClient.TradeRefund(r)
	if res.AliPayTradeRefund.FundChange == "Y" {
		retData := map[string]interface{}{
			"message": "退款成功",
		}
		return utils.CommJSONRet(c, utils.Success, retData)
	}

	return utils.CommJSONRet(c, utils.Success, map[string]interface{}{
		"message": "退款失败",
	})

}

func AlipayRefundQuery(c *gin.Context) error {
	var q = alipay.AliPayFastpayTradeRefundQuery{}

	var res, _ = aliClient.TradeFastpayRefundQuery(q)
	return c.JSON(http.StatusOK, res)
}

// AlipayCallBack godoc
// @Summary 支付宝支付回调
// @Description 支付宝支付回调
// @Tags api-pay
// @ID api-alipay-callback
// @Accept  json
// @Produce  json
// @Success 200 {string} RetStatus ""
// @Router /alipay/callback [post]
func AlipayCallBack(c *gin.Context) error {
	var bodyBytes []byte
	var err error
	if c.Request().Body != nil {
		bodyBytes, err = io.ReadAll(c.Request().Body)
		if err != nil {
			log.Error(err, "read notify body error")
			return c.String(http.StatusOK, "fail")
		}
	}
	retParams := string(bodyBytes)
	fmt.Println("body:", retParams)

	// 把刚刚读出来的再写进去
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	fmt.Println("param:", c.Request().PostFormValue("notify_id"))

	noti, err := aliClient.GetTradeNotification(c.Request())
	if noti != nil && noti.TradeStatus == alipay.K_TRADE_STATUS_TRADE_SUCCESS {
		UpdateOrderStatus(PAYED, retParams, noti.OutTradeNo)
		fmt.Println("支付成功")
	} else {
		//UpdateOrderStatus(FAILD, retParams, noti.OutTradeNo)
		//if err != nil {
		fmt.Printf("错误信息 %v \n", err)
		//}
		fmt.Printf("支付失败 %v\n", noti)
	}
	return c.String(http.StatusOK, "success")
}

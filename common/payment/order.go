package payment

import (
	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"github.com/go-admin-team/go-admin-core/sdk/service"

	"go-admin/app/admin/models"
)

type OrderSrv struct {
	service.Service
}

// CreateOrder 创建订单基础函数
func (e *OrderSrv) CreateOrder(c *gin.Context, outTradeNo string) (*models.UserPayOrder, error) {
	userId := user.GetUserId(c)

	order := new(models.UserPayOrder)
	if err := c.Bind(order); err != nil {
		log.Fatalf("bind error: %v", err)
		return order, err
	}
	order.OutTradeNo = outTradeNo
	order.Status = models.WAITING
	order.UserId = userId
	err := e.Orm.Create(&order)
	if err != nil {
		return order, err
	}

	return order, nil
}

// UpdateOrderStatus 更新订单状态
func (e *OrderSrv) UpdateOrderStatus(status UserPayStatus, retParams string, outTradeNo string) bool {
	order := new(models.UserPayOrder)
	order.Status = status
	order.RetParams = retParams
	db := e.Orm.Debug().First(&model, c.GetId())

	affected, err := tx.Where("out_trade_no = ?", outTradeNo).Update(order)
	if affected > 0 && err == nil {
		return true
	}

	return false
}

// GetOrder 获取订单信息
func (e *OrderSrv) GetOrder(c *gin.Context) UserPayOrder {
	tx := e.Orm.Debug().Begin()
	info := utils.GetUser(c)

	orderID := c.Param("order_id")

	order := models.UserPayOrder{}
	engine.Where("id = ? and user_id = ?", orderID, info.Uid).Get(&order)
	return order
}

// GetOrderByOutTrade 获取payobj通过out_trade_no
func (e *OrderSrv) GetOrderByOutTrade(outTradeNo string) models.UserPayOrder {
	order := models.UserPayOrder{}
	engine.Where("out_trade_no = ?", outTradeNo).Get(&order)
	return order
}

// GetOrderInfo godoc
// @Summary 获取订单信息
// @Description 获取订单信息
// @Tags api-job
// @ID api-job-info
// @Accept  json
// @Produce  json
// @Param order_id path int true "订单id"
// @Success 200 {object} pay.UserPayOrder
// @Router /order/info [get]
func (e *OrderSrv) GetOrderInfo(c *gin.Context) error {
	payObj := GetOrder(c)
	return utils.CommJSONRet(c, utils.Success, payObj)
}

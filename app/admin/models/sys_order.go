package models

import (
	"go-admin/common/models"
)

type SysOrder struct {
	models.Model

	PayParams      string `json:"payParams" gorm:"type:varchar(255);comment:PayParams"`
	RetParams      string `json:"retParams" gorm:"type:varchar(255);comment:RetParams"`
	PayFee         string `json:"payFee" gorm:"type:int(11);comment:PayFee"`
	Status         int    `json:"status" gorm:"type:int(11);comment:Status"`
	RefundFee      string `json:"refundFee" gorm:"type:int(11);comment:RefundFee"`
	UserId         int    `json:"userId" gorm:"type:bigint(20);comment:UserId"`
	OutTradeNo     string `json:"outTradeNo" gorm:"type:varchar(255);comment:OutTradeNo"`
	RefundOutTrade string `json:"refundOutTrade" gorm:"type:varchar(255);comment:RefundOutTrade"`
	models.ModelTime
	models.ControlBy
}

func (SysOrder) TableName() string {
	return "sys_order"
}

func (e *SysOrder) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysOrder) GetId() interface{} {
	return e.Id
}

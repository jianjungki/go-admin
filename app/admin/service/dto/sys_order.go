package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysOrderGetPageReq struct {
	dto.Pagination `search:"-"`
	SysOrderOrder
}

type SysOrderOrder struct {
	Id             int    `form:"idOrder"  search:"type:order;column:id;table:sys_order"`
	PayParams      string `form:"payParamsOrder"  search:"type:order;column:pay_params;table:sys_order"`
	RetParams      string `form:"retParamsOrder"  search:"type:order;column:ret_params;table:sys_order"`
	PayFee         string `form:"payFeeOrder"  search:"type:order;column:pay_fee;table:sys_order"`
	Status         int    `form:"statusOrder"  search:"type:order;column:status;table:sys_order"`
	RefundFee      string `form:"refundFeeOrder"  search:"type:order;column:refund_fee;table:sys_order"`
	UserId         int    `form:"userIdOrder"  search:"type:order;column:user_id;table:sys_order"`
	OutTradeNo     string `form:"outTradeNoOrder"  search:"type:order;column:out_trade_no;table:sys_order"`
	RefundOutTrade string `form:"refundOutTradeOrder"  search:"type:order;column:refund_out_trade;table:sys_order"`
}

func (m *SysOrderGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysOrderInsertReq struct {
	Id             int    `json:"-" comment:""` //
	PayParams      string `json:"payParams" comment:""`
	RetParams      string `json:"retParams" comment:""`
	PayFee         string `json:"payFee" comment:""`
	Status         int    `json:"status" comment:""`
	RefundFee      string `json:"refundFee" comment:""`
	UserId         int    `json:"userId" comment:""`
	OutTradeNo     string `json:"outTradeNo" comment:""`
	RefundOutTrade string `json:"refundOutTrade" comment:""`
	common.ControlBy
}

func (s *SysOrderInsertReq) Generate(model *models.SysOrder) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.PayParams = s.PayParams
	model.RetParams = s.RetParams
	model.PayFee = s.PayFee
	model.Status = s.Status
	model.RefundFee = s.RefundFee
	model.UserId = s.UserId
	model.OutTradeNo = s.OutTradeNo
	model.RefundOutTrade = s.RefundOutTrade
}

func (s *SysOrderInsertReq) GetId() interface{} {
	return s.Id
}

type SysOrderUpdateReq struct {
	Id             int    `uri:"id" comment:""` //
	PayParams      string `json:"payParams" comment:""`
	RetParams      string `json:"retParams" comment:""`
	PayFee         string `json:"payFee" comment:""`
	Status         int    `json:"status" comment:""`
	RefundFee      string `json:"refundFee" comment:""`
	UserId         int    `json:"userId" comment:""`
	OutTradeNo     string `json:"outTradeNo" comment:""`
	RefundOutTrade string `json:"refundOutTrade" comment:""`
	common.ControlBy
}

func (s *SysOrderUpdateReq) Generate(model *models.SysOrder) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.PayParams = s.PayParams
	model.RetParams = s.RetParams
	model.PayFee = s.PayFee
	model.Status = s.Status
	model.RefundFee = s.RefundFee
	model.UserId = s.UserId
	model.OutTradeNo = s.OutTradeNo
	model.RefundOutTrade = s.RefundOutTrade
}

func (s *SysOrderUpdateReq) GetId() interface{} {
	return s.Id
}

// SysOrderGetReq 功能获取请求参数
type SysOrderGetReq struct {
	Id int `uri:"id"`
}

func (s *SysOrderGetReq) GetId() interface{} {
	return s.Id
}

// SysOrderDeleteReq 功能删除请求参数
type SysOrderDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysOrderDeleteReq) GetId() interface{} {
	return s.Ids
}

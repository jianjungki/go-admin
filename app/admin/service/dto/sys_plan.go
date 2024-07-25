package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysPlanGetPageReq struct {
	dto.Pagination `search:"-"`
	SysPlanOrder
}

type SysPlanOrder struct {
	PlanId       string `form:"planIdOrder"  search:"type:order;column:plan_id;table:sys_plan"`
	Quota        string `form:"quotaOrder"  search:"type:order;column:quota;table:sys_plan"`
	Price        string `form:"priceOrder"  search:"type:order;column:price;table:sys_plan"`
	BillingCycle string `form:"billingCycleOrder"  search:"type:order;column:billing_cycle;table:sys_plan"`
	Remark       string `form:"remarkOrder"  search:"type:order;column:remark;table:sys_plan"`
	Name         string `form:"nameOrder"  search:"type:order;column:name;table:sys_plan"`
	Desc         string `form:"descOrder"  search:"type:order;column:desc;table:sys_plan"`
}

func (m *SysPlanGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysPlanInsertReq struct {
	PlanId       int    `json:"-" comment:""` //
	Quota        string `json:"quota" comment:""`
	Price        string `json:"price" comment:""`
	BillingCycle string `json:"billingCycle" comment:""`
	Remark       string `json:"remark" comment:""`
	Name         string `json:"name" comment:""`
	Desc         string `json:"desc" comment:""`
	common.ControlBy
}

func (s *SysPlanInsertReq) Generate(model *models.SysPlan) {
	if s.PlanId == 0 {
		model.Model = common.Model{Id: s.PlanId}
	}
	model.Quota = s.Quota
	model.Price = s.Price
	model.BillingCycle = s.BillingCycle
	model.Remark = s.Remark
	model.Name = s.Name
	model.Desc = s.Desc
}

func (s *SysPlanInsertReq) GetId() interface{} {
	return s.PlanId
}

type SysPlanUpdateReq struct {
	PlanId       int    `uri:"id" comment:""` //
	Quota        string `json:"quota" comment:""`
	Price        string `json:"price" comment:""`
	BillingCycle string `json:"billingCycle" comment:""`
	Remark       string `json:"remark" comment:""`
	Name         string `json:"name" comment:""`
	Desc         string `json:"desc" comment:""`
	common.ControlBy
}

func (s *SysPlanUpdateReq) Generate(model *models.SysPlan) {
	if s.PlanId == 0 {
		model.Model = common.Model{Id: s.PlanId}
	}
	model.Quota = s.Quota
	model.Price = s.Price
	model.BillingCycle = s.BillingCycle
	model.Remark = s.Remark
	model.Name = s.Name
	model.Desc = s.Desc
}

func (s *SysPlanUpdateReq) GetId() interface{} {
	return s.PlanId
}

// SysPlanGetReq 功能获取请求参数
type SysPlanGetReq struct {
	PlanId int `uri:"id"`
}

func (s *SysPlanGetReq) GetId() interface{} {
	return s.PlanId
}

// SysPlanDeleteReq 功能删除请求参数
type SysPlanDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysPlanDeleteReq) GetId() interface{} {
	return s.Ids
}

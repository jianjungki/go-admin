package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysUserPlanGetPageReq struct {
	dto.Pagination `search:"-"`
	SysUserPlanOrder
}

type SysUserPlanOrder struct {
	Id       int    `form:"idOrder"  search:"type:order;column:id;table:sys_user_plan"`
	Quota    int    `form:"quotaOrder"  search:"type:order;column:quota;table:sys_user_plan"`
	Used     int    `form:"usedOrder"  search:"type:order;column:used;table:sys_user_plan"`
	Leave    int    `form:"leaveOrder"  search:"type:order;column:leave;table:sys_user_plan"`
	UserId   int    `form:"userIdOrder"  search:"type:order;column:user_id;table:sys_user_plan"`
	PlanId   int    `form:"planIdOrder"  search:"type:order;column:plan_id;table:sys_user_plan"`
	StartAt  string `form:"startAtOrder"  search:"type:order;column:start_at;table:sys_user_plan"`
	UpdateAt string `form:"updateAtOrder"  search:"type:order;column:update_at;table:sys_user_plan"`
	CreateAt string `form:"createAtOrder"  search:"type:order;column:create_at;table:sys_user_plan"`
	EndAt    string `form:"endAtOrder"  search:"type:order;column:end_at;table:sys_user_plan"`
	DeleteAt string `form:"deleteAtOrder"  search:"type:order;column:delete_at;table:sys_user_plan"`
}

func (m *SysUserPlanGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysUserPlanInsertReq struct {
	Id      int       `json:"id" comment:""`
	Quota   int       `json:"quota" comment:""`
	Used    int       `json:"used" comment:""`
	Leave   int       `json:"leave" comment:""`
	UserId  int       `json:"userId" comment:""`
	PlanId  int       `json:"planId" comment:""`
	StartAt time.Time `json:"startAt" comment:""`
	EndAt   time.Time `json:"endAt" comment:""`
	common.ControlBy
}

func (s *SysUserPlanInsertReq) Generate(model *models.SysUserPlan) {
	model.Id = s.Id
	model.Quota = s.Quota
	model.Used = s.Used
	model.Leave = s.Leave
	model.UserId = s.UserId
	model.PlanId = s.PlanId
	model.StartAt = s.StartAt
	model.EndAt = s.EndAt
}

func (s *SysUserPlanInsertReq) GetId() interface{} {
	return s.Id
}

type SysUserPlanUpdateReq struct {
	Id       int       `json:"id" comment:""`
	Quota    int       `json:"quota" comment:""`
	Used     int       `json:"used" comment:""`
	Leave    int       `json:"leave" comment:""`
	UserId   int       `json:"userId" comment:""`
	PlanId   int       `json:"planId" comment:""`
	StartAt  time.Time `json:"startAt" comment:""`
	UpdateAt time.Time `json:"updateAt" comment:""`
	CreateAt time.Time `json:"createAt" comment:""`
	EndAt    time.Time `json:"endAt" comment:""`
	DeleteAt time.Time `json:"deleteAt" comment:""`
	common.ControlBy
}

func (s *SysUserPlanUpdateReq) Generate(model *models.SysUserPlan) {
	model.Id = s.Id
	model.Quota = s.Quota
	model.Used = s.Used
	model.Leave = s.Leave
	model.UserId = s.UserId
	model.PlanId = s.PlanId
	model.StartAt = s.StartAt
	model.UpdateAt = s.UpdateAt
	model.CreateAt = s.CreateAt
	model.EndAt = s.EndAt
	model.DeleteAt = s.DeleteAt
}

func (s *SysUserPlanUpdateReq) GetId() interface{} {
	return s.Id
}

// SysUserPlanGetReq 功能获取请求参数
type SysUserPlanGetReq struct {
}

func (s *SysUserPlanGetReq) GetId() interface{} {
	return s.GetId()
}

// SysUserPlanDeleteReq 功能删除请求参数
type SysUserPlanDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysUserPlanDeleteReq) GetId() interface{} {
	return s.Ids
}

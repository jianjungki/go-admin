package models

import (
	"time"

	"go-admin/common/models"
)

type SysUserPlan struct {
	models.Model

	Id       int       `json:"id" gorm:"type:int(11);comment:Id"`
	Quota    int       `json:"quota" gorm:"type:int(11);comment:Quota"`
	Used     int       `json:"used" gorm:"type:int(11);comment:Used"`
	Leave    int       `json:"leave" gorm:"type:int(11);comment:Leave"`
	UserId   int       `json:"userId" gorm:"type:int(11);comment:UserId"`
	PlanId   int       `json:"planId" gorm:"type:int(11);comment:PlanId"`
	StartAt  time.Time `json:"startAt" gorm:"type:datetime;comment:StartAt"`
	UpdateAt time.Time `json:"updateAt" gorm:"type:datetime;comment:UpdateAt"`
	CreateAt time.Time `json:"createAt" gorm:"type:datetime;comment:CreateAt"`
	EndAt    time.Time `json:"endAt" gorm:"type:datetime;comment:EndAt"`
	DeleteAt time.Time `json:"deleteAt" gorm:"type:datetime;comment:DeleteAt"`
	models.ModelTime
	models.ControlBy
}

func (SysUserPlan) TableName() string {
	return "sys_user_plan"
}

func (e *SysUserPlan) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysUserPlan) GetId() interface{} {
	return e.Id
}

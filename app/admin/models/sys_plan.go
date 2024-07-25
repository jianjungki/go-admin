package models

import (
	"go-admin/common/models"
)

type SysPlan struct {
	models.Model

	Quota        string `json:"quota" gorm:"type:int(11);comment:Quota"`
	Price        string `json:"price" gorm:"type:int(11);comment:Price"`
	BillingCycle string `json:"billingCycle" gorm:"type:varchar(50);comment:BillingCycle"`
	Remark       string `json:"remark" gorm:"type:varchar(50);comment:Remark"`
	Name         string `json:"name" gorm:"type:varchar(50);comment:Name"`
	Desc         string `json:"desc" gorm:"type:text;comment:Desc"`
	models.ModelTime
	models.ControlBy
}

func (SysPlan) TableName() string {
	return "sys_plan"
}

func (e *SysPlan) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysPlan) GetId() interface{} {
	return e.Id
}

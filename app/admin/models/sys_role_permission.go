package models

import (

	"go-admin/common/models"

)

type SysRolePermission struct {
    models.Model
    
    Permission string `json:"permission" gorm:"type:varchar(50);comment:Permission"` 
    Method string `json:"method" gorm:"type:varchar(50);comment:Method"` 
    Path string `json:"path" gorm:"type:varchar(50);comment:Path"` 
    RoleId string `json:"roleId" gorm:"type:int(11);comment:RoleId"` 
    models.ModelTime
    models.ControlBy
}

func (SysRolePermission) TableName() string {
    return "sys_role_permission"
}

func (e *SysRolePermission) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysRolePermission) GetId() interface{} {
	return e.Id
}
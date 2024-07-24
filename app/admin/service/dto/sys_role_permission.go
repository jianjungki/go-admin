package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysRolePermissionGetPageReq struct {
	dto.Pagination `search:"-"`
	SysRolePermissionOrder
}

type SysRolePermissionOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:sys_role_permission"`
	Permission string `form:"permissionOrder"  search:"type:order;column:permission;table:sys_role_permission"`
	Method     string `form:"methodOrder"  search:"type:order;column:method;table:sys_role_permission"`
	Path       string `form:"pathOrder"  search:"type:order;column:path;table:sys_role_permission"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_role_permission"`
	UpdatedAt  string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_role_permission"`
	DeletedAt  string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sys_role_permission"`
	RoleId     string `form:"roleIdOrder"  search:"type:order;column:role_id;table:sys_role_permission"`
}

func (m *SysRolePermissionGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysRolePermissionInsertReq struct {
	Id         int    `json:"-" comment:""` //
	Permission string `json:"permission" comment:""`
	Method     string `json:"method" comment:""`
	Path       string `json:"path" comment:""`
	RoleId     int    `json:"roleId" comment:""`
	common.ControlBy
}

func (s *SysRolePermissionInsertReq) Generate(model *models.SysRolePermission) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Permission = s.Permission
	model.Method = s.Method
	model.Path = s.Path
	model.RoleId = s.RoleId
}

func (s *SysRolePermissionInsertReq) GetId() interface{} {
	return s.Id
}

type SysRolePermissionUpdateReq struct {
	Id         int    `uri:"id" comment:""` //
	Permission string `json:"permission" comment:""`
	Method     string `json:"method" comment:""`
	Path       string `json:"path" comment:""`
	RoleId     int    `json:"roleId" comment:""`
	common.ControlBy
}

func (s *SysRolePermissionUpdateReq) Generate(model *models.SysRolePermission) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Permission = s.Permission
	model.Method = s.Method
	model.Path = s.Path
	model.RoleId = s.RoleId
}

func (s *SysRolePermissionUpdateReq) GetId() interface{} {
	return s.Id
}

// SysRolePermissionGetReq 功能获取请求参数
type SysRolePermissionGetReq struct {
	Id int `uri:"id"`
}

func (s *SysRolePermissionGetReq) GetId() interface{} {
	return s.Id
}

// SysRolePermissionDeleteReq 功能删除请求参数
type SysRolePermissionDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysRolePermissionDeleteReq) GetId() interface{} {
	return s.Ids
}

type SysRolePermissionCheckReq struct {
	AuthMethod string `json:"resource_method"`
	AuthPath   string `json:"resource_path"`
}

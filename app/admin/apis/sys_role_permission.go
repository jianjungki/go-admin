package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"github.com/lunny/log"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SysRolePermission struct {
	api.Api
}

// GetPage 获取角色接口权限列表
// @Summary 获取角色接口权限列表
// @Description 获取角色接口权限列表
// @Tags 角色接口权限
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysRolePermission}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-role-permission [get]
// @Security Bearer
func (e SysRolePermission) GetPage(c *gin.Context) {
	req := dto.SysRolePermissionGetPageReq{}
	s := service.SysRolePermission{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysRolePermission, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取角色接口权限失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取角色接口权限
// @Summary 获取角色接口权限
// @Description 获取角色接口权限
// @Tags 角色接口权限
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysRolePermission} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-role-permission/{id} [get]
// @Security Bearer
func (e SysRolePermission) Get(c *gin.Context) {
	req := dto.SysRolePermissionGetReq{}
	s := service.SysRolePermission{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysRolePermission

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取角色接口权限失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建角色接口权限
// @Summary 创建角色接口权限
// @Description 创建角色接口权限
// @Tags 角色接口权限
// @Accept application/json
// @Product application/json
// @Param data body dto.SysRolePermissionInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-role-permission [post]
// @Security Bearer
func (e SysRolePermission) Insert(c *gin.Context) {
	req := dto.SysRolePermissionInsertReq{}
	s := service.SysRolePermission{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建角色接口权限失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改角色接口权限
// @Summary 修改角色接口权限
// @Description 修改角色接口权限
// @Tags 角色接口权限
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysRolePermissionUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-role-permission/{id} [put]
// @Security Bearer
func (e SysRolePermission) Update(c *gin.Context) {
	req := dto.SysRolePermissionUpdateReq{}
	s := service.SysRolePermission{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改角色接口权限失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除角色接口权限
// @Summary 删除角色接口权限
// @Description 删除角色接口权限
// @Tags 角色接口权限
// @Param data body dto.SysRolePermissionDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-role-permission [delete]
// @Security Bearer
func (e SysRolePermission) Delete(c *gin.Context) {
	s := service.SysRolePermission{}
	req := dto.SysRolePermissionDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除角色接口权限失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// CheckPermission
// @Summary 获取个人信息
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-role-permission/check [post]
// @Security Bearer
func (e SysRolePermission) CheckPermission(c *gin.Context) {
	req := dto.SysRolePermissionCheckReq{}
	s := service.SysUser{}
	r := service.SysRole{}
	rp := service.SysRolePermission{}
	err := e.MakeContext(c).
		MakeOrm().Bind(&req).
		MakeService(&r.Service).
		MakeService(&s.Service).
		MakeService(&rp.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var roles = make([]string, 1)
	roles[0] = user.GetRoleName(c)
	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"
	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"

	p := actions.GetPermissionFromContext(c)

	rpModel := []models.SysRolePermission{}
	if err := rp.GetByRoleId(user.GetRoleId(c), p, &rpModel); err != nil {
		log.Debugf("get permission by role error: %v", err.Error())
		return
	}

	var passVal = false
	for _, rpItem := range rpModel {
		log.Printf("method: %v   path: %v", req.AuthMethod, req.AuthPath)
		if req.AuthMethod == rpItem.Method &&
			req.AuthPath == rpItem.Path {
			passVal = true
		}
	}
	if passVal {
		e.OK(nil, "permission passed")
	} else {
		e.Error(400, nil, "permission denied")
	}

}

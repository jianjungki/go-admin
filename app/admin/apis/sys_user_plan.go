package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SysUserPlan struct {
	api.Api
}

// GetPage 获取SysUserPlan列表
// @Summary 获取SysUserPlan列表
// @Description 获取SysUserPlan列表
// @Tags SysUserPlan
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysUserPlan}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user-plan [get]
// @Security Bearer
func (e SysUserPlan) GetPage(c *gin.Context) {
	req := map[string]string{}
	s := service.SysUserPlan{}
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

}

// Get 获取SysUserPlan
// @Summary 获取SysUserPlan
// @Description 获取SysUserPlan
// @Tags SysUserPlan
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysUserPlan} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user-plan/{id} [get]
// @Security Bearer
func (e SysUserPlan) Get(c *gin.Context) {
	req := dto.SysUserPlanGetReq{}
	s := service.SysUserPlan{}
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
	var object models.SysUserPlan

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取SysUserPlan失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建SysUserPlan
// @Summary 创建SysUserPlan
// @Description 创建SysUserPlan
// @Tags SysUserPlan
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserPlanInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-user-plan [post]
// @Security Bearer
func (e SysUserPlan) Insert(c *gin.Context) {
	req := dto.SysUserPlanInsertReq{}
	s := service.SysUserPlan{}
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
		e.Error(500, err, fmt.Sprintf("创建SysUserPlan失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改SysUserPlan
// @Summary 修改SysUserPlan
// @Description 修改SysUserPlan
// @Tags SysUserPlan
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysUserPlanUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-user-plan/{id} [put]
// @Security Bearer
func (e SysUserPlan) Update(c *gin.Context) {
	req := dto.SysUserPlanUpdateReq{}
	s := service.SysUserPlan{}
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
		e.Error(500, err, fmt.Sprintf("修改SysUserPlan失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除SysUserPlan
// @Summary 删除SysUserPlan
// @Description 删除SysUserPlan
// @Tags SysUserPlan
// @Param data body dto.SysUserPlanDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-user-plan [delete]
// @Security Bearer
func (e SysUserPlan) Delete(c *gin.Context) {
	s := service.SysUserPlan{}
	req := dto.SysUserPlanDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除SysUserPlan失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

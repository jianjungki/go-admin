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

type SysPlan struct {
	api.Api
}

// GetPage 获取SysPlan列表
// @Summary 获取SysPlan列表
// @Description 获取SysPlan列表
// @Tags SysPlan
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysPlan}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-plan [get]
// @Security Bearer
func (e SysPlan) GetPage(c *gin.Context) {
    req := dto.SysPlanGetPageReq{}
    s := service.SysPlan{}
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
	list := make([]models.SysPlan, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取SysPlan失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取SysPlan
// @Summary 获取SysPlan
// @Description 获取SysPlan
// @Tags SysPlan
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysPlan} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-plan/{id} [get]
// @Security Bearer
func (e SysPlan) Get(c *gin.Context) {
	req := dto.SysPlanGetReq{}
	s := service.SysPlan{}
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
	var object models.SysPlan

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取SysPlan失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建SysPlan
// @Summary 创建SysPlan
// @Description 创建SysPlan
// @Tags SysPlan
// @Accept application/json
// @Product application/json
// @Param data body dto.SysPlanInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-plan [post]
// @Security Bearer
func (e SysPlan) Insert(c *gin.Context) {
    req := dto.SysPlanInsertReq{}
    s := service.SysPlan{}
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
		e.Error(500, err, fmt.Sprintf("创建SysPlan失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改SysPlan
// @Summary 修改SysPlan
// @Description 修改SysPlan
// @Tags SysPlan
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysPlanUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-plan/{id} [put]
// @Security Bearer
func (e SysPlan) Update(c *gin.Context) {
    req := dto.SysPlanUpdateReq{}
    s := service.SysPlan{}
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
		e.Error(500, err, fmt.Sprintf("修改SysPlan失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除SysPlan
// @Summary 删除SysPlan
// @Description 删除SysPlan
// @Tags SysPlan
// @Param data body dto.SysPlanDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-plan [delete]
// @Security Bearer
func (e SysPlan) Delete(c *gin.Context) {
    s := service.SysPlan{}
    req := dto.SysPlanDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除SysPlan失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}

package apis

import (
	"go-admin/app/admin/models"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type OCR struct {
	api.Api
}

// MixOCR
// @Summary 混合多种ocr结果
// @Tags OCR
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/ocr/mix [post]
// @Security Bearer
func (e SysDept) MixOCR(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetPageReq{}
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
	list := make([]models.SysDept, 0)
	list, err = s.SetDeptPage(&req)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list, "查询成功")
}

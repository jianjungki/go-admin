package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SysUserPlan struct {
	service.Service
}

// GetPage 获取SysUserPlan列表
func (e *SysUserPlan) GetPage(c *dto.SysUserPlanGetPageReq, p *actions.DataPermission, list *[]models.SysUserPlan, count *int64) error {
	var err error
	var data models.SysUserPlan

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysUserPlanService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysUserPlan对象
func (e *SysUserPlan) Get(d *dto.SysUserPlanGetReq, p *actions.DataPermission, model *models.SysUserPlan) error {
	var data models.SysUserPlan

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysUserPlan error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysUserPlan对象
func (e *SysUserPlan) Insert(c *dto.SysUserPlanInsertReq) error {
	var err error
	var data models.SysUserPlan
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysUserPlanService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysUserPlan对象
func (e *SysUserPlan) Update(c *dto.SysUserPlanUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SysUserPlan{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("SysUserPlanService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysUserPlan
func (e *SysUserPlan) Remove(d *dto.SysUserPlanDeleteReq, p *actions.DataPermission) error {
	var data models.SysUserPlan

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysUserPlan error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

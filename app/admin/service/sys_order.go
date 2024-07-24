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

type SysOrder struct {
	service.Service
}

// GetPage 获取SysOrder列表
func (e *SysOrder) GetPage(c *dto.SysOrderGetPageReq, p *actions.DataPermission, list *[]models.SysOrder, count *int64) error {
	var err error
	var data models.SysOrder

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysOrderService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysOrder对象
func (e *SysOrder) Get(d *dto.SysOrderGetReq, p *actions.DataPermission, model *models.SysOrder) error {
	var data models.SysOrder

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysOrder error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysOrder对象
func (e *SysOrder) Insert(c *dto.SysOrderInsertReq) error {
    var err error
    var data models.SysOrder
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysOrderService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysOrder对象
func (e *SysOrder) Update(c *dto.SysOrderUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.SysOrder{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SysOrderService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除SysOrder
func (e *SysOrder) Remove(d *dto.SysOrderDeleteReq, p *actions.DataPermission) error {
	var data models.SysOrder

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSysOrder error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

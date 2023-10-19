package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/sequence/models"
	"go-admin/app/sequence/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Location struct {
	service.Service
}

// GetPage 获取Location列表
func (e *Location) GetPage(c *dto.LocationGetPageReq, p *actions.DataPermission, list *[]models.Location, count *int64) error {
	var err error
	var data models.Location

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("LocationService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Location对象
func (e *Location) Get(d *dto.LocationGetReq, p *actions.DataPermission, model *models.Location) error {
	var data models.Location

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetLocation error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Location对象
func (e *Location) Insert(c *dto.LocationInsertReq) error {
    var err error
    var data models.Location
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("LocationService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Location对象
func (e *Location) Update(c *dto.LocationUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Location{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("LocationService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Location
func (e *Location) Remove(d *dto.LocationDeleteReq, p *actions.DataPermission) error {
	var data models.Location

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveLocation error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

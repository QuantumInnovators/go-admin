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

type Meta struct {
	service.Service
}

// GetPage 获取Meta列表
func (e *Meta) GetPage(c *dto.MetaGetPageReq, p *actions.DataPermission, list *[]models.Meta, count *int64) error {
	var err error
	var data models.Meta

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("MetaService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Meta对象
func (e *Meta) Get(d *dto.MetaGetReq, p *actions.DataPermission, model *models.Meta) error {
	var data models.Meta

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetMeta error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Meta对象
func (e *Meta) Insert(c *dto.MetaInsertReq) error {
    var err error
    var data models.Meta
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("MetaService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Meta对象
func (e *Meta) Update(c *dto.MetaUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Meta{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("MetaService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Meta
func (e *Meta) Remove(d *dto.MetaDeleteReq, p *actions.DataPermission) error {
	var data models.Meta

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveMeta error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

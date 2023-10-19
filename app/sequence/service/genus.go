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

type Genus struct {
	service.Service
}

// GetPage 获取Genus列表
func (e *Genus) GetPage(c *dto.GenusGetPageReq, p *actions.DataPermission, list *[]models.Genus, count *int64) error {
	var err error
	var data models.Genus

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("GenusService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Genus对象
func (e *Genus) Get(d *dto.GenusGetReq, p *actions.DataPermission, model *models.Genus) error {
	var data models.Genus

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetGenus error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Genus对象
func (e *Genus) Insert(c *dto.GenusInsertReq) error {
    var err error
    var data models.Genus
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("GenusService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Genus对象
func (e *Genus) Update(c *dto.GenusUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Genus{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("GenusService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Genus
func (e *Genus) Remove(d *dto.GenusDeleteReq, p *actions.DataPermission) error {
	var data models.Genus

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveGenus error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

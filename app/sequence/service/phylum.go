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

type Phylum struct {
	service.Service
}

// GetPage 获取Phylum列表
func (e *Phylum) GetPage(c *dto.PhylumGetPageReq, p *actions.DataPermission, list *[]models.Phylum, count *int64) error {
	var err error
	var data models.Phylum

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("PhylumService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Phylum对象
func (e *Phylum) Get(d *dto.PhylumGetReq, p *actions.DataPermission, model *models.Phylum) error {
	var data models.Phylum

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetPhylum error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Phylum对象
func (e *Phylum) Insert(c *dto.PhylumInsertReq) error {
    var err error
    var data models.Phylum
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("PhylumService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Phylum对象
func (e *Phylum) Update(c *dto.PhylumUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Phylum{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("PhylumService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Phylum
func (e *Phylum) Remove(d *dto.PhylumDeleteReq, p *actions.DataPermission) error {
	var data models.Phylum

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemovePhylum error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

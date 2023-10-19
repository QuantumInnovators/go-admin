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

type Primer struct {
	service.Service
}

// GetPage 获取Primer列表
func (e *Primer) GetPage(c *dto.PrimerGetPageReq, p *actions.DataPermission, list *[]models.Primer, count *int64) error {
	var err error
	var data models.Primer

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("PrimerService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Primer对象
func (e *Primer) Get(d *dto.PrimerGetReq, p *actions.DataPermission, model *models.Primer) error {
	var data models.Primer

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetPrimer error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Primer对象
func (e *Primer) Insert(c *dto.PrimerInsertReq) error {
    var err error
    var data models.Primer
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("PrimerService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Primer对象
func (e *Primer) Update(c *dto.PrimerUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Primer{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("PrimerService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Primer
func (e *Primer) Remove(d *dto.PrimerDeleteReq, p *actions.DataPermission) error {
	var data models.Primer

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemovePrimer error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

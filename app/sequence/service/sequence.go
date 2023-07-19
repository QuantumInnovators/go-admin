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

type Sequence struct {
	service.Service
}

func getTableName(name string, data models.Sequence) string {
	if name != "" {
		return name
	}
	return data.TableName()
}

// GetPage 获取Sequence列表
func (e *Sequence) GetPage(c *dto.SequenceGetPageReq, p *actions.DataPermission, list *[]models.Sequence, count *int64) error {
	var err error
	var data models.Sequence

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SequenceService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Sequence对象
func (e *Sequence) Get(d *dto.SequenceGetReq, p *actions.DataPermission, model *models.Sequence) error {
	var data models.Sequence

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSequence error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Sequence对象
func (e *Sequence) Insert(c *dto.SequenceInsertReq) error {
	var err error
	var data models.Sequence
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SequenceService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Sequence对象
func (e *Sequence) Update(c *dto.SequenceUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Sequence{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("SequenceService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Sequence
func (e *Sequence) Remove(d *dto.SequenceDeleteReq, p *actions.DataPermission) error {
	var data models.Sequence

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSequence error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

func (e *Sequence) GetByClass(d *dto.SequenceGetByClassReq, p *actions.DataPermission, model *models.Sequence) error {
	var data models.Sequence

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetClassID()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSequence error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

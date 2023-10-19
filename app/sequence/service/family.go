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

type Family struct {
	service.Service
}

// GetPage 获取Family列表
func (e *Family) GetPage(c *dto.FamilyGetPageReq, p *actions.DataPermission, list *[]models.Family, count *int64) error {
	var err error
	var data models.Family

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FamilyService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Family对象
func (e *Family) Get(d *dto.FamilyGetReq, p *actions.DataPermission, model *models.Family) error {
	var data models.Family

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFamily error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Family对象
func (e *Family) Insert(c *dto.FamilyInsertReq) error {
    var err error
    var data models.Family
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FamilyService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Family对象
func (e *Family) Update(c *dto.FamilyUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Family{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("FamilyService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Family
func (e *Family) Remove(d *dto.FamilyDeleteReq, p *actions.DataPermission) error {
	var data models.Family

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveFamily error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

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

type Kingdom struct {
	service.Service
}

// GetPage 获取Kingdom列表
func (e *Kingdom) GetPage(c *dto.KingdomGetPageReq, p *actions.DataPermission, list *[]models.Kingdom, count *int64) error {
	var err error
	var data models.Kingdom

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("KingdomService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Kingdom对象
func (e *Kingdom) Get(d *dto.KingdomGetReq, p *actions.DataPermission, model *models.Kingdom) error {
	var data models.Kingdom

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetKingdom error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Kingdom对象
func (e *Kingdom) Insert(c *dto.KingdomInsertReq) error {
    var err error
    var data models.Kingdom
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("KingdomService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Kingdom对象
func (e *Kingdom) Update(c *dto.KingdomUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Kingdom{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("KingdomService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Kingdom
func (e *Kingdom) Remove(d *dto.KingdomDeleteReq, p *actions.DataPermission) error {
	var data models.Kingdom

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveKingdom error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

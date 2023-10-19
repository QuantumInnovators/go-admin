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

type Order struct {
	service.Service
}

// GetPage 获取Order列表
func (e *Order) GetPage(c *dto.OrderGetPageReq, p *actions.DataPermission, list *[]models.Order, count *int64) error {
	var err error
	var data models.Order

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrderService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Order对象
func (e *Order) Get(d *dto.OrderGetReq, p *actions.DataPermission, model *models.Order) error {
	var data models.Order

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrder error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Order对象
func (e *Order) Insert(c *dto.OrderInsertReq) error {
    var err error
    var data models.Order
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrderService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Order对象
func (e *Order) Update(c *dto.OrderUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Order{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrderService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Order
func (e *Order) Remove(d *dto.OrderDeleteReq, p *actions.DataPermission) error {
	var data models.Order

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrder error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

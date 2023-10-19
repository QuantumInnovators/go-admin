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

type Class struct {
	service.Service
}

// GetPage 获取Class列表
func (e *Class) GetPage(c *dto.ClassGetPageReq, p *actions.DataPermission, list *[]models.Class, count *int64) error {
	var err error
	var data models.Class

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ClassService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Class对象
func (e *Class) Get(d *dto.ClassGetReq, p *actions.DataPermission, model *models.Class) error {
	var data models.Class

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetClass error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Class对象
func (e *Class) Insert(c *dto.ClassInsertReq) error {
    var err error
    var data models.Class
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ClassService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Class对象
func (e *Class) Update(c *dto.ClassUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Class{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("ClassService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Class
func (e *Class) Remove(d *dto.ClassDeleteReq, p *actions.DataPermission) error {
	var data models.Class

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveClass error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

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

type Species struct {
	service.Service
}

// GetPage 获取Species列表
func (e *Species) GetPage(c *dto.SpeciesGetPageReq, p *actions.DataPermission, list *[]models.Species, count *int64) error {
	var err error
	var data models.Species

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SpeciesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Species对象
func (e *Species) Get(d *dto.SpeciesGetReq, p *actions.DataPermission, model *models.Species) error {
	var data models.Species

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSpecies error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Species对象
func (e *Species) Insert(c *dto.SpeciesInsertReq) error {
    var err error
    var data models.Species
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SpeciesService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Species对象
func (e *Species) Update(c *dto.SpeciesUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Species{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SpeciesService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Species
func (e *Species) Remove(d *dto.SpeciesDeleteReq, p *actions.DataPermission) error {
	var data models.Species

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSpecies error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

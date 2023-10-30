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

type Projects struct {
	service.Service
}

// GetPage 获取Projects列表
func (e *Projects) GetPage(c *dto.ProjectsGetPageReq, p *actions.DataPermission, list *[]models.Projects, count *int64) error {
	var err error
	var data models.Projects

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ProjectsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Projects对象
func (e *Projects) Get(d *dto.ProjectsGetReq, p *actions.DataPermission, model *models.Projects) error {
	var data models.Projects

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetProjects error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Projects对象
func (e *Projects) Insert(c *dto.ProjectsInsertReq) error {
    var err error
    var data models.Projects
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ProjectsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Projects对象
func (e *Projects) Update(c *dto.ProjectsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Projects{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("ProjectsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Projects
func (e *Projects) Remove(d *dto.ProjectsDeleteReq, p *actions.DataPermission) error {
	var data models.Projects

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveProjects error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

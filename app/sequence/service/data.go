package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/sequence/models"
	"go-admin/common/actions"
)

type ClassList struct {
	service.Service
}

// GetTotalClassList 获取ClassList列表
func (e *ClassList) GetTotalClassList(p *actions.DataPermission, list *models.ClassList) error {
	listKingdom := make([]models.Kingdom, 0)
	var kingdom models.Kingdom
	err := e.Orm.Model(&kingdom).
		Scopes(
			actions.Permission(kingdom.TableName(), p),
		).
		Find(&listKingdom).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("ClassService GetKingdom error:%s \r\n", err)
		return err
	}
	// 门
	var listPhylum = make([]models.Phylum, 0)
	var phylum models.Phylum
	err = e.Orm.Model(&phylum).Scopes(actions.Permission(phylum.TableName(), p)).Find(&listPhylum).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("ClassService GetPhylum error:%s \r\n", err)
		return err
	}
	// 纲
	var listCategory = make([]models.Category, 0)
	var category models.Category
	err = e.Orm.Model(&category).Scopes(actions.Permission(category.TableName(), p)).Find(&listCategory).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("ClassService GetCategory error:%s \r\n", err)
		return err
	}

	// 组合返回list
	for _, v := range listKingdom {
		list.Data = append(list.Data, &models.ClassListKingdom{
			Data: v,
			Id:   v.Id,
			Desc: v.Desc,
		})
	}
	for _, v := range listPhylum {
		for _, v1 := range list.Data {
			if v1.Data.Id == v.KingdomId {
				v1.Phylum = append(v1.Phylum, &models.ClassListPhylum{
					Id:   v.Id,
					Desc: v.Desc,
					Data: v,
				})
			}
		}
	}
	return nil
}

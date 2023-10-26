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
	var listClass = make([]models.Class, 0)
	var class models.Class
	err = e.Orm.Model(&class).Scopes(actions.Permission(class.TableName(), p)).Find(&listClass).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("ClassService GetClass error:%s \r\n", err)
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
			if int64(v1.Data.Id) == v.ParentId {
				v1.Phylum = append(v1.Phylum, &models.ClassListPhylum{
					Id:   v.Id,
					Desc: v.Desc,
					Data: v,
				})
			}
		}
	}
	for _, v := range listClass {
		for _, v1 := range list.Data {
			for _, v2 := range v1.Phylum {
				if int64(v2.Data.Id) == v.ParentId {
					v2.Class = append(v2.Class, &models.ClassListClass{
						Id:   v.Id,
						Desc: v.Desc,
						Data: v,
					})
				}
			}
		}
	}
	return nil
}

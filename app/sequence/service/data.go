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
	// 目
	var listOrder = make([]models.Order, 0)
	var order models.Order
	err = e.Orm.Model(&order).Scopes(actions.Permission(order.TableName(), p)).Find(&listOrder).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("ClassService GetOrder error:%s \r\n", err)
		return err
	}
	// 科
	var listFamily = make([]models.Family, 0)
	var family models.Family
	err = e.Orm.Model(&family).Scopes(actions.Permission(family.TableName(), p)).Find(&listFamily).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("ClassService GetFamily error:%s \r\n", err)
		return err
	}
	// 属
	var listGenus = make([]models.Genus, 0)
	var genus models.Genus
	err = e.Orm.Model(&genus).Scopes(actions.Permission(genus.TableName(), p)).Find(&listGenus).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("ClassService GetGenus error:%s \r\n", err)
		return err
	}
	// 种
	var listSpecies = make([]models.Species, 0)
	var species models.Species
	err = e.Orm.Model(&species).Scopes(actions.Permission(species.TableName(), p)).Find(&listSpecies).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("ClassService GetSpecies error:%s \r\n", err)
		return err
	}

	// 组合返回list
	for _, v := range listKingdom {
		list.Data = append(list.Data, &models.ClassData{
			Id:    v.Id,
			Label: v.Desc,
		})
	}
	for _, v := range listPhylum {
		for _, v1 := range list.Data {
			if int64(v1.Id) == v.ParentId {
				v1.Children = append(v1.Children, &models.ClassData{
					Id:    v.Id,
					Label: v.NameCn,
				})
			}
		}
	}
	for _, v := range listClass {
		for _, v1 := range list.Data {
			for _, v2 := range v1.Children {
				if int64(v2.Id) == v.ParentId {
					v2.Children = append(v2.Children, &models.ClassData{
						Id:    v.Id,
						Label: v.NameCn,
					})
				}
			}
		}
	}
	for _, v := range listOrder {
		for _, v1 := range list.Data {
			for _, v2 := range v1.Children {
				for _, v3 := range v2.Children {
					if int64(v3.Id) == v.ParentId {
						v3.Children = append(v3.Children, &models.ClassData{
							Id:    v.Id,
							Label: v.NameCn,
						})
					}
				}
			}
		}
	}
	for _, v := range listFamily {
		for _, v1 := range list.Data {
			for _, v2 := range v1.Children {
				for _, v3 := range v2.Children {
					for _, v4 := range v3.Children {
						if int64(v4.Id) == v.ParentId {
							v4.Children = append(v4.Children, &models.ClassData{
								Id:    v.Id,
								Label: v.NameCn,
							})
						}
					}
				}
			}
		}
	}
	// todo zyx 优化写法
	return nil
}

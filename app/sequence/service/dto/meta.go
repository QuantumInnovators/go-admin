package dto

import (

	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type MetaGetPageReq struct {
	dto.Pagination     `search:"-"`
    MetaOrder
}

type MetaOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:meta"`
    CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:meta"`
    KingdomId string `form:"kingdomIdOrder"  search:"type:order;column:kingdom_id;table:meta"`
    PhylumId string `form:"phylumIdOrder"  search:"type:order;column:phylum_id;table:meta"`
    ClassId string `form:"classIdOrder"  search:"type:order;column:class_id;table:meta"`
    OrderId string `form:"orderIdOrder"  search:"type:order;column:order_id;table:meta"`
    FamilyId string `form:"familyIdOrder"  search:"type:order;column:family_id;table:meta"`
    GenusId string `form:"genusIdOrder"  search:"type:order;column:genus_id;table:meta"`
    SpeciesId string `form:"speciesIdOrder"  search:"type:order;column:species_id;table:meta"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:meta"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:meta"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:meta"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:meta"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:meta"`
    
}

func (m *MetaGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type MetaInsertReq struct {
    Id int `json:"-" comment:"自增id"` // 自增id
    CategoryId string `json:"categoryId" comment:"类别"`
    KingdomId string `json:"kingdomId" comment:"界"`
    PhylumId string `json:"phylumId" comment:"门"`
    ClassId string `json:"classId" comment:"纲"`
    OrderId string `json:"orderId" comment:"目"`
    FamilyId string `json:"familyId" comment:"科"`
    GenusId string `json:"genusId" comment:"属"`
    SpeciesId string `json:"speciesId" comment:"种"`
    common.ControlBy
}

func (s *MetaInsertReq) Generate(model *models.Meta)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CategoryId = s.CategoryId
    model.KingdomId = s.KingdomId
    model.PhylumId = s.PhylumId
    model.ClassId = s.ClassId
    model.OrderId = s.OrderId
    model.FamilyId = s.FamilyId
    model.GenusId = s.GenusId
    model.SpeciesId = s.SpeciesId
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *MetaInsertReq) GetId() interface{} {
	return s.Id
}

type MetaUpdateReq struct {
    Id int `uri:"id" comment:"自增id"` // 自增id
    CategoryId string `json:"categoryId" comment:"类别"`
    KingdomId string `json:"kingdomId" comment:"界"`
    PhylumId string `json:"phylumId" comment:"门"`
    ClassId string `json:"classId" comment:"纲"`
    OrderId string `json:"orderId" comment:"目"`
    FamilyId string `json:"familyId" comment:"科"`
    GenusId string `json:"genusId" comment:"属"`
    SpeciesId string `json:"speciesId" comment:"种"`
    common.ControlBy
}

func (s *MetaUpdateReq) Generate(model *models.Meta)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CategoryId = s.CategoryId
    model.KingdomId = s.KingdomId
    model.PhylumId = s.PhylumId
    model.ClassId = s.ClassId
    model.OrderId = s.OrderId
    model.FamilyId = s.FamilyId
    model.GenusId = s.GenusId
    model.SpeciesId = s.SpeciesId
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *MetaUpdateReq) GetId() interface{} {
	return s.Id
}

// MetaGetReq 功能获取请求参数
type MetaGetReq struct {
     Id int `uri:"id"`
}
func (s *MetaGetReq) GetId() interface{} {
	return s.Id
}

// MetaDeleteReq 功能删除请求参数
type MetaDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *MetaDeleteReq) GetId() interface{} {
	return s.Ids
}

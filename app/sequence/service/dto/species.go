package dto

import (
	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SpeciesGetPageReq struct {
	dto.Pagination `search:"-"`
	SpeciesOrder
}

type SpeciesOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:species"`
	Name       string `form:"nameOrder"  search:"type:order;column:name;table:species"`
	NameCn     string `form:"nameCnOrder"  search:"type:order;column:name_cn;table:species"`
	Desc       string `form:"descOrder"  search:"type:order;column:desc;table:species"`
	ParentId   string `form:"parentIdOrder"  search:"type:order;column:parent_id;table:species"`
	CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:species"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:species"`
	UpdatedAt  string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:species"`
	DeletedAt  string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:species"`
	CreateBy   string `form:"createByOrder"  search:"type:order;column:create_by;table:species"`
	UpdateBy   string `form:"updateByOrder"  search:"type:order;column:update_by;table:species"`
}

func (m *SpeciesGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SpeciesInsertReq struct {
	Id         int    `json:"-" comment:"主键编码"` // 主键编码
	Name       string `json:"name" comment:"拉丁文名"`
	NameCn     string `json:"nameCn" comment:"中文名"`
	Desc       string `json:"desc" comment:"Desc"`
	ParentId   int64  `json:"parentId" comment:"所属属"`
	CategoryId string `json:"categoryId" comment:"所属类"`
	common.ControlBy
}

func (s *SpeciesInsertReq) Generate(model *models.Species) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.NameCn = s.NameCn
	model.Desc = s.Desc
	model.ParentId = s.ParentId
	model.CategoryId = s.CategoryId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SpeciesInsertReq) GetId() interface{} {
	return s.Id
}

type SpeciesUpdateReq struct {
	Id         int    `uri:"id" comment:"主键编码"` // 主键编码
	Name       string `json:"name" comment:"拉丁文名"`
	NameCn     string `json:"nameCn" comment:"中文名"`
	Desc       string `json:"desc" comment:"Desc"`
	ParentId   int64  `json:"parentId" comment:"所属属"`
	CategoryId string `json:"categoryId" comment:"所属类"`
	common.ControlBy
}

func (s *SpeciesUpdateReq) Generate(model *models.Species) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.NameCn = s.NameCn
	model.Desc = s.Desc
	model.ParentId = s.ParentId
	model.CategoryId = s.CategoryId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SpeciesUpdateReq) GetId() interface{} {
	return s.Id
}

// SpeciesGetReq 功能获取请求参数
type SpeciesGetReq struct {
	Id int `uri:"id"`
}

func (s *SpeciesGetReq) GetId() interface{} {
	return s.Id
}

// SpeciesDeleteReq 功能删除请求参数
type SpeciesDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SpeciesDeleteReq) GetId() interface{} {
	return s.Ids
}

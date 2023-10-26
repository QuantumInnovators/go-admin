package dto

import (
	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type PhylumGetPageReq struct {
	dto.Pagination `search:"-"`
	PhylumOrder
}

type PhylumOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:phylum"`
	Name       string `form:"nameOrder"  search:"type:order;column:name;table:phylum"`
	NameCn     string `form:"nameCnOrder"  search:"type:order;column:name_cn;table:phylum"`
	Desc       string `form:"descOrder"  search:"type:order;column:desc;table:phylum"`
	ParentId   string `form:"parentIdOrder"  search:"type:order;column:parent_id;table:phylum"`
	CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:phylum"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:phylum"`
	UpdatedAt  string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:phylum"`
	DeletedAt  string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:phylum"`
	CreateBy   string `form:"createByOrder"  search:"type:order;column:create_by;table:phylum"`
	UpdateBy   string `form:"updateByOrder"  search:"type:order;column:update_by;table:phylum"`
}

func (m *PhylumGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type PhylumInsertReq struct {
	Id         int    `json:"-" comment:"主键编码"` // 主键编码
	Name       string `json:"name" comment:"拉丁文名"`
	NameCn     string `json:"nameCn" comment:"中文名"`
	Desc       string `json:"desc" comment:"Desc"`
	ParentId   int64  `json:"parentId" comment:"所属界"`
	CategoryId string `json:"categoryId" comment:"所属类"`
	common.ControlBy
}

func (s *PhylumInsertReq) Generate(model *models.Phylum) {
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

func (s *PhylumInsertReq) GetId() interface{} {
	return s.Id
}

type PhylumUpdateReq struct {
	Id         int    `uri:"id" comment:"主键编码"` // 主键编码
	Name       string `json:"name" comment:"拉丁文名"`
	NameCn     string `json:"nameCn" comment:"中文名"`
	Desc       string `json:"desc" comment:"Desc"`
	ParentId   int64  `json:"parentId" comment:"所属界"`
	CategoryId string `json:"categoryId" comment:"所属类"`
	common.ControlBy
}

func (s *PhylumUpdateReq) Generate(model *models.Phylum) {
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

func (s *PhylumUpdateReq) GetId() interface{} {
	return s.Id
}

// PhylumGetReq 功能获取请求参数
type PhylumGetReq struct {
	Id int `uri:"id"`
}

func (s *PhylumGetReq) GetId() interface{} {
	return s.Id
}

// PhylumDeleteReq 功能删除请求参数
type PhylumDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *PhylumDeleteReq) GetId() interface{} {
	return s.Ids
}

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
	Desc       string `form:"descOrder"  search:"type:order;column:desc;table:phylum"`
	KingdomId  string `form:"kingdomIdOrder"  search:"type:order;column:kingdom_id;table:phylum"`
	CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:phylum"`
	DeletedAt  string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:phylum"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:phylum"`
	UpdatedAt  string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:phylum"`
	CreateBy   string `form:"createByOrder"  search:"type:order;column:create_by;table:phylum"`
	UpdateBy   string `form:"updateByOrder"  search:"type:order;column:update_by;table:phylum"`
}

func (m *PhylumGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type PhylumInsertReq struct {
	Id         int    `json:"-" comment:""` //
	Name       string `json:"name" comment:"拉丁文名"`
	Desc       string `json:"desc" comment:""`
	KingdomId  int    `json:"kingdomId" comment:"所属界"`
	CategoryId string `json:"categoryId" comment:"所属种群"`
	common.ControlBy
}

func (s *PhylumInsertReq) Generate(model *models.Phylum) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Desc = s.Desc
	model.KingdomId = s.KingdomId
	model.CategoryId = s.CategoryId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *PhylumInsertReq) GetId() interface{} {
	return s.Id
}

type PhylumUpdateReq struct {
	Id         int    `uri:"id" comment:""` //
	Name       string `json:"name" comment:"拉丁文名"`
	Desc       string `json:"desc" comment:""`
	KingdomId  int    `json:"kingdomId" comment:"所属界"`
	CategoryId string `json:"categoryId" comment:"所属种群"`
	common.ControlBy
}

func (s *PhylumUpdateReq) Generate(model *models.Phylum) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Desc = s.Desc
	model.KingdomId = s.KingdomId
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

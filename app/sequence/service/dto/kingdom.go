package dto

import (
	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type KingdomGetPageReq struct {
	dto.Pagination `search:"-"`
	KingdomOrder
}

type KingdomOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:kingdom"`
	Name       string `form:"nameOrder"  search:"type:order;column:name;table:kingdom"`
	Desc       string `form:"descOrder"  search:"type:order;column:desc;table:kingdom"`
	CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:kingdom"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:kingdom"`
	UpdatedAt  string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:kingdom"`
	DeletedAt  string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:kingdom"`
	CreateBy   string `form:"createByOrder"  search:"type:order;column:create_by;table:kingdom"`
	UpdateBy   string `form:"updateByOrder"  search:"type:order;column:update_by;table:kingdom"`
}

func (m *KingdomGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type KingdomInsertReq struct {
	Id         int    `json:"-" comment:"主键编码"` // 主键编码
	Name       string `json:"name" comment:"拉丁文名"`
	Desc       string `json:"desc" comment:"中文名"`
	CategoryId string `json:"categoryId" comment:"所属类"`
	common.ControlBy
}

func (s *KingdomInsertReq) Generate(model *models.Kingdom) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Desc = s.Desc
	model.CategoryId = s.CategoryId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *KingdomInsertReq) GetId() interface{} {
	return s.Id
}

type KingdomUpdateReq struct {
	Id         int    `uri:"id" comment:"主键编码"` // 主键编码
	Name       string `json:"name" comment:"拉丁文名"`
	Desc       string `json:"desc" comment:"中文名"`
	CategoryId string `json:"categoryId" comment:"所属类"`
	common.ControlBy
}

func (s *KingdomUpdateReq) Generate(model *models.Kingdom) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Desc = s.Desc
	model.CategoryId = s.CategoryId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *KingdomUpdateReq) GetId() interface{} {
	return s.Id
}

// KingdomGetReq 功能获取请求参数
type KingdomGetReq struct {
	Id int `uri:"id"`
}

func (s *KingdomGetReq) GetId() interface{} {
	return s.Id
}

// KingdomDeleteReq 功能删除请求参数
type KingdomDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *KingdomDeleteReq) GetId() interface{} {
	return s.Ids
}

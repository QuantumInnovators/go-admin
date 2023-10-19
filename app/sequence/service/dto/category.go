package dto

import (

	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CategoryGetPageReq struct {
	dto.Pagination     `search:"-"`
    CategoryOrder
}

type CategoryOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:category"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:category"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:category"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:category"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:category"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:category"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:category"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:category"`
    
}

func (m *CategoryGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type CategoryInsertReq struct {
    Id int `json:"-" comment:"自增id"` // 自增id
    Name string `json:"name" comment:"拉丁文名称"`
    Desc string `json:"desc" comment:"描述"`
    common.ControlBy
}

func (s *CategoryInsertReq) Generate(model *models.Category)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *CategoryInsertReq) GetId() interface{} {
	return s.Id
}

type CategoryUpdateReq struct {
    Id int `uri:"id" comment:"自增id"` // 自增id
    Name string `json:"name" comment:"拉丁文名称"`
    Desc string `json:"desc" comment:"描述"`
    common.ControlBy
}

func (s *CategoryUpdateReq) Generate(model *models.Category)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *CategoryUpdateReq) GetId() interface{} {
	return s.Id
}

// CategoryGetReq 功能获取请求参数
type CategoryGetReq struct {
     Id int `uri:"id"`
}
func (s *CategoryGetReq) GetId() interface{} {
	return s.Id
}

// CategoryDeleteReq 功能删除请求参数
type CategoryDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *CategoryDeleteReq) GetId() interface{} {
	return s.Ids
}

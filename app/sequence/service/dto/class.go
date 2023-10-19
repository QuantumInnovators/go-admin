package dto

import (

	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ClassGetPageReq struct {
	dto.Pagination     `search:"-"`
    ClassOrder
}

type ClassOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:class"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:class"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:class"`
    PhylumId string `form:"phylumIdOrder"  search:"type:order;column:phylum_id;table:class"`
    CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:class"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:class"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:class"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:class"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:class"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:class"`
    
}

func (m *ClassGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ClassInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:"拉丁文名"`
    Desc string `json:"desc" comment:""`
    PhylumId string `json:"phylumId" comment:"所属门"`
    CategoryId string `json:"categoryId" comment:"所属类"`
    common.ControlBy
}

func (s *ClassInsertReq) Generate(model *models.Class)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.PhylumId = s.PhylumId
    model.CategoryId = s.CategoryId
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ClassInsertReq) GetId() interface{} {
	return s.Id
}

type ClassUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:"拉丁文名"`
    Desc string `json:"desc" comment:""`
    PhylumId string `json:"phylumId" comment:"所属门"`
    CategoryId string `json:"categoryId" comment:"所属类"`
    common.ControlBy
}

func (s *ClassUpdateReq) Generate(model *models.Class)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.PhylumId = s.PhylumId
    model.CategoryId = s.CategoryId
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ClassUpdateReq) GetId() interface{} {
	return s.Id
}

// ClassGetReq 功能获取请求参数
type ClassGetReq struct {
     Id int `uri:"id"`
}
func (s *ClassGetReq) GetId() interface{} {
	return s.Id
}

// ClassDeleteReq 功能删除请求参数
type ClassDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ClassDeleteReq) GetId() interface{} {
	return s.Ids
}

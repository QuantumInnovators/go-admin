package dto

import (

	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrderGetPageReq struct {
	dto.Pagination     `search:"-"`
    OrderOrder
}

type OrderOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:order"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:order"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:order"`
    ClassId string `form:"classIdOrder"  search:"type:order;column:class_id;table:order"`
    CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:order"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:order"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:order"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:order"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:order"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:order"`
    
}

func (m *OrderGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrderInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:"拉丁文名"`
    Desc string `json:"desc" comment:""`
    ClassId string `json:"classId" comment:"所属纲"`
    CategoryId string `json:"categoryId" comment:"所属类"`
    common.ControlBy
}

func (s *OrderInsertReq) Generate(model *models.Order)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.ClassId = s.ClassId
    model.CategoryId = s.CategoryId
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrderInsertReq) GetId() interface{} {
	return s.Id
}

type OrderUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:"拉丁文名"`
    Desc string `json:"desc" comment:""`
    ClassId string `json:"classId" comment:"所属纲"`
    CategoryId string `json:"categoryId" comment:"所属类"`
    common.ControlBy
}

func (s *OrderUpdateReq) Generate(model *models.Order)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.ClassId = s.ClassId
    model.CategoryId = s.CategoryId
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrderUpdateReq) GetId() interface{} {
	return s.Id
}

// OrderGetReq 功能获取请求参数
type OrderGetReq struct {
     Id int `uri:"id"`
}
func (s *OrderGetReq) GetId() interface{} {
	return s.Id
}

// OrderDeleteReq 功能删除请求参数
type OrderDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrderDeleteReq) GetId() interface{} {
	return s.Ids
}

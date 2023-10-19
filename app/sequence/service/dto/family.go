package dto

import (

	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FamilyGetPageReq struct {
	dto.Pagination     `search:"-"`
    FamilyOrder
}

type FamilyOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:family"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:family"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:family"`
    OrderId string `form:"orderIdOrder"  search:"type:order;column:order_id;table:family"`
    CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:family"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:family"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:family"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:family"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:family"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:family"`
    
}

func (m *FamilyGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FamilyInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:"拉丁文名"`
    Desc string `json:"desc" comment:""`
    OrderId string `json:"orderId" comment:"所属目"`
    CategoryId string `json:"categoryId" comment:"所属类"`
    common.ControlBy
}

func (s *FamilyInsertReq) Generate(model *models.Family)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.OrderId = s.OrderId
    model.CategoryId = s.CategoryId
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FamilyInsertReq) GetId() interface{} {
	return s.Id
}

type FamilyUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:"拉丁文名"`
    Desc string `json:"desc" comment:""`
    OrderId string `json:"orderId" comment:"所属目"`
    CategoryId string `json:"categoryId" comment:"所属类"`
    common.ControlBy
}

func (s *FamilyUpdateReq) Generate(model *models.Family)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.OrderId = s.OrderId
    model.CategoryId = s.CategoryId
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FamilyUpdateReq) GetId() interface{} {
	return s.Id
}

// FamilyGetReq 功能获取请求参数
type FamilyGetReq struct {
     Id int `uri:"id"`
}
func (s *FamilyGetReq) GetId() interface{} {
	return s.Id
}

// FamilyDeleteReq 功能删除请求参数
type FamilyDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FamilyDeleteReq) GetId() interface{} {
	return s.Ids
}

package dto

import (

	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type GenusGetPageReq struct {
	dto.Pagination     `search:"-"`
    GenusOrder
}

type GenusOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:genus"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:genus"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:genus"`
    FamilyId string `form:"familyIdOrder"  search:"type:order;column:family_id;table:genus"`
    CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:genus"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:genus"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:genus"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:genus"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:genus"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:genus"`
    
}

func (m *GenusGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type GenusInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:"拉丁文名"`
    Desc string `json:"desc" comment:""`
    FamilyId string `json:"familyId" comment:"所属科"`
    CategoryId string `json:"categoryId" comment:"所属类"`
    common.ControlBy
}

func (s *GenusInsertReq) Generate(model *models.Genus)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.FamilyId = s.FamilyId
    model.CategoryId = s.CategoryId
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *GenusInsertReq) GetId() interface{} {
	return s.Id
}

type GenusUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:"拉丁文名"`
    Desc string `json:"desc" comment:""`
    FamilyId string `json:"familyId" comment:"所属科"`
    CategoryId string `json:"categoryId" comment:"所属类"`
    common.ControlBy
}

func (s *GenusUpdateReq) Generate(model *models.Genus)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.FamilyId = s.FamilyId
    model.CategoryId = s.CategoryId
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *GenusUpdateReq) GetId() interface{} {
	return s.Id
}

// GenusGetReq 功能获取请求参数
type GenusGetReq struct {
     Id int `uri:"id"`
}
func (s *GenusGetReq) GetId() interface{} {
	return s.Id
}

// GenusDeleteReq 功能删除请求参数
type GenusDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *GenusDeleteReq) GetId() interface{} {
	return s.Ids
}

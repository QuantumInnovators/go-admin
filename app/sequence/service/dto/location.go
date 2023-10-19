package dto

import (

	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type LocationGetPageReq struct {
	dto.Pagination     `search:"-"`
    LocationOrder
}

type LocationOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:location"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:location"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:location"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:location"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:location"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:location"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:location"`
    
}

func (m *LocationGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type LocationInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:"地点简称"`
    common.ControlBy
}

func (s *LocationInsertReq) Generate(model *models.Location)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *LocationInsertReq) GetId() interface{} {
	return s.Id
}

type LocationUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:"地点简称"`
    common.ControlBy
}

func (s *LocationUpdateReq) Generate(model *models.Location)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *LocationUpdateReq) GetId() interface{} {
	return s.Id
}

// LocationGetReq 功能获取请求参数
type LocationGetReq struct {
     Id int `uri:"id"`
}
func (s *LocationGetReq) GetId() interface{} {
	return s.Id
}

// LocationDeleteReq 功能删除请求参数
type LocationDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *LocationDeleteReq) GetId() interface{} {
	return s.Ids
}

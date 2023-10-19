package dto

import (

	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type NotationGetPageReq struct {
	dto.Pagination     `search:"-"`
    NotationOrder
}

type NotationOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:notation"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:notation"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:notation"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:notation"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:notation"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:notation"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:notation"`
    
}

func (m *NotationGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type NotationInsertReq struct {
    Id int `json:"-" comment:""` // 
    Desc string `json:"desc" comment:"描述"`
    common.ControlBy
}

func (s *NotationInsertReq) Generate(model *models.Notation)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Desc = s.Desc
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *NotationInsertReq) GetId() interface{} {
	return s.Id
}

type NotationUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Desc string `json:"desc" comment:"描述"`
    common.ControlBy
}

func (s *NotationUpdateReq) Generate(model *models.Notation)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Desc = s.Desc
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *NotationUpdateReq) GetId() interface{} {
	return s.Id
}

// NotationGetReq 功能获取请求参数
type NotationGetReq struct {
     Id int `uri:"id"`
}
func (s *NotationGetReq) GetId() interface{} {
	return s.Id
}

// NotationDeleteReq 功能删除请求参数
type NotationDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *NotationDeleteReq) GetId() interface{} {
	return s.Ids
}

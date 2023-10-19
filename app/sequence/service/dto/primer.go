package dto

import (

	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type PrimerGetPageReq struct {
	dto.Pagination     `search:"-"`
    PrimerOrder
}

type PrimerOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:primer"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:primer"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:primer"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:primer"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:primer"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:primer"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:primer"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:primer"`
    
}

func (m *PrimerGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type PrimerInsertReq struct {
    Id int `json:"-" comment:"自增id"` // 自增id
    Name string `json:"name" comment:"拉丁文名称"`
    Desc string `json:"desc" comment:"描述"`
    common.ControlBy
}

func (s *PrimerInsertReq) Generate(model *models.Primer)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *PrimerInsertReq) GetId() interface{} {
	return s.Id
}

type PrimerUpdateReq struct {
    Id int `uri:"id" comment:"自增id"` // 自增id
    Name string `json:"name" comment:"拉丁文名称"`
    Desc string `json:"desc" comment:"描述"`
    common.ControlBy
}

func (s *PrimerUpdateReq) Generate(model *models.Primer)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Desc = s.Desc
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *PrimerUpdateReq) GetId() interface{} {
	return s.Id
}

// PrimerGetReq 功能获取请求参数
type PrimerGetReq struct {
     Id int `uri:"id"`
}
func (s *PrimerGetReq) GetId() interface{} {
	return s.Id
}

// PrimerDeleteReq 功能删除请求参数
type PrimerDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *PrimerDeleteReq) GetId() interface{} {
	return s.Ids
}

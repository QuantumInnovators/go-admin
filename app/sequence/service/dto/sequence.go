package dto

import (
	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SequenceGetPageReq struct {
	dto.Pagination `search:"-"`
	SequenceOrder
}

type SequenceOrder struct {
	Id                  string `form:"idOrder"  search:"type:order;column:id;table:sequence"`
	SequenceId          string `form:"sequenceIdOrder"  search:"type:order;column:sequence_id;table:sequence"`
	SequenceDescription string `form:"sequenceDescriptionOrder"  search:"type:order;column:sequence_description;table:sequence"`
	Sequence            string `form:"sequenceOrder"  search:"type:order;column:sequence;table:sequence"`
}

func (m *SequenceGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SequenceInsertReq struct {
	Id                  int    `json:"-" comment:""` //
	SequenceId          string `json:"sequenceId" comment:""`
	SequenceDescription string `json:"sequenceDescription" comment:""`
	Sequence            string `json:"sequence" comment:""`
	common.ControlBy
}

func (s *SequenceInsertReq) Generate(model *models.Sequence) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.SequenceId = s.SequenceId
	model.SequenceDescription = s.SequenceDescription
	model.Sequence = s.Sequence
}

func (s *SequenceInsertReq) GetId() interface{} {
	return s.Id
}

type SequenceUpdateReq struct {
	Id                  int    `uri:"id" comment:""` //
	SequenceId          string `json:"sequenceId" comment:""`
	SequenceDescription string `json:"sequenceDescription" comment:""`
	Sequence            string `json:"sequence" comment:""`
	common.ControlBy
}

func (s *SequenceUpdateReq) Generate(model *models.Sequence) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.SequenceId = s.SequenceId
	model.SequenceDescription = s.SequenceDescription
	model.Sequence = s.Sequence
}

func (s *SequenceUpdateReq) GetId() interface{} {
	return s.Id
}

// SequenceGetReq 功能获取请求参数
type SequenceGetReq struct {
	Id int `uri:"id"`
}

func (s *SequenceGetReq) GetId() interface{} {
	return s.Id
}

// SequenceDeleteReq 功能删除请求参数
type SequenceDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SequenceDeleteReq) GetId() interface{} {
	return s.Ids
}

type SequenceGetByClassReq struct {
	ClassID int `json:"classId"`
}

func (s *SequenceGetByClassReq) GetClassID() interface{} {
	return s.ClassID
}

type SequenceSearchReq struct {
	dto.Pagination `search:"-"`
	Source         int    `search:"-"`                                                                               // 数据库来源 0-total 1-ncbi 2-local
	Key            string `json:"key" form:"key"  search:"type:contains;column:sequence_description;table:sequence"` // 查询关键词
}

func (s *SequenceSearchReq) GetNeedSearch() interface{} {
	return *s
}

func (s *SequenceSearchReq) Transfer2Local() *SequenceLocalSearchReq {
	return &SequenceLocalSearchReq{
		Pagination: s.Pagination,
		Source:     s.Source,
		Key:        s.Key,
	}
}

type SequenceLocalSearchReq struct {
	dto.Pagination `search:"-"`
	Source         int    `search:"-"`                                                                                     // 数据库来源 0-total 1-ncbi 2-local
	Key            string `json:"key" form:"key"  search:"type:contains;column:sequence_description;table:sequence_local"` // 查询关键词
}

func (s *SequenceLocalSearchReq) GetNeedSearch() interface{} {
	return *s
}

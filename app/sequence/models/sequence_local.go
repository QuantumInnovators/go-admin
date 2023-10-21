package models

import (
	"go-admin/common/models"
)

type SequenceLocal struct {
	models.Model

	SequenceId          string `json:"sequenceId" gorm:"type:varchar(100);comment:SequenceId"`
	SequenceDescription string `json:"sequenceDescription" gorm:"type:varchar(255);comment:SequenceDescription"`
	Sequence            string `json:"sequence" gorm:"type:text;comment:Sequence"`
	models.ModelTime
	models.ControlBy
}

func (SequenceLocal) TableName() string {
	return "sequence_local"
}

func (e *SequenceLocal) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SequenceLocal) GetId() interface{} {
	return e.Id
}

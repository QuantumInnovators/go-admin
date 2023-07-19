package models

import (
	"go-admin/common/models"
)

type Sequence struct {
	models.Model

	SequenceId          string `json:"sequenceId" gorm:"type:varchar(100);comment:SequenceId"`
	SequenceDescription string `json:"sequenceDescription" gorm:"type:varchar(255);comment:SequenceDescription"`
	Sequence            string `json:"sequence" gorm:"type:text;comment:Sequence"`
	models.ModelTime
	models.ControlBy
}

func (Sequence) TableName() string {
	return "sequence"
}

func (e *Sequence) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Sequence) GetId() interface{} {
	return e.Id
}

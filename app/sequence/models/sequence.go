package models

import (
	"go-admin/common/models"
)

type Sequence struct {
	models.Model

	SequenceId          string `json:"sequenceId" gorm:"type:varchar(100);comment:SequenceId"`
	Name                string `json:"name" gorm:"type:varchar(255);comment:Name"`
	NameZh              string `json:"nameCN" gorm:"type:varchar(255);comment:中文名称"`
	SequenceDescription string `json:"sequenceDescription" gorm:"type:varchar(255);comment:SequenceDescription"`
	Sequence            string `json:"sequence" gorm:"type:text;comment:Sequence"`
	Type                string `json:"type" gorm:"type:varchar(255);comment:Type"`
	PrimerName          string `json:"primerName" gorm:"type:varchar(255);comment:PrimerName"`
	models.ModelTime
	models.ControlBy
}

func (Sequence) TableName() string {
	return "sequence"
}

func (Sequence) LocalTableName() string {
	return "sequence_local"
}

func (e *Sequence) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Sequence) GetId() interface{} {
	return e.Id
}

package models

import (

	"go-admin/common/models"

)

type Notation struct {
    models.Model
    
    Desc string `json:"desc" gorm:"type:text;comment:描述"` 
    models.ModelTime
    models.ControlBy
}

func (Notation) TableName() string {
    return "notation"
}

func (e *Notation) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Notation) GetId() interface{} {
	return e.Id
}
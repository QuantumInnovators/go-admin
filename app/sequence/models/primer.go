package models

import (

	"go-admin/common/models"

)

type Primer struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(255);comment:拉丁文名称"` 
    Desc string `json:"desc" gorm:"type:varchar(255);comment:描述"` 
    models.ModelTime
    models.ControlBy
}

func (Primer) TableName() string {
    return "primer"
}

func (e *Primer) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Primer) GetId() interface{} {
	return e.Id
}
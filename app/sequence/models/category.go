package models

import (

	"go-admin/common/models"

)

type Category struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(255);comment:拉丁文名称"` 
    Desc string `json:"desc" gorm:"type:varchar(255);comment:描述"` 
    models.ModelTime
    models.ControlBy
}

func (Category) TableName() string {
    return "category"
}

func (e *Category) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Category) GetId() interface{} {
	return e.Id
}
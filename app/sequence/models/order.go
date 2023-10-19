package models

import (

	"go-admin/common/models"

)

type Order struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(255);comment:拉丁文名"` 
    Desc string `json:"desc" gorm:"type:varchar(255);comment:Desc"` 
    ClassId string `json:"classId" gorm:"type:int;comment:所属纲"` 
    CategoryId string `json:"categoryId" gorm:"type:json;comment:所属类"` 
    models.ModelTime
    models.ControlBy
}

func (Order) TableName() string {
    return "order"
}

func (e *Order) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Order) GetId() interface{} {
	return e.Id
}
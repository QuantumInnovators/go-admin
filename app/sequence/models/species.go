package models

import (

	"go-admin/common/models"

)

type Species struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(255);comment:拉丁文名"` 
    Desc string `json:"desc" gorm:"type:varchar(255);comment:Desc"` 
    GenusId string `json:"genusId" gorm:"type:int;comment:所属属"` 
    CategoryId string `json:"categoryId" gorm:"type:json;comment:所属类"` 
    models.ModelTime
    models.ControlBy
}

func (Species) TableName() string {
    return "species"
}

func (e *Species) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Species) GetId() interface{} {
	return e.Id
}
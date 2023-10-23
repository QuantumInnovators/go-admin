package models

import (
	"go-admin/common/models"
)

type Class struct {
	models.Model

	Name       string `json:"name" gorm:"type:varchar(255);comment:拉丁文名"`
	Desc       string `json:"desc" gorm:"type:varchar(255);comment:Desc"`
	PhylumId   int    `json:"phylumId" gorm:"type:int;comment:所属门"`
	CategoryId string `json:"categoryId" gorm:"type:json;comment:所属类"`
	models.ModelTime
	models.ControlBy
}

func (Class) TableName() string {
	return "class"
}

func (e *Class) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Class) GetId() interface{} {
	return e.Id
}

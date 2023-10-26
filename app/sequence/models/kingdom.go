package models

import (
	"go-admin/common/models"
)

type Kingdom struct {
	models.Model

	Name       string `json:"name" gorm:"type:varchar(255);comment:拉丁文名"`
	Desc       string `json:"desc" gorm:"type:varchar(255);comment:中文名"`
	CategoryId string `json:"categoryId" gorm:"type:varchar(255);comment:所属类"`
	models.ModelTime
	models.ControlBy
}

func (Kingdom) TableName() string {
	return "kingdom"
}

func (e *Kingdom) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Kingdom) GetId() interface{} {
	return e.Id
}

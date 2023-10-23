package models

import (
	"go-admin/common/models"
)

type Class struct {
	Id         int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Name       string `json:"name" gorm:"type:varchar(255);comment:拉丁文名"`
	Desc       string `json:"desc" gorm:"type:varchar(255);comment:Desc"`
	PhylumId   string `json:"phylumId" gorm:"type:int;comment:所属门"`
	CategoryId int    `json:"categoryId" gorm:"type:int;comment:所属类"`
	models.ModelTime
	models.ControlBy
}

func (Class) TableName() string {
	return "class"
}

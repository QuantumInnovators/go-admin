package models

import (
	"go-admin/common/models"
)

type Family struct {
	Id         int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Name       string `json:"name" gorm:"type:varchar(255);comment:拉丁文名"`
	NameCN     string `json:"nameCN" gorm:"type:varchar(255);comment:中文名"`
	Desc       string `json:"desc" gorm:"type:varchar(255);comment:Desc"`
	ParentId   int    `json:"parentId" gorm:"type:bigint;comment:所属目"`
	CategoryId string `json:"categoryId" gorm:"type:varchar(255);comment:所属类"`
	models.ModelTime
	models.ControlBy
}

func (Family) TableName() string {
	return "family"
}

package models

import (
	"go-admin/common/models"
)

type Genus struct {
	models.Model

	Name       string `json:"name" gorm:"type:varchar(255);comment:拉丁文名"`
	NameCn     string `json:"nameCn" gorm:"type:varchar(255);comment:中文名"`
	Desc       string `json:"desc" gorm:"type:varchar(255);comment:Desc"`
	ParentId   int64  `json:"parentId" gorm:"type:bigint;comment:所属科"`
	CategoryId string `json:"categoryId" gorm:"type:varchar(255);comment:所属类"`
	models.ModelTime
	models.ControlBy
}

func (Genus) TableName() string {
	return "genus"
}

func (e *Genus) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Genus) GetId() interface{} {
	return e.Id
}

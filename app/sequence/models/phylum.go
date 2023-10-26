package models

import (
	"go-admin/common/models"
)

type Phylum struct {
	models.Model

	Name       string `json:"name" gorm:"type:varchar(255);comment:拉丁文名"`
	NameCn     string `json:"nameCn" gorm:"type:varchar(255);comment:中文名"`
	Desc       string `json:"desc" gorm:"type:varchar(255);comment:Desc"`
	ParentId   int64  `json:"parentId" gorm:"type:bigint;comment:所属界"`
	CategoryId string `json:"categoryId" gorm:"type:varchar(255);comment:所属类"`
	models.ModelTime
	models.ControlBy
}

func (Phylum) TableName() string {
	return "phylum"
}

func (e *Phylum) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Phylum) GetId() interface{} {
	return e.Id
}

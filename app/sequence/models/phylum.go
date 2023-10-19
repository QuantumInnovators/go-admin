package models

import (
	"go-admin/common/models"
)

type Phylum struct {
	models.Model

	Name       string `json:"name" gorm:"type:varchar(255);comment:拉丁文名"`
	Desc       string `json:"desc" gorm:"type:varchar(255);comment:Desc"`
	KingdomId  int    `json:"kingdomId" gorm:"type:int;comment:所属界"`
	CategoryId string `json:"categoryId" gorm:"type:json;comment:所属种群"`
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

package models

import (
	"go-admin/common/models"
)

type Phylum struct {
	Id         int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Name       string `json:"name" gorm:"type:varchar(255);comment:拉丁文名"`
	Desc       string `json:"desc" gorm:"type:varchar(255);comment:Desc"`
	KingdomId  int    `json:"kingdomId" gorm:"type:int;comment:所属界"`
	CategoryId int    `json:"categoryId" gorm:"type:int;comment:所属类"`
	models.ModelTime
	models.ControlBy
}

func (Phylum) TableName() string {
	return "phylum"
}

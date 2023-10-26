package models

import (
	"go-admin/common/models"
)

type Meta struct {
	Id         int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	CategoryId string `json:"categoryId" gorm:"type:bigint;comment:类别"`
	KingdomId  string `json:"kingdomId" gorm:"type:bigint;comment:界"`
	PhylumId   string `json:"phylumId" gorm:"type:bigint;comment:门"`
	ClassId    string `json:"classId" gorm:"type:bigint;comment:纲"`
	OrderId    string `json:"orderId" gorm:"type:bigint;comment:目"`
	FamilyId   string `json:"familyId" gorm:"type:bigint;comment:科"`
	GenusId    string `json:"genusId" gorm:"type:bigint;comment:属"`
	SpeciesId  string `json:"speciesId" gorm:"type:bigint;comment:种"`
	models.ModelTime
	models.ControlBy
}

func (Meta) TableName() string {
	return "meta"
}

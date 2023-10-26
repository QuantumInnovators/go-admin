package models

import (
	"go-admin/common/models"
)

type Meta struct {
	models.Model

	CategoryId   int64  `json:"categoryId" gorm:"type:bigint;comment:类别"`
	CategoryName string `json:"categoryName" gorm:"type:varchar(255);comment:类别"`
	KingdomId    int64  `json:"kingdomId" gorm:"type:bigint;comment:界"`
	KingdomName  string `json:"kingdomName" gorm:"type:varchar(255);comment:界"`
	PhylumId     int64  `json:"phylumId" gorm:"type:bigint;comment:门"`
	PhylumName   string `json:"phylumName" gorm:"type:varchar(255);comment:门"`
	ClassId      int64  `json:"classId" gorm:"type:bigint;comment:纲"`
	ClassName    string `json:"className" gorm:"type:varchar(255);comment:纲"`
	OrderId      int64  `json:"orderId" gorm:"type:bigint;comment:目"`
	OrderName    string `json:"orderName" gorm:"type:varchar(255);comment:目"`
	FamilyId     int64  `json:"familyId" gorm:"type:bigint;comment:科"`
	FamilyName   string `json:"familyName" gorm:"type:varchar(255);comment:科"`
	GenusId      int64  `json:"genusId" gorm:"type:bigint;comment:属"`
	GenusName    string `json:"genusName" gorm:"type:varchar(255);comment:属"`
	SpeciesId    int64  `json:"speciesId" gorm:"type:bigint;comment:种"`
	SpeciesName  string `json:"speciesName" gorm:"type:varchar(255);comment:种"`
	models.ModelTime
	models.ControlBy
}

func (Meta) TableName() string {
	return "meta"
}

func (e *Meta) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Meta) GetId() interface{} {
	return e.Id
}

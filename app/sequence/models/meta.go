package models

import (
	"go-admin/common/models"
)

type Meta struct {
	models.Model

	CategoryId  string `json:"categoryId" gorm:"type:int;comment:类别"`
	KingdomId   string `json:"kingdomId" gorm:"type:int;comment:界"`
	KingdomName string `json:"kingdomName" gorm:"type:varchar(255);comment:界"`
	PhylumId    string `json:"phylumId" gorm:"type:int;comment:门"`
	PhylumName  string `json:"phylumName" gorm:"type:varchar(255);comment:门"`
	ClassId     string `json:"classId" gorm:"type:int;comment:纲"`
	ClassName   string `json:"className" gorm:"type:varchar(255);comment:纲"`
	OrderId     string `json:"orderId" gorm:"type:int;comment:目"`
	OrderName   string `json:"orderName" gorm:"type:varchar(255);comment:目"`
	FamilyId    string `json:"familyId" gorm:"type:int;comment:科"`
	FamilyName  string `json:"familyName" gorm:"type:varchar(255);comment:科"`
	GenusId     string `json:"genusId" gorm:"type:int;comment:属"`
	GenusName   string `json:"genusName" gorm:"type:varchar(255);comment:属"`
	SpeciesId   string `json:"speciesId" gorm:"type:int;comment:种"`
	SpeciesName string `json:"speciesName" gorm:"type:varchar(255);comment:种"`
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

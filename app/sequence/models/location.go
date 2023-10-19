package models

import (

	"go-admin/common/models"

)

type Location struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(255);comment:地点简称"` 
    models.ModelTime
    models.ControlBy
}

func (Location) TableName() string {
    return "location"
}

func (e *Location) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Location) GetId() interface{} {
	return e.Id
}
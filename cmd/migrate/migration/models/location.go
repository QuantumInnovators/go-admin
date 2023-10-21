package models

import (
	"go-admin/common/models"
)

type Location struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Name string `json:"name" gorm:"type:varchar(255);comment:地点简称"`
	models.ModelTime
	models.ControlBy
}

func (Location) TableName() string {
	return "location"
}

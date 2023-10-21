package models

import (
	"go-admin/common/models"
)

type Notation struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Desc string `json:"desc" gorm:"type:text;comment:描述"`
	models.ModelTime
	models.ControlBy
}

func (Notation) TableName() string {
	return "notation"
}

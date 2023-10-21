package models

import (
	"go-admin/common/models"
)

type Category struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Name string `json:"name" gorm:"type:varchar(255);comment:拉丁文名称"`
	Desc string `json:"desc" gorm:"type:varchar(255);comment:描述"`
	models.ModelTime
	models.ControlBy
}

func (Category) TableName() string {
	return "category"
}

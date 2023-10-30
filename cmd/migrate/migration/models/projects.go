package models

import (
	"go-admin/common/models"
)

type Projects struct {
	Id             int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	FolderName     string `json:"folderName" gorm:"type:varchar(255);comment:目录名称"`
	IsProject      bool   `json:"isProject" gorm:"type:tinyint(1);comment:是否是项目名称"`
	IsFolder       bool   `json:"isFolder" gorm:"type:tinyint(1);comment:是否为目录"`
	Path           string `json:"path" gorm:"type:varchar(255);comment:文件路径"`
	ParentFolderId int    `json:"parentFolderId" gorm:"type:int;comment:目录id"`
	models.ModelTime
	models.ControlBy
}

func (Projects) TableName() string {
	return "projects"
}

package models

import (

	"go-admin/common/models"

)

type Projects struct {
    models.Model
    
    FolderName string `json:"folderName" gorm:"type:varchar(255);comment:目录名称"` 
    IsProject int64 `json:"isProject" gorm:"type:tinyint(1);comment:是否是项目名称"` 
    IsFolder int64 `json:"isFolder" gorm:"type:tinyint(1);comment:是否为目录"` 
    Path string `json:"path" gorm:"type:varchar(255);comment:文件路径"` 
    ParentFolderId int64 `json:"parentFolderId" gorm:"type:int;comment:目录id"` 
    models.ModelTime
    models.ControlBy
}

func (Projects) TableName() string {
    return "projects"
}

func (e *Projects) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Projects) GetId() interface{} {
	return e.Id
}
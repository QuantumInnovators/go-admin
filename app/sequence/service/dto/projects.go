package dto

import (

	"go-admin/app/sequence/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ProjectsGetPageReq struct {
	dto.Pagination     `search:"-"`
    ProjectsOrder
}

type ProjectsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:projects"`
    FolderName string `form:"folderNameOrder"  search:"type:order;column:folder_name;table:projects"`
    IsProject string `form:"isProjectOrder"  search:"type:order;column:is_project;table:projects"`
    IsFolder string `form:"isFolderOrder"  search:"type:order;column:is_folder;table:projects"`
    Path string `form:"pathOrder"  search:"type:order;column:path;table:projects"`
    ParentFolderId string `form:"parentFolderIdOrder"  search:"type:order;column:parent_folder_id;table:projects"`
    
}

func (m *ProjectsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ProjectsInsertReq struct {
    Id int `json:"-" comment:""` // 
    FolderName string `json:"folderName" comment:"目录名称"`
    IsProject int64 `json:"isProject" comment:"是否是项目名称"`
    IsFolder int64 `json:"isFolder" comment:"是否为目录"`
    Path string `json:"path" comment:"文件路径"`
    ParentFolderId int64 `json:"parentFolderId" comment:"目录id"`
    common.ControlBy
}

func (s *ProjectsInsertReq) Generate(model *models.Projects)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.FolderName = s.FolderName
    model.IsProject = s.IsProject
    model.IsFolder = s.IsFolder
    model.Path = s.Path
    model.ParentFolderId = s.ParentFolderId
}

func (s *ProjectsInsertReq) GetId() interface{} {
	return s.Id
}

type ProjectsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    FolderName string `json:"folderName" comment:"目录名称"`
    IsProject int64 `json:"isProject" comment:"是否是项目名称"`
    IsFolder int64 `json:"isFolder" comment:"是否为目录"`
    Path string `json:"path" comment:"文件路径"`
    ParentFolderId int64 `json:"parentFolderId" comment:"目录id"`
    common.ControlBy
}

func (s *ProjectsUpdateReq) Generate(model *models.Projects)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.FolderName = s.FolderName
    model.IsProject = s.IsProject
    model.IsFolder = s.IsFolder
    model.Path = s.Path
    model.ParentFolderId = s.ParentFolderId
}

func (s *ProjectsUpdateReq) GetId() interface{} {
	return s.Id
}

// ProjectsGetReq 功能获取请求参数
type ProjectsGetReq struct {
     Id int `uri:"id"`
}
func (s *ProjectsGetReq) GetId() interface{} {
	return s.Id
}

// ProjectsDeleteReq 功能删除请求参数
type ProjectsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ProjectsDeleteReq) GetId() interface{} {
	return s.Ids
}

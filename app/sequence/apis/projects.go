package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/sequence/models"
	"go-admin/app/sequence/service"
	"go-admin/app/sequence/service/dto"
	"go-admin/common/actions"
)

type Projects struct {
	api.Api
}

// GetPage 获取Projects列表
// @Summary 获取Projects列表
// @Description 获取Projects列表
// @Tags Projects
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Projects}} "{"code": 200, "data": [...]}"
// @Router /api/v1/projects [get]
// @Security Bearer
func (e Projects) GetPage(c *gin.Context) {
    req := dto.ProjectsGetPageReq{}
    s := service.Projects{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.Projects, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Projects失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Projects
// @Summary 获取Projects
// @Description 获取Projects
// @Tags Projects
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Projects} "{"code": 200, "data": [...]}"
// @Router /api/v1/projects/{id} [get]
// @Security Bearer
func (e Projects) Get(c *gin.Context) {
	req := dto.ProjectsGetReq{}
	s := service.Projects{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Projects

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Projects失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建Projects
// @Summary 创建Projects
// @Description 创建Projects
// @Tags Projects
// @Accept application/json
// @Product application/json
// @Param data body dto.ProjectsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/projects [post]
// @Security Bearer
func (e Projects) Insert(c *gin.Context) {
    req := dto.ProjectsInsertReq{}
    s := service.Projects{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建Projects失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Projects
// @Summary 修改Projects
// @Description 修改Projects
// @Tags Projects
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ProjectsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/projects/{id} [put]
// @Security Bearer
func (e Projects) Update(c *gin.Context) {
    req := dto.ProjectsUpdateReq{}
    s := service.Projects{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改Projects失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除Projects
// @Summary 删除Projects
// @Description 删除Projects
// @Tags Projects
// @Param data body dto.ProjectsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/projects [delete]
// @Security Bearer
func (e Projects) Delete(c *gin.Context) {
    s := service.Projects{}
    req := dto.ProjectsDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除Projects失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}

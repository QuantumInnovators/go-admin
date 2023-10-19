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

type Notation struct {
	api.Api
}

// GetPage 获取Notation列表
// @Summary 获取Notation列表
// @Description 获取Notation列表
// @Tags Notation
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Notation}} "{"code": 200, "data": [...]}"
// @Router /api/v1/notation [get]
// @Security Bearer
func (e Notation) GetPage(c *gin.Context) {
    req := dto.NotationGetPageReq{}
    s := service.Notation{}
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
	list := make([]models.Notation, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Notation失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Notation
// @Summary 获取Notation
// @Description 获取Notation
// @Tags Notation
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Notation} "{"code": 200, "data": [...]}"
// @Router /api/v1/notation/{id} [get]
// @Security Bearer
func (e Notation) Get(c *gin.Context) {
	req := dto.NotationGetReq{}
	s := service.Notation{}
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
	var object models.Notation

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Notation失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建Notation
// @Summary 创建Notation
// @Description 创建Notation
// @Tags Notation
// @Accept application/json
// @Product application/json
// @Param data body dto.NotationInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/notation [post]
// @Security Bearer
func (e Notation) Insert(c *gin.Context) {
    req := dto.NotationInsertReq{}
    s := service.Notation{}
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
		e.Error(500, err, fmt.Sprintf("创建Notation失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Notation
// @Summary 修改Notation
// @Description 修改Notation
// @Tags Notation
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.NotationUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/notation/{id} [put]
// @Security Bearer
func (e Notation) Update(c *gin.Context) {
    req := dto.NotationUpdateReq{}
    s := service.Notation{}
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
		e.Error(500, err, fmt.Sprintf("修改Notation失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除Notation
// @Summary 删除Notation
// @Description 删除Notation
// @Tags Notation
// @Param data body dto.NotationDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/notation [delete]
// @Security Bearer
func (e Notation) Delete(c *gin.Context) {
    s := service.Notation{}
    req := dto.NotationDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Notation失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}

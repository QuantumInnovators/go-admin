package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"

	"go-admin/app/sequence/models"
	"go-admin/app/sequence/service"
	"go-admin/app/sequence/service/dto"
	"go-admin/common/actions"
)

type Meta struct {
	api.Api
}

// GetPage 获取Meta列表
// @Summary 获取Meta列表
// @Description 获取Meta列表
// @Tags Meta
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Meta}} "{"code": 200, "data": [...]}"
// @Router /api/v1/meta [get]
// @Security Bearer
func (e Meta) GetPage(c *gin.Context) {
	req := dto.MetaGetPageReq{}
	s := service.Meta{}
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
	list := make([]models.Meta, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Meta失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Meta
// @Summary 获取Meta
// @Description 获取Meta
// @Tags Meta
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Meta} "{"code": 200, "data": [...]}"
// @Router /api/v1/meta/{id} [get]
// @Security Bearer
func (e Meta) Get(c *gin.Context) {
	req := dto.MetaGetReq{}
	s := service.Meta{}
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
	var object models.Meta

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Meta失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Meta
// @Summary 创建Meta
// @Description 创建Meta
// @Tags Meta
// @Accept application/json
// @Product application/json
// @Param data body dto.MetaInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/meta [post]
// @Security Bearer
func (e Meta) Insert(c *gin.Context) {
	req := dto.MetaInsertReq{}
	s := service.Meta{}
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
		e.Error(500, err, fmt.Sprintf("创建Meta失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Meta
// @Summary 修改Meta
// @Description 修改Meta
// @Tags Meta
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.MetaUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/meta/{id} [put]
// @Security Bearer
func (e Meta) Update(c *gin.Context) {
	req := dto.MetaUpdateReq{}
	s := service.Meta{}
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
		e.Error(500, err, fmt.Sprintf("修改Meta失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Meta
// @Summary 删除Meta
// @Description 删除Meta
// @Tags Meta
// @Param data body dto.MetaDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/meta [delete]
// @Security Bearer
func (e Meta) Delete(c *gin.Context) {
	s := service.Meta{}
	req := dto.MetaDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Meta失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

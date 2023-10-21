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

type Primer struct {
	api.Api
}

// GetPage 获取Primer列表
// @Summary 获取Primer列表
// @Description 获取Primer列表
// @Tags Primer
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Primer}} "{"code": 200, "data": [...]}"
// @Router /api/v1/primer [get]
// @Security Bearer
func (e Primer) GetPage(c *gin.Context) {
	req := dto.PrimerGetPageReq{}
	s := service.Primer{}
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
	list := make([]models.Primer, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Primer失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Primer
// @Summary 获取Primer
// @Description 获取Primer
// @Tags Primer
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Primer} "{"code": 200, "data": [...]}"
// @Router /api/v1/primer/{id} [get]
// @Security Bearer
func (e Primer) Get(c *gin.Context) {
	req := dto.PrimerGetReq{}
	s := service.Primer{}
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
	var object models.Primer

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Primer失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Primer
// @Summary 创建Primer
// @Description 创建Primer
// @Tags Primer
// @Accept application/json
// @Product application/json
// @Param data body dto.PrimerInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/primer [post]
// @Security Bearer
func (e Primer) Insert(c *gin.Context) {
	req := dto.PrimerInsertReq{}
	s := service.Primer{}
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
		e.Error(500, err, fmt.Sprintf("创建Primer失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Primer
// @Summary 修改Primer
// @Description 修改Primer
// @Tags Primer
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.PrimerUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/primer/{id} [put]
// @Security Bearer
func (e Primer) Update(c *gin.Context) {
	req := dto.PrimerUpdateReq{}
	s := service.Primer{}
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
		e.Error(500, err, fmt.Sprintf("修改Primer失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Primer
// @Summary 删除Primer
// @Description 删除Primer
// @Tags Primer
// @Param data body dto.PrimerDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/primer [delete]
// @Security Bearer
func (e Primer) Delete(c *gin.Context) {
	s := service.Primer{}
	req := dto.PrimerDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Primer失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

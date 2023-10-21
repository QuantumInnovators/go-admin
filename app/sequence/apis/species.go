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

type Species struct {
	api.Api
}

// GetPage 获取Species列表
// @Summary 获取Species列表
// @Description 获取Species列表
// @Tags Species
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Species}} "{"code": 200, "data": [...]}"
// @Router /api/v1/species [get]
// @Security Bearer
func (e Species) GetPage(c *gin.Context) {
	req := dto.SpeciesGetPageReq{}
	s := service.Species{}
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
	list := make([]models.Species, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Species失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Species
// @Summary 获取Species
// @Description 获取Species
// @Tags Species
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Species} "{"code": 200, "data": [...]}"
// @Router /api/v1/species/{id} [get]
// @Security Bearer
func (e Species) Get(c *gin.Context) {
	req := dto.SpeciesGetReq{}
	s := service.Species{}
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
	var object models.Species

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Species失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Species
// @Summary 创建Species
// @Description 创建Species
// @Tags Species
// @Accept application/json
// @Product application/json
// @Param data body dto.SpeciesInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/species [post]
// @Security Bearer
func (e Species) Insert(c *gin.Context) {
	req := dto.SpeciesInsertReq{}
	s := service.Species{}
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
		e.Error(500, err, fmt.Sprintf("创建Species失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Species
// @Summary 修改Species
// @Description 修改Species
// @Tags Species
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SpeciesUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/species/{id} [put]
// @Security Bearer
func (e Species) Update(c *gin.Context) {
	req := dto.SpeciesUpdateReq{}
	s := service.Species{}
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
		e.Error(500, err, fmt.Sprintf("修改Species失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Species
// @Summary 删除Species
// @Description 删除Species
// @Tags Species
// @Param data body dto.SpeciesDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/species [delete]
// @Security Bearer
func (e Species) Delete(c *gin.Context) {
	s := service.Species{}
	req := dto.SpeciesDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Species失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

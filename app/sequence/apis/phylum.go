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

type Phylum struct {
	api.Api
}

// GetPage 获取Phylum列表
// @Summary 获取Phylum列表
// @Description 获取Phylum列表
// @Tags Phylum
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Phylum}} "{"code": 200, "data": [...]}"
// @Router /api/v1/phylum [get]
// @Security Bearer
func (e Phylum) GetPage(c *gin.Context) {
	req := dto.PhylumGetPageReq{}
	s := service.Phylum{}
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
	list := make([]models.Phylum, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Phylum失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Phylum
// @Summary 获取Phylum
// @Description 获取Phylum
// @Tags Phylum
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Phylum} "{"code": 200, "data": [...]}"
// @Router /api/v1/phylum/{id} [get]
// @Security Bearer
func (e Phylum) Get(c *gin.Context) {
	req := dto.PhylumGetReq{}
	s := service.Phylum{}
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
	var object models.Phylum

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Phylum失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Phylum
// @Summary 创建Phylum
// @Description 创建Phylum
// @Tags Phylum
// @Accept application/json
// @Product application/json
// @Param data body dto.PhylumInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/phylum [post]
// @Security Bearer
func (e Phylum) Insert(c *gin.Context) {
	req := dto.PhylumInsertReq{}
	s := service.Phylum{}
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
		e.Error(500, err, fmt.Sprintf("创建Phylum失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Phylum
// @Summary 修改Phylum
// @Description 修改Phylum
// @Tags Phylum
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.PhylumUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/phylum/{id} [put]
// @Security Bearer
func (e Phylum) Update(c *gin.Context) {
	req := dto.PhylumUpdateReq{}
	s := service.Phylum{}
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
		e.Error(500, err, fmt.Sprintf("修改Phylum失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Phylum
// @Summary 删除Phylum
// @Description 删除Phylum
// @Tags Phylum
// @Param data body dto.PhylumDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/phylum [delete]
// @Security Bearer
func (e Phylum) Delete(c *gin.Context) {
	s := service.Phylum{}
	req := dto.PhylumDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Phylum失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

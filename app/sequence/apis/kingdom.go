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

type Kingdom struct {
	api.Api
}

// GetPage 获取Kingdom列表
// @Summary 获取Kingdom列表
// @Description 获取Kingdom列表
// @Tags Kingdom
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Kingdom}} "{"code": 200, "data": [...]}"
// @Router /api/v1/kingdom [get]
// @Security Bearer
func (e Kingdom) GetPage(c *gin.Context) {
	req := dto.KingdomGetPageReq{}
	s := service.Kingdom{}
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
	list := make([]models.Kingdom, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Kingdom失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Kingdom
// @Summary 获取Kingdom
// @Description 获取Kingdom
// @Tags Kingdom
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Kingdom} "{"code": 200, "data": [...]}"
// @Router /api/v1/kingdom/{id} [get]
// @Security Bearer
func (e Kingdom) Get(c *gin.Context) {
	req := dto.KingdomGetReq{}
	s := service.Kingdom{}
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
	var object models.Kingdom

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Kingdom失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Kingdom
// @Summary 创建Kingdom
// @Description 创建Kingdom
// @Tags Kingdom
// @Accept application/json
// @Product application/json
// @Param data body dto.KingdomInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/kingdom [post]
// @Security Bearer
func (e Kingdom) Insert(c *gin.Context) {
	req := dto.KingdomInsertReq{}
	s := service.Kingdom{}
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
		e.Error(500, err, fmt.Sprintf("创建Kingdom失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Kingdom
// @Summary 修改Kingdom
// @Description 修改Kingdom
// @Tags Kingdom
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.KingdomUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/kingdom/{id} [put]
// @Security Bearer
func (e Kingdom) Update(c *gin.Context) {
	req := dto.KingdomUpdateReq{}
	s := service.Kingdom{}
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
		e.Error(500, err, fmt.Sprintf("修改Kingdom失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Kingdom
// @Summary 删除Kingdom
// @Description 删除Kingdom
// @Tags Kingdom
// @Param data body dto.KingdomDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/kingdom [delete]
// @Security Bearer
func (e Kingdom) Delete(c *gin.Context) {
	s := service.Kingdom{}
	req := dto.KingdomDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Kingdom失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

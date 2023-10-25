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

	for idx, item := range list {
		// 从 kingdom 表里查询 kingdom 的名称
		var kingdom models.Kingdom
		e.Orm.Model(&kingdom).Where("id = ?", item.KingdomId).First(&kingdom)
		list[idx].KingdomName = kingdom.Name
		// 从 phylum 表里查询 phylum 的名称
		var phylum models.Phylum
		e.Orm.Model(&phylum).Where("id = ?", item.PhylumId).First(&phylum)
		list[idx].PhylumName = phylum.Name
		// 从 class 表里查询 class 的名称
		var class models.Class
		e.Orm.Model(&class).Where("id = ?", item.ClassId).First(&class)
		list[idx].ClassName = class.Name
		// 从 order 表里查询 order 的名称
		var order models.Order
		e.Orm.Model(&order).Where("id = ?", item.OrderId).First(&order)
		list[idx].OrderName = order.Name
		// 从 family 表里查询 family 的名称
		var family models.Family
		e.Orm.Model(&family).Where("id = ?", item.FamilyId).First(&family)
		list[idx].FamilyName = family.Name
		// 从 genus 表里查询 genus 的名称
		var genus models.Genus
		e.Orm.Model(&genus).Where("id = ?", item.GenusId).First(&genus)
		list[idx].GenusName = genus.Name
		// 从 species 表里查询 species 的名称
		var species models.Species
		e.Orm.Model(&species).Where("id = ?", item.SpeciesId).First(&species)
		list[idx].SpeciesName = species.Name
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

package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"strings"

	"go-admin/app/sequence/models"
	"go-admin/app/sequence/service"
	"go-admin/app/sequence/service/dto"
	"go-admin/common/actions"
)

type Sequence struct {
	api.Api
}

// GetPage 获取Sequence列表
// @Summary 获取Sequence列表
// @Description 获取Sequence列表
// @Tags Sequence
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sequence [get]
// @Security Bearer
func (e Sequence) GetPage(c *gin.Context) {
	req := dto.SequenceGetPageReq{}
	s := service.Sequence{}
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
	list := make([]models.Sequence, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Sequence失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Sequence
// @Summary 获取Sequence
// @Description 获取Sequence
// @Tags Sequence
// @Param id path int false "id"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sequence/{id} [get]
// @Security Bearer
func (e Sequence) Get(c *gin.Context) {
	req := dto.SequenceGetReq{}
	s := service.Sequence{}
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
	var object models.Sequence

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Sequence失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Sequence
// @Summary 创建Sequence
// @Description 创建Sequence
// @Tags Sequence
// @Accept application/json
// @Product application/json
// @Param data body dto.SequenceInsertReq true "data"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sequence [post]
// @Security Bearer
func (e Sequence) Insert(c *gin.Context) {
	req := dto.SequenceInsertReq{}
	s := service.Sequence{}
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
		e.Error(500, err, fmt.Sprintf("创建Sequence失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Sequence
// @Summary 修改Sequence
// @Description 修改Sequence
// @Tags Sequence
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SequenceUpdateReq true "body"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sequence/{id} [put]
// @Security Bearer
func (e Sequence) Update(c *gin.Context) {
	req := dto.SequenceUpdateReq{}
	s := service.Sequence{}
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
		e.Error(500, err, fmt.Sprintf("修改Sequence失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Sequence
// @Summary 删除Sequence
// @Description 删除Sequence
// @Tags Sequence
// @Param data body dto.SequenceDeleteReq true "body"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sequence [delete]
// @Security Bearer
func (e Sequence) Delete(c *gin.Context) {
	s := service.Sequence{}
	req := dto.SequenceDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Sequence失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// GetByClassID
// @Summary 获取Sequence
// @Description 获取Sequence
// @Tags Sequence
// @Param class_id path int false "class_id"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sequence/class/{class_id} [get]
// @Security Bearer
func (e Sequence) GetByClassID(c *gin.Context) {
	req := dto.SequenceGetByClassReq{}
	s := service.Sequence{}
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
	var object models.Sequence

	p := actions.GetPermissionFromContext(c)
	err = s.GetByClass(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Sequence失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Search
// @Summary 查询Sequence信息
// @Description 查询Sequence信息
// @Tags Sequence
// @Param data body object true "查询条件"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sequence/search [post]
// @Security Bearer
func (e Sequence) Search(c *gin.Context) {
	req := dto.SequenceSearchReq{}
	s := service.Sequence{}
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
	type packData struct {
		Source int               `json:"source"`
		List   []models.Sequence `json:"list"`
	}
	list := make([]packData, 0)
	var count int64

	var data models.Sequence
	var dataLocal models.SequenceLocal
	var searchTables []string
	var iData []interface{}
	switch req.Source {
	case 0:
		searchTables = append(searchTables, data.TableName())
		iData = append(iData, data)
		list = append(list, packData{Source: 1, List: make([]models.Sequence, 0)})
		searchTables = append(searchTables, dataLocal.TableName())
		iData = append(iData, dataLocal)
		list = append(list, packData{Source: 2, List: make([]models.Sequence, 0)})
	case 1:
		searchTables = append(searchTables, data.TableName())
		iData = append(iData, data)
		list = append(list, packData{Source: 1, List: make([]models.Sequence, 0)})
	case 2:
		searchTables = append(searchTables, dataLocal.TableName())
		iData = append(iData, dataLocal)
		list = append(list, packData{Source: 2, List: make([]models.Sequence, 0)})
	}

	for idx, tableName := range searchTables {
		retList := make([]models.Sequence, 0)
		err = s.GetFromSourceByKey(&req, p, &retList, &count, iData[idx], tableName)
		if err != nil {
			e.Error(500, err, fmt.Sprintf("获取Sequence失败，\r\n失败信息 %s", err.Error()))
			return
		}

		// 处理一下retList数据
		for idx, item := range retList {
			if item.Name != "" {
				continue
			}
			// 按照一定格式从desc字符串转换
			// Caenidae sp. Cy2020 sp1 voucher Caenidae_Cy2020_sp1 small subunit ribosomal RNA gene and internal transcribed spacer 1, partial sequence
			// 获取 Caenidae sp. 为 name
			// 获取 partial 及以后为 type
			var desc string
			index := strings.Index(item.SequenceDescription, "partial")
			if index != -1 {
				retList[idx].Type = item.SequenceDescription[index:]
				desc = item.SequenceDescription[:index]
			}
			// 查询所有种的名字，如果在desc中存在，则提取出来
			var speciesList []models.Species
			e.Orm.Model(&models.Species{}).Find(&speciesList)
			for _, species := range speciesList {
				if strings.Contains(desc, species.Name) {
					retList[idx].Name = species.Name
					retList[idx].NameZh = species.Desc
					break
				}
			}
			if retList[idx].Name != "" {
				continue
			}
			// 查询所有的属名字，如果在desc中存在，则提取出来
			var genusList []models.Genus
			e.Orm.Model(&models.Genus{}).Find(&genusList)
			for _, genus := range genusList {
				if strings.Contains(desc, genus.Name) {
					retList[idx].Name = genus.Name
					retList[idx].NameZh = genus.Desc
					break
				}
			}
			if retList[idx].Name != "" {
				continue
			}
			// 查询所有的科名字，如果在desc中存在，则提取出来
			var familyList []models.Family
			e.Orm.Model(&models.Family{}).Find(&familyList)
			for _, family := range familyList {
				if strings.Contains(desc, family.Name) {
					retList[idx].Name = family.Name
					retList[idx].NameZh = family.Desc
					break
				}
			}
		}
		if retList != nil {
			list[idx].List = retList
		}
	}
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

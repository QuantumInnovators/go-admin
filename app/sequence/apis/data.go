package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/sequence/models"
	"go-admin/app/sequence/service"
	"go-admin/app/sequence/service/dto"
	"go-admin/common/actions"
)

type Data struct {
	api.Api
}

// GetTotalClassList .
func (e Data) GetTotalClassList(c *gin.Context) {
	req := dto.GetTotalClassListReq{}
	s := service.ClassList{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := models.ClassList{}
	err = s.GetTotalClassList(p, &list)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(list, "查询成功")
}

package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-admin/common/middleware"

	"go-admin/app/sequence/apis"
	"go-admin/common/actions"
	_ "go-admin/docs/sequence"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSequenceRouter)
}

// registerSequenceRouter
func registerSequenceRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Sequence{}
	r := v1.Group("/sequence").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.GET("/class/:classId", actions.PermissionAction(), api.GetByClassID)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
		r.POST("/uploadFile", actions.PermissionAction(), api.UploadFile)
		r.POST("/search", actions.PermissionAction(), api.Search)
	}
}

func swaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/sequence/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("sequence")))
}

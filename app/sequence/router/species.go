package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/sequence/apis"
	"go-admin/common/middleware"
	"go-admin/common/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSpeciesRouter)
}

// registerSpeciesRouter
func registerSpeciesRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Species{}
	r := v1.Group("/species").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}
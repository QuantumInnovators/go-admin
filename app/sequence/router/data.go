package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/sequence/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDataRouter)
}

// registerSequenceRouter
func registerDataRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Data{}
	r := v1.Group("/logic").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/classList", api.GetTotalClassList)
	}
}

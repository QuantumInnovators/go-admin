// Package api .
package api

import "go-admin/app/sequence/router"

func init() {
	AppRouters = append(AppRouters, router.InitRouter)
}

/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:15
 * @File  : router.go
 * @Description: ...
 */
package router

import (
	"github.com/gin-gonic/gin"
	"go-lion/middleware"
)

var Router *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.New()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
}

func SetRouter() {
	Router.Use(middleware.CORS())
}

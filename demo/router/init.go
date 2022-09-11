package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"demo/config"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	corsConfig := cors.DefaultConfig()
	//同源策略, cors中间件鉴权登录
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	if config.Config.Dev {
		//开发环境下, 绕过鉴权
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = config.Config.Server.AllowOrigins
	}
	Router.Use(cors.New(corsConfig))
	SetRouter()
}

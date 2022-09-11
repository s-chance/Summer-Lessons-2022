package router

import (
	"demo/controller"
	"demo/utility/middleware"
)

func SetRouter() {
	Router.POST("/api/user/register", controller.Register)
	Router.POST("/api/user/login", controller.Login)
	Router.GET("/api/user", middleware.Authorization, controller.GetUserInfo)
	Router.POST("/api/class", middleware.Authorization, controller.AddClass)
	Router.DELETE("/api/class", middleware.Authorization, controller.DeleteClass)
	Router.PUT("/api/class", middleware.Authorization, controller.UpdateClass)
	Router.GET("/api/class", middleware.Authorization, controller.GetClasses)
}

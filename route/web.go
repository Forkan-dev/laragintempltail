package route

import (
	"project1/app/http/controller"
	"project1/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func SetWebRoutes(router *gin.Engine) {
	web := router.Group("/")
	{
		// Routes requiring authentication
		auth := web.Group("/")
		auth.Use(middleware.Auth)
		{
			auth.GET("/dashboard", controller.DashboardIndex)
			auth.GET("/user", controller.UserIndex)
			auth.GET("/user/create", controller.CreateUser)
			auth.POST("/user/store", controller.StoreUser)
			auth.GET("/user/edit/:id", controller.EditUser)
			auth.GET("/user/delete/:id", controller.DeteleUser)
			auth.GET("/user/logout", controller.Logout)
		}

		// Routes for guests (not logged in)
		guest := web.Group("/")
		guest.Use(middleware.Guest)
		{
			guest.GET("/login", controller.Login)
			guest.POST("/login", controller.LoginUser)
			guest.GET("/register", controller.Register)
			guest.POST("/register", controller.StoreUser)
		}
	}
}

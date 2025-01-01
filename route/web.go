package route

import (
	"project1/app/http/controller"

	"github.com/gin-gonic/gin"
)

func SetWebRoutes(router *gin.Engine) {
	web := router.Group("/")
	{
		web.GET("/dashboard", controller.DashboardIndex)

		web.GET("/user", controller.UserIndex)

		web.GET("/login", controller.Login)

	}
	// router.Static("/resources", "./resources")
	// router.Static("/public", "./public")
	// ginHtmlRenderer := router.HTMLRender
	// router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}
	// Define a route to serve the templated content
}

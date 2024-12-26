package route

import (
	"net/http"
	"project1/app/http/controller"
	"project1/views/pages/auth"
	"project1/views/pages/user"

	"github.com/gin-gonic/gin"
)

func SetWebRoutes(router *gin.Engine) {
	web := router.Group("/")
	{
		web.Use(middleware.Auth(


		))
		web.GET("/dashboard", controller.UserIndex)

		web.GET("/user", func(c *gin.Context) {
			// Render the "app.tmpl" template from the "views" folder
			c.HTML(http.StatusOK, "", user.Index())
		})

		web.GET("/login", controller.Login)

	}
	// router.Static("/resources", "./resources")
	// router.Static("/public", "./public")
	// ginHtmlRenderer := router.HTMLRender
	// router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}
	// Define a route to serve the templated content
}

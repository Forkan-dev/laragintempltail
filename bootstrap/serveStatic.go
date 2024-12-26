package bootstrap

import (
	"project1/gintemplrenderer"

	"github.com/gin-gonic/gin"
)

func ServeStatic(router *gin.Engine) {
	router.Static("/resources", "./resources")
	router.Static("/public", "./public")

	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}
	// Define a route to serve the templated content
}

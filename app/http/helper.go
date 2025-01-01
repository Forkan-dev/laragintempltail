package http

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func View(c *gin.Context, template templ.Component) {
	templ.Component.Render(template, c, c.Writer)
}

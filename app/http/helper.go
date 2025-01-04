package http

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func View(c *gin.Context, template templ.Component) {
	templ.Component.Render(template, c, c.Writer)
}

func Redirect(c *gin.Context, url string, message string) {
	c.Redirect(http.StatusFound, url)
}

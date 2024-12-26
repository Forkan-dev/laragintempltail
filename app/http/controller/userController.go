package controller

import (
	"net/http"
	"project1/views/pages/dashboard"

	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	// Render the "app.tmpl" template from the "views" folder
	c.HTML(http.StatusOK, "", dashboard.Dashboard())
}

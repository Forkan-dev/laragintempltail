package controller

import (
	"net/http"
	"project1/views/pages/auth"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Render the "app.tmpl" template from the "views" folder
	c.HTML(http.StatusOK, "", auth.Login())
}

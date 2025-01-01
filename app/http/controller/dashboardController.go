package controller

import (
	helper "project1/app/http"
	"project1/views/pages/dashboard"

	"github.com/gin-gonic/gin"
)

func DashboardIndex(c *gin.Context) {
	helper.View(c, dashboard.Dashboard())
}

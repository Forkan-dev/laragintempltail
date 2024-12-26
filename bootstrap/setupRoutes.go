package bootstrap

import (
	"project1/route"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	route.SetWebRoutes(router)

}

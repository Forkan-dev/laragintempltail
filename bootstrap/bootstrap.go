package bootstrap

import "github.com/gin-gonic/gin"

func Serve(port string) {
	router := gin.Default()
	ServeStatic(router)
	SetupRoutes(router)
	router.Run(port)
}

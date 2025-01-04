package bootstrap

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Serve(port string) {

	router := gin.Default()
	// Use cookie-based session store
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	ServeStatic(router)
	SetupRoutes(router)

	router.Run(port)
}

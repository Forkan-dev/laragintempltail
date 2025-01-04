package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authUser := sessions.Default(c).Get("user")
	if authUser == nil {
		c.Redirect(http.StatusFound, "/login")
	}
	c.Next()
}

func Guest(c *gin.Context) {
	authUser := sessions.Default(c).Get("user")
	if authUser != nil {
		c.Redirect(http.StatusFound, "/dashboard")
	}
	c.Next()
}

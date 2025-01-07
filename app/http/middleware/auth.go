package middleware

import (
	"net/http"
	"project1/app/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authUser := sessions.Default(c).Get("user")
	if authUser == nil {
		c.Redirect(http.StatusFound, "/login")
	}

	// slice authUser
	authUserStruct := utils.AuthUser{
		Id:    authUser.([]interface{})[0].(uint),
		Name:  authUser.([]interface{})[1].(string),
		Email: authUser.([]interface{})[2].(string),
	}

	c.Set("user", authUserStruct)
	c.Next()
}

func Guest(c *gin.Context) {
	authUser := sessions.Default(c).Get("user")
	if authUser != nil {
		c.Redirect(http.StatusFound, "/dashboard")
	}
	c.Next()
}

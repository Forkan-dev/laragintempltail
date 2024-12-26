package middleware

import "github.com/gin-gonic/gin"

func Auth(c *gin.Context) {
	// Check if the user is authenticated

	// If not, redirect to the login page
	// If the user is authenticated, call the next handler

}

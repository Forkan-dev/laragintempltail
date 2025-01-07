package utils

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Flash(c *gin.Context, key string, message string) {
	session := sessions.Default(c)
	session.Delete(key)
	session.AddFlash(message, key)
	session.Save()
}

func FlashGet(c *gin.Context, key string) string {
	session := sessions.Default(c)
	message := session.Flashes(key)
	if len(message) == 0 {
		return ""
	}

	session.Save()

	var errorMessage string
	if len(message) > 0 {
		errorMessage = message[0].(string) // Type assertion to string
	} else {
		panic("Flash message not found")
	}

	return errorMessage
}

func SessionSet(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.Delete(key)
	session.Set(key, value)
	session.Save()
	// Save the session
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session", "message": err.Error()})
		return
	}
}

func SessionGet(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	value := session.Get(key)
	return value
}

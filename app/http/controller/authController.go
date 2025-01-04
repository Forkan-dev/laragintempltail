package controller

import (
	"net/http"
	"project1/app/models"
	"project1/database"
	"project1/views/pages/auth"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Render the "app.tmpl" template from the "views" folder
	authUser := sessions.Default(c).Get("user")
	if authUser != nil {
		c.Redirect(http.StatusFound, "/dashboard")
	}
	c.HTML(http.StatusOK, "", auth.Login())
}

func LoginUser(c *gin.Context) {
	var user models.User
	email := strings.TrimSpace(c.PostForm("email"))
	email = strings.ToLower(email) // Normalize to lowercase
	err := database.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"email": c.PostForm("email"),
		})

		return
	}

	// Check if the password is correct
	if !user.CheckPassword(c.PostForm("password")) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid password",
		})
	}

	// save in session
	// Get session store
	session := sessions.Default(c)

	userData := []interface{}{
		user.ID,
		user.Username,
		user.Email,
	}

	session.Set("user", userData) // You can store just the user ID or other information as needed

	// Save the session
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session", "message": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}

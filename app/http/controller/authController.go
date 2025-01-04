package controller

import (
	"fmt"
	"net/http"
	"project1/app/dto"
	"project1/app/models"
	"project1/database"
	"project1/views/pages/auth"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	message := session.Flashes("error")
	session.Save()

	var errorMessage string
	if len(message) > 0 {
		errorMessage = message[0].(string) // Type assertion to string
	}

	data := dto.ResponseData{
		ErrorResponse: dto.ErrorResponse{
			Error: errorMessage,
		},
	}
	fmt.Println("data", message)
	// admin@gmail.com
	c.HTML(http.StatusOK, "", auth.Login(data))
}

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "", auth.Register())
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
		session := sessions.Default(c)
		session.AddFlash("Invalid username or password!", "error")
		session.Save()
		c.Redirect(http.StatusFound, "/login")
		return
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

package controller

import (
	"fmt"
	"net/http"
	"project1/app/dto"
	"project1/app/models"
	"project1/app/utils"
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
		// session.AddFlash("Invalid Email", "error")
		// session.Save()

		utils.Flash(c, "error", "Invalid Email")
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// Check if the password is correct
	if !user.CheckPassword(c.PostForm("password")) {
		utils.Flash(c, "error", "Invalid Password")
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// save in session
	// Get session store

	userData := []interface{}{
		user.ID,
		user.Username,
		user.Email,
	}

	utils.SessionSet(c, "user", userData) // You can store just the user ID or other information as needed

	c.Redirect(http.StatusFound, "/dashboard")
}

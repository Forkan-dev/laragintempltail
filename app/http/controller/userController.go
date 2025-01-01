package controller

import (
	"fmt"
	"net/http"
	"project1/app/dto"
	helper "project1/app/http"
	"project1/app/models"
	"project1/database"
	userView "project1/views/pages/user"

	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	// Initialize a variable to hold user data
	var user []models.User

	if database.DB == nil {
		fmt.Println("DB is nil")
		return
	}
	if err := database.DB.Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user",
		})
		return
	}

	data := makeUserIndexData(user)
	fmt.Println(data.Users[0].Email)

	helper.View(c, userView.Index(data))
}

func makeUserIndexData(user []models.User) dto.Data {

	return dto.Data{
		Title: "User List",
		Users: user,
	}
}

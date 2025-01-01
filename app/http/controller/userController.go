package controller

import (
	"net/http"
	"project1/app/dto"
	helper "project1/app/http"
	"project1/app/models"
	"project1/database"
	userView "project1/views/pages/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	// Initialize a variable to hold user data
	var user []models.User

	if err := database.DB.Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user",
		})
		return
	}

	data := makeUserIndexData(user)

	helper.View(c, userView.Index(data))
}

func CreateUser(c *gin.Context) {
	helper.View(c, userView.CreateUser())
}

func StoreUser(c *gin.Context) {
	// Initialize a variable to hold user data
	var user models.User

	user.Username = c.PostForm("username")
	user.Email = c.PostForm("email")
	user.Address = c.PostForm("address")
	user.Age, _ = strconv.Atoi(c.PostForm("age"))

	// c.JSON(http.StatusOK, gin.H{
	// 	"data": c.PostForm("username"),
	// })
	// return

	// Save the user data
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save user",
		})
		return
	}

	// redirect
	c.Redirect(http.StatusFound, "/user")
}

func makeUserIndexData(user []models.User) dto.Data {

	return dto.Data{
		Title: "User List",
		Users: user,
	}
}

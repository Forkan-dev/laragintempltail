package controller

import (
	"fmt"
	"math/rand"
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

	// Getting random character
	pass := generateRandPassword()
	var user models.User
	c.Set("password", pass)
	fmt.Println(c.Get("password"))
	if err := c.ShouldBind(&user); err != nil {
		// Check for validation errors
		// General binding error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = pass

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save user",
		})
		return
	}

	// redirect back with success message

	c.Redirect(http.StatusFound, "/user")
}

func DeteleUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}

	c.Redirect(http.StatusFound, "/user")
}

func makeUserIndexData(user []models.User) dto.Data {

	return dto.Data{
		Title: "User List",
		Users: user,
	}
}

func generateRandPassword() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var pass string
	for i := 0; i < 8; i++ {
		pass += string(charset[rand.Intn(len(charset))])
	}

	return pass
}

package controller

import (
	"fmt"
	"net/http"
	"project1/app/dto"
	helper "project1/app/http"
	"project1/app/models"
	"project1/app/service"
	"project1/database"
	userView "project1/views/pages/user"

	"github.com/gin-contrib/sessions"
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

	data := makeUserIndexData(user, c)

	helper.View(c, userView.Index(data))
}

func CreateUser(c *gin.Context) {
	helper.View(c, userView.CreateUser())
}

func StoreUser(c *gin.Context) {

	// Getting random character
	pass := service.MakePassword(c.PostForm("password"))
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
			"error":   "Failed to save user",
			"message": err.Error(),
		})
		return
	}

	// redirect back with success message

	c.Redirect(http.StatusFound, "/user")
}

func EditUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user",
		})
		return
	}

	data := dto.EditData{
		Title: "Edit User",
		User:  user,
	}

	helper.View(c, userView.EditUser(data))
}

func DeteleUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.Unscoped().Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}

	c.Redirect(http.StatusFound, "/user")
}

func makeUserIndexData(user []models.User, c *gin.Context) dto.Data {

	authUser := getUserFromSession(c)
	fmt.Println(authUser)
	return dto.Data{
		Title: "User List",
		Users: user,
	}
}

func getUserFromSession(c *gin.Context) interface{} {
	session := sessions.Default(c)
	user := session.Get("user")

	return user
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	// specific key
	session.Delete("user")
	session.Save()

	// c.Redirect(http.StatusFound, "/login")
	// i want redirect back
	c.Redirect(http.StatusFound, c.Request.Referer())
}

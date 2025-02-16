package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"project1/app/dto"
	helper "project1/app/http"
	"project1/app/models"
	"project1/app/service"
	"project1/database"
	userView "project1/views/pages/user"
	"strconv"

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

	// c.Redirect(http.StatusFound, "/user")
	c.Header("HX-Location", "/user")
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

func UpdateUser(c *gin.Context) {
	email := c.PostForm("email")
	fileHeader, err := c.FormFile("avater")
	if err != nil {
		fmt.Println(fileHeader.Filename)
		c.JSON(400, gin.H{"error": "No file uploaded"})
		return
	}

	// Ensure the directory exists
	dir := "public/images"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "Failed to create directory"})
		return
	}

	// Construct the full path
	fullPath := filepath.Join(dir, fileHeader.Filename)
	fmt.Println(fullPath)

	// Save the file
	if err := c.SaveUploadedFile(fileHeader, fullPath); err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "Failed to save file"})
		return
	}

	// Retrieve the user from the database using the email
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "User not found",
			"message": err.Error(),
		})
		return
	}

	pass := service.MakePassword(c.PostForm("password"))

	user.Username = c.PostForm("username")
	user.Email = email
	user.Password = pass
	user.Avater = fullPath
	user.Address = c.PostForm("address")
	user.Age, _ = strconv.Atoi(c.PostForm("age"))

	if err := database.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to save user",
			"message": err.Error(),
		})
		return
	}

	sessions.Default(c).AddFlash("User updated successfully", "success")

	c.Header("HX-Location", "/user")
}

func DeteleUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user",
		})

		return
	}

	authUser := getUserFromSession(c).([]interface{})
	fmt.Println(authUser[2].(string), user.Email)
	if authUser[2].(string) == user.Email {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You can't delete yourself",
		})

		return
	}

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

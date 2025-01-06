package database

import (
	"fmt"
	"project1/app/models"
	"project1/database/seeder"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declare the global DB variable
var DB *gorm.DB

// ConnetDB establishes a connection to the database and migrates the models
func ConnetDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/test_go_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	// Open the database connection and assign it to the global DB variable
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// If the connection fails, panic
		panic("Failed to connect to database!")
	}

	// Auto-migrate the User model
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		fmt.Println("Error migrating database:", err)
		return
	}

	// Optionally, you can log successful database connection
	fmt.Println("Database connected successfully!")
	seeder.SeedUser(10, DB)
}

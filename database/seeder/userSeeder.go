package seeder

import (
	"fmt"
	"project1/app/models"
	"project1/app/service"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func SeedUser(count int, db *gorm.DB) {
	userExist := db.First(&models.User{})

	if userExist != nil {
		fmt.Println("User already exist", userExist)
		return
	}
	for i := 0; i < count; i++ {
		user := models.User{
			Username: gofakeit.Name(),
			Email:    gofakeit.Email(),
			Age:      gofakeit.Number(1, 100),
			Address:  gofakeit.Address().Street,
			Password: service.MakePassword("password"),
			// Add other fields according to your User model
		}

		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("Failed to seed user %d: %v\n", i+1, err)
		}
	}

	fmt.Printf("Successfully seeded %d users!\n", count)
}

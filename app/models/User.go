package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Address  string `json:"address" gorm:"type:text" form:"address"`
	Age      int    `json:"age" gorm:"type:int" form:"age"`
	Avater   string `json:"avater" gorm:"type:varchar(255)" form:"avater"`
	Password string `json:"password" form:"password"`
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

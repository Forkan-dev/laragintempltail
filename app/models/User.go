package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Address  string `json:"address" gorm:"type:text" form:"address"`
	Age      int    `json:"age" gorm:"type:int" form:"age"`
	Password string `json:"password" form:"password"`
}

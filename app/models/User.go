package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address" gorm:"type:text"`
	Age      int    `json:"age"   gorm:"type:int"`
	Password string `json:"password"`
}

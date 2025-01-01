package dto

import "project1/app/models"

type Data struct {
	Title string
	Users []models.User
}

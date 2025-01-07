package utils

import "context"

type AuthUser struct {
	Id    uint
	Name  string
	Email string
}

func Auth(c context.Context) AuthUser {
	return c.Value("user").(AuthUser)
}

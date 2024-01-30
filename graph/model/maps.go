package model

import (
	"github.com/pwnd27/go_app/db"
)

func UserResponse(u db.User) User {
	passwordChangedAt := u.PasswordChangedAt.String()
	createdAt := u.CreatedAt.String()
	return User{
		Username:          u.Username,
		FullName:          u.FullName,
		Email:             u.Email,
		PasswordChangedAt: &passwordChangedAt,
		CreatedAt:         &createdAt,
		Image:             nil,
	}
}

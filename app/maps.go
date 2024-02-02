package app

import (
	"github.com/pwnd27/go_app/db"
	"github.com/pwnd27/go_app/graph/model"
)

func UserResponse(u db.User) model.User {
	passwordChangedAt := u.PasswordChangedAt.String()
	createdAt := u.CreatedAt.String()
	return model.User{
		Username:          u.Username,
		FullName:          u.FullName,
		Email:             u.Email,
		PasswordChangedAt: &passwordChangedAt,
		CreatedAt:         &createdAt,
		Image:             nil,
	}
}

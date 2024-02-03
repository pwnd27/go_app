package app

import (
	"github.com/pwnd27/go_app/db"
	"github.com/pwnd27/go_app/graph/model"
)

func UserResponse(u db.User) model.User {
	return model.User{
		ID:                u.ID,
		Username:          u.Username,
		FullName:          u.FullName,
		Email:             u.Email,
		PasswordChangedAt: &u.PasswordChangedAt,
		CreatedAt:         &u.CreatedAt,
		Image:             nil,
	}
}

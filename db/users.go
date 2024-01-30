package db

import (
	"context"
	"fmt"
	"time"
)

const createUser = `
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
`

type AddUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

type User struct {
	Username          string
	HashedPassword    string
	FullName          string
	Email             string
	PasswordChangedAt time.Time
	CreatedAt         time.Time
}

func (store *SQLStore) AddUser(ctx context.Context, arg AddUserParams) (User, error) {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return User{}, err
	}
	row := tx.QueryRow(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i User
	err = row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return User{}, fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return User{}, err
	}
	err = tx.Commit(ctx)
	return i, err
}

const listUsers = `
SELECT username, hashed_password, full_name, email, password_changed_at, created_at FROM users
`

type ListUsersParams struct {
}

func (store *SQLStore) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	emptyArg := ListUsersParams{}
	if arg != emptyArg {
		return nil, fmt.Errorf("что-то пошло не так")
	}
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err = rows.Scan(
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
			&i.PasswordChangedAt,
			&i.CreatedAt,
		); err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				return nil, fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		}
		items = append(items, i)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	err = tx.Commit(ctx)
	return items, err
}

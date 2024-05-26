package models

import "time"

type User struct {
	Username    string    `db:"username"`
	PassHash    []byte    `db:"password"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

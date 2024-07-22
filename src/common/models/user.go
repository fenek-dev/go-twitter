package models

import "time"

type User struct {
	Id          string    `db:"id" json:"id" redis:"id"`
	Username    string    `db:"username" json:"username" redis:"username"`
	PassHash    []byte    `db:"password" json:"-" redis:"-"`
	Description string    `db:"description" json:"description" redis:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at" redis:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at" redis:"updated_at"`
}

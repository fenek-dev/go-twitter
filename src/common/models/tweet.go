package models

import "time"

type Tweet struct {
	ID        string    `db:"id"`
	Content   string    `db:"content"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

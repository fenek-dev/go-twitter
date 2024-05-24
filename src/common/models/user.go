package models

import "time"

type User struct {
	Username    string
	Email       string
	PassHash    []byte
	Description string
	CreatedAt   time.Duration
	UpdatedAt   time.Duration
}

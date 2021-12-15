package models

import "time"

type User struct {
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

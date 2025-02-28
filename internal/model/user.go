package model

import "time"

type User struct {
	ID        int
	Login     string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

package model

import (
	"time"
)

// User model
type User struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	NickName   string    `json:"nickname"`
	PW         string    `json:"pw"`
	CreateDate time.Time `json:"create_date"`
}

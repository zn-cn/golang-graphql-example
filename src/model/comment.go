package model

import (
	"time"
)

// Comment model
type Comment struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	PostID     int       `json:"post_id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	CreateDate time.Time `json:"create_date"`
}

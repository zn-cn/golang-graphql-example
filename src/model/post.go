package model

import (
	"time"
)

// Post model
type Post struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	PraiseNum  int       `json:"praise_num"`
	CommentNum int       `json:"comment_num"`
	CreateDate time.Time `json:"create_date"`
}

package model

import (
	"time"
)

// Follow model
type Follow struct {
	ID         int       `json:"id"`
	FollowerID int       `json:"follower_id"`
	FolloweeID int       `json:"followee_id"`
	CreateDate time.Time `json:"create_date"`
}

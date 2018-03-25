package handler

import (
	"db"
	"strconv"

	"github.com/graphql-go/graphql"
)

var followArgs = graphql.FieldConfigArgument{
	"token": &graphql.ArgumentConfig{
		Description: "Token to verify",
		Type:        graphql.NewNonNull(graphql.String),
	},
	"follower_id": &graphql.ArgumentConfig{
		Description: "Follower ID",
		Type:        graphql.NewNonNull(graphql.ID),
	},
	"followee_id": &graphql.ArgumentConfig{
		Description: "Followee ID",
		Type:        graphql.NewNonNull(graphql.ID),
	},
}

func follow(p graphql.ResolveParams) (interface{}, error) {
	// JWT token verify
	token := p.Args["token"].(string)
	if ok, checkErr := auth(token); !ok || checkErr != nil {
		return nil, checkErr
	}

	followerID, followerErr := strconv.Atoi(p.Args["follower_id"].(string))
	if followerErr != nil {
		return nil, followerErr
	}
	followeeID, followeeErr := strconv.Atoi(p.Args["followee_id"].(string))
	if followeeErr != nil {
		return nil, followeeErr
	}
	err := db.Follow(followerID, followeeID)
	return (err == nil), err
}

var unfollowArgs = graphql.FieldConfigArgument{
	"token": &graphql.ArgumentConfig{
		Description: "Token to verify",
		Type:        graphql.NewNonNull(graphql.String),
	},
	"follower_id": &graphql.ArgumentConfig{
		Description: "UnFollower ID",
		Type:        graphql.NewNonNull(graphql.ID),
	},
	"followee_id": &graphql.ArgumentConfig{
		Description: "UnFollowee ID",
		Type:        graphql.NewNonNull(graphql.ID),
	},
}

func unfollow(p graphql.ResolveParams) (interface{}, error) {
	// JWT token verify
	token := p.Args["token"].(string)
	if ok, checkErr := auth(token); !ok || checkErr != nil {
		return nil, checkErr
	}

	followerID, followerErr := strconv.Atoi(p.Args["follower_id"].(string))
	if followerErr != nil {
		return nil, followerErr
	}
	followeeID, followeeErr := strconv.Atoi(p.Args["followee_id"].(string))
	if followeeErr != nil {
		return nil, followeeErr
	}
	err := db.Unfollow(followerID, followeeID)
	return (err == nil), err
}

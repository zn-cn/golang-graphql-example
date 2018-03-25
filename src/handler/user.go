package handler

import (
	"db"
	"model"
	"strconv"

	"github.com/graphql-go/graphql"
)

// UserType handler
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*model.User); ok == true {
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"nickname": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*model.User); ok == true {
					return user.NickName, nil
				}
				return nil, nil
			},
		},
		"email": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*model.User); ok == true {
					return user.Email, nil
				}
				return nil, nil
			},
		},
		"create_date": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*model.User); ok == true {
					return user.CreateDate.Format("2006-01-02 15:04:05"), nil
				}
				return nil, nil
			},
		},
		// "pw": &graphql.Field{
		// 	Type: graphql.NewNonNull(graphql.String),
		// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// 		if user, ok := p.Source.(*model.User); ok == true {
		// 			return user.PW, nil
		// 		}
		// 		return nil, nil
		// 	},
		// },
	},
})

func init() {
	UserType.AddFieldConfig("post", &graphql.Field{
		Type: PostType,
		Args: graphql.FieldConfigArgument{
			"post_id": &graphql.ArgumentConfig{
				Description: "Post ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				postID, postErr := strconv.Atoi(p.Args["post_id"].(string))
				if postErr != nil {
					return nil, postErr
				}
				return db.GetPostByIDAndUser(postID, user.ID)
			}
			return model.Post{}, nil
		},
	})
	UserType.AddFieldConfig("posts", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(PostType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				return db.GetPostsForUser(user.ID)
			}
			return []model.Post{}, nil
		},
	})
	UserType.AddFieldConfig("follower", &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"follower_id": &graphql.ArgumentConfig{
				Description: "Follower ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				followerID, followerErr := strconv.Atoi(p.Args["follower_id"].(string))
				if followerErr != nil {
					return nil, followerErr
				}
				return db.GetFollowerByIDAndUser(followerID, user.ID)
			}
			return model.User{}, nil
		},
	})
	UserType.AddFieldConfig("followers", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				return db.GetFollowersForUser(user.ID)
			}
			return []model.User{}, nil
		},
	})
	UserType.AddFieldConfig("followee", &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"followee_id": &graphql.ArgumentConfig{
				Description: "Followee ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				followeeID, followeeErr := strconv.Atoi(p.Args["followee_id"].(string))
				if followeeErr != nil {
					return nil, followeeErr
				}
				return db.GetFolloweeByIDAndUser(followeeID, user.ID)
			}
			return model.User{}, nil
		},
	})
	UserType.AddFieldConfig("followees", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				return db.GetFolloweesForUser(user.ID)
			}
			return []model.User{}, nil
		},
	})
}

var userArgs = graphql.FieldConfigArgument{
	"user_id": &graphql.ArgumentConfig{
		Description: "User ID",
		Type:        graphql.NewNonNull(graphql.ID),
	},
}

func getUserByID(p graphql.ResolveParams) (interface{}, error) {
	userIDStr := p.Args["user_id"].(string)
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, err
	}
	return db.GetUserByID(userID)
}

var createUserArgs = graphql.FieldConfigArgument{
	"nickname": &graphql.ArgumentConfig{
		Description: "User NickName to create",
		Type:        graphql.NewNonNull(graphql.String),
	},
	"email": &graphql.ArgumentConfig{
		Description: "User Email to create",
		Type:        graphql.NewNonNull(graphql.String),
	},
	"pw": &graphql.ArgumentConfig{
		Description: "User PW to create",
		Type:        graphql.NewNonNull(graphql.String),
	},
}

func createUser(p graphql.ResolveParams) (interface{}, error) {
	nickname := p.Args["nickname"].(string)
	email := p.Args["email"].(string)
	pw := p.Args["pw"].(string)

	user := &model.User{
		NickName: nickname,
		Email:    email,
		PW:       pw,
	}
	err := db.InsertUser(user)
	return user, err
}

var removeUserArgs = graphql.FieldConfigArgument{
	"token": &graphql.ArgumentConfig{
		Description: "Token to verify",
		Type:        graphql.NewNonNull(graphql.String),
	},
	"user_id": &graphql.ArgumentConfig{
		Description: "User ID to remove",
		Type:        graphql.NewNonNull(graphql.ID),
	},
}

func removeUser(p graphql.ResolveParams) (interface{}, error) {
	// JWT token verify
	token := p.Args["token"].(string)
	if ok, checkErr := auth(token); !ok || checkErr != nil {
		return nil, checkErr
	}

	userID, err := strconv.Atoi(p.Args["user_id"].(string))
	if err != nil {
		return nil, err
	}
	err = db.RemoveUserByID(userID)
	return (err == nil), err
}

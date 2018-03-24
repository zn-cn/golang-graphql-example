package handler

import (
	"conf"
	"db"
	"model"
	"strconv"
	"util"

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

func login(auth *model.Auth) (bool, error) {
	if ok, err := db.Login(&(auth.User)); !ok {
		return ok, err
	}

	customClaims := util.CustomClaims{
		UserID: auth.User.ID,
		Email:  auth.User.Email,
	}
	var err error
	auth.Token, err = util.CreateJWTToken(conf.Config.JWT.Secret, conf.Config.JWT.SignMethod, customClaims)
	if err != nil {
		return false, err
	}
	return true, nil
}

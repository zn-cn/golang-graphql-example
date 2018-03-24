package handler

import (
	"db"
	"model"
	"strconv"

	"github.com/graphql-go/graphql"
)

// MutationType handler
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "MutationType",
	Fields: graphql.Fields{
		// curl -XPOST http://localhost:1323/graphql -d 'mutation{createUser(nickname:"tofar",email:"yun_tofar@163.com",pw:"1234567879"){id,email,nickname, create_date}}'
		"createUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
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
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
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
			},
		},
		"login": &graphql.Field{
			Type: AuthType,
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Description: "User Email to login",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"pw": &graphql.ArgumentConfig{
					Description: "User PW to login",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["email"].(string)
				pw := p.Args["pw"].(string)
				auth := &model.Auth{
					User: model.User{
						Email: email,
						PW:    pw,
					},
				}
				if ok, err := login(auth); !ok {
					return nil, err
				}
				return auth, nil
			},
		},
		"removeUser": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Description: "User ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userID, err := strconv.Atoi(p.Args["user_id"].(string))
				if err != nil {
					return nil, err
				}
				err = db.RemoveUserByID(userID)
				return (err == nil), err
			},
		},
		"follow": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"follower_id": &graphql.ArgumentConfig{
					Description: "Follower ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"followee_id": &graphql.ArgumentConfig{
					Description: "Followee ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
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
			},
		},
		"unfollow": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"follower_id": &graphql.ArgumentConfig{
					Description: "UnFollower ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"followee_id": &graphql.ArgumentConfig{
					Description: "UnFollowee ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
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
			},
		},
		"createPost": &graphql.Field{
			Type: PostType,
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Description: "User ID of the people who creating a new post",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"title": &graphql.ArgumentConfig{
					Description: "Post Title",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"body": &graphql.ArgumentConfig{
					Description: "Post Body",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userID, userErr := strconv.Atoi(p.Args["user_id"].(string))
				if userErr != nil {
					return nil, userErr
				}
				title := p.Args["title"].(string)
				body := p.Args["body"].(string)
				post := &model.Post{
					UserID: userID,
					Title:  title,
					Body:   body,
				}
				err := db.InsertPost(post)
				return post, err
			},
		},
		"removePost": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"post_id": &graphql.ArgumentConfig{
					Description: "Post ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, postErr := strconv.Atoi(p.Args["post_id"].(string))
				if postErr != nil {
					return nil, postErr
				}
				err := db.RemovePostByID(id)
				return (err == nil), err
			},
		},
		"updatePost": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"post_id": &graphql.ArgumentConfig{
					Description: "Post ID of post to praise",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"title": &graphql.ArgumentConfig{
					Description: "Title of post",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"body": &graphql.ArgumentConfig{
					Description: "Body of post",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				postID, postErr := strconv.Atoi(p.Args["post_id"].(string))
				if postErr != nil {
					return nil, postErr
				}
				title := p.Args["title"].(string)
				body := p.Args["body"].(string)
				err := db.UpdatePost(postID, title, body)
				return (err == nil), err
			},
		},
		"praisePost": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"post_id": &graphql.ArgumentConfig{
					Description: "Post ID to praise",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				postID, postErr := strconv.Atoi(p.Args["post_id"].(string))
				if postErr != nil {
					return nil, postErr
				}
				err := db.PraisePost(postID)
				return (err == nil), err
			},
		},
		"unpraisePost": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"post_id": &graphql.ArgumentConfig{
					Description: "Post ID to unpraise",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				postID, postErr := strconv.Atoi(p.Args["post_id"].(string))
				if postErr != nil {
					return nil, postErr
				}
				err := db.UnPraisePost(postID)
				return (err == nil), err
			},
		},
		"createComment": &graphql.Field{
			Type: CommentType,
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Description: "User id fo the people who creating the new comment",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"post_id": &graphql.ArgumentConfig{
					Description: "Post ID of post commented",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"title": &graphql.ArgumentConfig{
					Description: "Comment Title",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"body": &graphql.ArgumentConfig{
					Description: "Comment Body",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userID, userErr := strconv.Atoi(p.Args["user_id"].(string))
				if userErr != nil {
					return nil, userErr
				}
				postID, postErr := strconv.Atoi(p.Args["post_id"].(string))
				if postErr != nil {
					return nil, postErr
				}
				title := p.Args["title"].(string)
				body := p.Args["body"].(string)
				comment := &model.Comment{
					UserID: userID,
					PostID: postID,
					Title:  title,
					Body:   body,
				}
				err := db.InsertComment(comment)
				return comment, err
			},
		},
		"removeComment": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"comment_id": &graphql.ArgumentConfig{
					Description: "Comment ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				commentID, commentErr := strconv.Atoi(p.Args["comment_id"].(string))
				if commentErr != nil {
					return nil, commentErr
				}
				err := db.RemoveCommentByID(commentID)
				return (err == nil), err
			},
		},
	},
})

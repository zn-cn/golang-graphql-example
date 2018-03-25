package handler

import (
	"db"
	"model"
	"strconv"

	"github.com/graphql-go/graphql"
)

// PostType handler
var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*model.Post); ok == true {
					return post.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*model.Post); ok == true {
					return post.Title, nil
				}
				return nil, nil
			},
		},
		"body": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*model.Post); ok == true {
					return post.Body, nil
				}
				return nil, nil
			},
		},
		"praise_num": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*model.Post); ok == true {
					return post.PraiseNum, nil
				}
				return nil, nil
			},
		},
		"comment_num": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*model.Post); ok == true {
					return post.CommentNum, nil
				}
				return nil, nil
			},
		},
		"create_date": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*model.Post); ok == true {
					return post.CreateDate.Format("2006-01-02 15:04:05"), nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	PostType.AddFieldConfig("user", &graphql.Field{
		Type: graphql.NewNonNull(UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*model.Post); ok == true {
				return db.GetUserByID(post.UserID)
			}
			return nil, nil
		},
	})
	PostType.AddFieldConfig("comment", &graphql.Field{
		Type: CommentType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*model.Post); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return db.GetCommentByIDAndPost(id, post.ID)
			}
			return nil, nil
		},
	})
	PostType.AddFieldConfig("comments", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(CommentType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*model.Post); ok == true {
				return db.GetCommentsForPost(post.ID)
			}
			return []model.Comment{}, nil
		},
	})
}

var createPostArgs = graphql.FieldConfigArgument{
	"token": &graphql.ArgumentConfig{
		Description: "Token to verify",
		Type:        graphql.NewNonNull(graphql.String),
	},
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
}

func createPost(p graphql.ResolveParams) (interface{}, error) {
	// JWT token verify
	token := p.Args["token"].(string)
	if ok, checkErr := auth(token); !ok || checkErr != nil {
		return nil, checkErr
	}

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
}

var removePostArgs = graphql.FieldConfigArgument{
	"token": &graphql.ArgumentConfig{
		Description: "Token to verify",
		Type:        graphql.NewNonNull(graphql.String),
	},
	"post_id": &graphql.ArgumentConfig{
		Description: "Post ID to remove",
		Type:        graphql.NewNonNull(graphql.ID),
	},
}

func removePost(p graphql.ResolveParams) (interface{}, error) {
	// JWT token verify
	token := p.Args["token"].(string)
	if ok, checkErr := auth(token); !ok || checkErr != nil {
		return nil, checkErr
	}

	id, postErr := strconv.Atoi(p.Args["post_id"].(string))
	if postErr != nil {
		return nil, postErr
	}
	err := db.RemovePostByID(id)
	return (err == nil), err
}

var updatePostArgs = graphql.FieldConfigArgument{
	"token": &graphql.ArgumentConfig{
		Description: "Token to verify",
		Type:        graphql.NewNonNull(graphql.String),
	},
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
}

func updatePost(p graphql.ResolveParams) (interface{}, error) {
	// JWT token verify
	token := p.Args["token"].(string)
	if ok, checkErr := auth(token); !ok || checkErr != nil {
		return nil, checkErr
	}

	postID, postErr := strconv.Atoi(p.Args["post_id"].(string))
	if postErr != nil {
		return nil, postErr
	}
	title := p.Args["title"].(string)
	body := p.Args["body"].(string)
	err := db.UpdatePost(postID, title, body)
	return (err == nil), err
}

var praisePostArgs = graphql.FieldConfigArgument{
	"token": &graphql.ArgumentConfig{
		Description: "Token to verify",
		Type:        graphql.NewNonNull(graphql.String),
	},
	"post_id": &graphql.ArgumentConfig{
		Description: "Post ID to praise",
		Type:        graphql.NewNonNull(graphql.ID),
	},
}

func praisePost(p graphql.ResolveParams) (interface{}, error) {
	// JWT token verify
	token := p.Args["token"].(string)
	if ok, checkErr := auth(token); !ok || checkErr != nil {
		return nil, checkErr
	}

	postID, postErr := strconv.Atoi(p.Args["post_id"].(string))
	if postErr != nil {
		return nil, postErr
	}
	err := db.PraisePost(postID)
	return (err == nil), err
}

var unpraisePostArgs = graphql.FieldConfigArgument{
	"token": &graphql.ArgumentConfig{
		Description: "Token to verify",
		Type:        graphql.NewNonNull(graphql.String),
	},
	"post_id": &graphql.ArgumentConfig{
		Description: "Post ID to unpraise",
		Type:        graphql.NewNonNull(graphql.ID),
	},
}

func unpraisePost(p graphql.ResolveParams) (interface{}, error) {
	// JWT token verify
	token := p.Args["token"].(string)
	if ok, checkErr := auth(token); !ok || checkErr != nil {
		return nil, checkErr
	}

	postID, postErr := strconv.Atoi(p.Args["post_id"].(string))
	if postErr != nil {
		return nil, postErr
	}
	err := db.UnPraisePost(postID)
	return (err == nil), err
}

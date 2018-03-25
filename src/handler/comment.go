package handler

import (
	"db"
	"model"
	"strconv"

	"github.com/graphql-go/graphql"
)

// CommentType handler
var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(*model.Comment); ok == true {
					return comment.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(*model.Comment); ok == true {
					return comment.Title, nil
				}
				return nil, nil
			},
		},
		"body": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(*model.Comment); ok == true {
					return comment.Body, nil
				}
				return nil, nil
			},
		},
		"create_date": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(*model.Comment); ok == true {
					return comment.CreateDate.Format("2006-01-02 15:04:05"), nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	CommentType.AddFieldConfig("user", &graphql.Field{
		Type: UserType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if comment, ok := p.Source.(*model.Comment); ok == true {
				return db.GetUserByID(comment.UserID)
			}
			return nil, nil
		},
	})
	CommentType.AddFieldConfig("post", &graphql.Field{
		Type: PostType,
		Args: graphql.FieldConfigArgument{
			"post_id": &graphql.ArgumentConfig{
				Description: "Post ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			postID, postErr := strconv.Atoi(p.Args["post_id"].(string))
			if postErr != nil {
				return nil, postErr
			}
			return db.GetPostByID(postID)
		},
	})
}

var createCommentArgs = graphql.FieldConfigArgument{
	"token": &graphql.ArgumentConfig{
		Description: "Token to verify",
		Type:        graphql.NewNonNull(graphql.String),
	},
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
}

func createComment(p graphql.ResolveParams) (interface{}, error) {
	// JWT token verify
	token := p.Args["token"].(string)
	if ok, checkErr := auth(token); !ok || checkErr != nil {
		return nil, checkErr
	}

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
}

var removeCommentArgs = graphql.FieldConfigArgument{
	"token": &graphql.ArgumentConfig{
		Description: "Token to verify",
		Type:        graphql.NewNonNull(graphql.String),
	},
	"comment_id": &graphql.ArgumentConfig{
		Description: "Comment ID to remove",
		Type:        graphql.NewNonNull(graphql.ID),
	},
}

func removeComment(p graphql.ResolveParams) (interface{}, error) {
	// JWT token verify
	token := p.Args["token"].(string)
	if ok, checkErr := auth(token); !ok || checkErr != nil {
		return nil, checkErr
	}

	commentID, commentErr := strconv.Atoi(p.Args["comment_id"].(string))
	if commentErr != nil {
		return nil, commentErr
	}
	err := db.RemoveCommentByID(commentID)
	return (err == nil), err
}

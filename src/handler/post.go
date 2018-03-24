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

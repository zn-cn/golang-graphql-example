// Package handler implements a simple handler for http
package handler

import (
	"github.com/graphql-go/graphql"
)

// MutationType handler
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "MutationType",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type:    UserType,
			Args:    createUserArgs,
			Resolve: createUser,
		},
		"login": &graphql.Field{
			Type:    AuthType,
			Args:    loginArgs,
			Resolve: login,
		},
		"removeUser": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    removeUserArgs,
			Resolve: removeUser,
		},
		"follow": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    followArgs,
			Resolve: follow,
		},
		"unfollow": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    unfollowArgs,
			Resolve: unfollow,
		},
		"createPost": &graphql.Field{
			Type:    PostType,
			Args:    createPostArgs,
			Resolve: createPost,
		},
		"removePost": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    removePostArgs,
			Resolve: removePost,
		},
		"updatePost": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    updatePostArgs,
			Resolve: updatePost,
		},
		"praisePost": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    praisePostArgs,
			Resolve: praisePost,
		},
		"unpraisePost": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    unpraisePostArgs,
			Resolve: unpraisePost,
		},
		"createComment": &graphql.Field{
			Type:    CommentType,
			Args:    createCommentArgs,
			Resolve: createComment,
		},
		"removeComment": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    removeCommentArgs,
			Resolve: removeComment,
		},
	},
})

package handler

import (
	"github.com/graphql-go/graphql"
)

// QueryType handler
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "QueryType",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type:    UserType,
			Args:    userArgs,
			Resolve: getUserByID,
		},
	},
})

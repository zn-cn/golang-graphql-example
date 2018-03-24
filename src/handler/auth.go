package handler

import (
	"model"

	"github.com/graphql-go/graphql"
)

// AuthType handler
var AuthType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Auth",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if auth, ok := p.Source.(*model.Auth); ok {
					return auth.Token, nil
				}
				return nil, nil
			},
		},
		"user": &graphql.Field{
			Type: UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if auth, ok := p.Source.(*model.Auth); ok {
					return &auth.User, nil
				}
				return nil, nil
			},
		},
	},
})

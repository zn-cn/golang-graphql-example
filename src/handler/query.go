package handler

import (
	"db"
	"strconv"

	"github.com/graphql-go/graphql"
)

// QueryType handler
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "QueryType",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Description: "User ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userIDStr := p.Args["user_id"].(string)
				userID, err := strconv.Atoi(userIDStr)
				if err != nil {
					return nil, err
				}
				return db.GetUserByID(userID)
			},
		},
	},
})

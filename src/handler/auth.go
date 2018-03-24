package handler

import (
	"conf"
	"db"
	"model"
	"util"

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

func auth(token string) (bool, error) {
	claims := util.CustomClaims{}
	ok, err := util.ValidateJWTToken(token, conf.Config.JWT.Secret, &claims)
	if !ok || err != nil {
		return ok, err
	}
	if ok, checkErr := db.CheckUserValid(claims.UserID, claims.Email); !ok || checkErr != nil {
		return false, checkErr
	}
	return true, nil
}

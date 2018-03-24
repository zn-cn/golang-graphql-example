package main

import (
	"conf"
	"db"
	"handler"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	confErr := conf.InitConfig("dev")
	if confErr != nil {
		log.Fatal(confErr)
	}

	dbErr := db.Init()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    handler.QueryType,
		Mutation: handler.MutationType,
	})

	http.Handle("/graphql", handler.Handler(schema))
	log.Fatal(http.ListenAndServe(":1323", nil))
}

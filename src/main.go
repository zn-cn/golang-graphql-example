package main

import (
	"conf"
	"db"
	"handler"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	// gh "github.com/graphql-go/handler"
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

	// h := gh.New(&gh.Config{
	// 	Schema:   &schema,
	// 	Pretty:   true,
	// 	GraphiQL: true,
	// })

	http.Handle("/graphql", handler.Handler(schema))
	// http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":1323", nil))
}

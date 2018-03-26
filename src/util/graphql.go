package util

import (
	"github.com/graphql-go/graphql"
)

// FieldArg util
type FieldArg struct {
	Description string
	Type        graphql.Type
}

// GetGraphqlArgs util （已弃用，感觉写了也没变简洁）
func GetGraphqlArgs(args map[string]FieldArg) graphql.FieldConfigArgument {
	var fieldConfigArg = graphql.FieldConfigArgument{}
	for k, v := range args {
		fieldConfigArg[k] = &graphql.ArgumentConfig{
			Description: v.Description,
			Type:        graphql.NewNonNull(v.Type),
		}
	}
	return fieldConfigArg
}

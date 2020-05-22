/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:23
 * @File  : lion_schema.go
 * @Description: ...
 */
package gql

import "github.com/graphql-go/graphql"

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

var ResultType = graphql.NewObject(graphql.ObjectConfig{
	Name: "resultType",
	Fields: graphql.Fields{
		"state": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "200:成功;401:未授权;其他:错误",
		},
	},
})

var PageInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "pageInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"page": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"page_size": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"filter": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

func GeneratePageType(oType graphql.Type, name string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Fields: graphql.Fields{
			"total": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"items": &graphql.Field{
				Type: graphql.NewList(oType),
			},
		},
	})
}

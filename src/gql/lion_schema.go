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

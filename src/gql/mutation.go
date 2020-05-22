/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:23
 * @File  : mutation.go
 * @Description: ...
 */
package gql

import "github.com/graphql-go/graphql"

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Mutation",
	Fields: graphql.Fields{},
})

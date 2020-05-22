/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:23
 * @File  : query.go
 * @Description: ...
 */
package gql

import (
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Query",
	Fields: graphql.Fields{},
})

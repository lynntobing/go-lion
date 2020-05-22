/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:23
 * @File  : graphql.go
 * @Description: ...
 */
package gql

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func GraphqlHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
		//Playground:       false,
		//RootObjectFn:     nil,
		//ResultCallbackFn: nil,
		//FormatErrorFn:    nil,
	})
	return func(context *gin.Context) {
		h.ServeHTTP(context.Writer, context.Request)
	}
}

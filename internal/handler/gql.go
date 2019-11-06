package handler

import (
	gqlgenhandler "github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/ramonmacias/graphql/internal/gql"
	resolvers "github.com/ramonmacias/graphql/internal/gql/resolvers/generated"
)

func GraphqlHandler() gin.HandlerFunc {
	c := gql.Config{
		Resolvers: &resolvers.Resolver{},
	}
	h := gqlgenhandler.GraphQL(gql.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

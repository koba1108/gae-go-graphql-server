package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/koba1108/gae-go-graphql-server/gql"
	"github.com/koba1108/gae-go-graphql-server/gql/resolvers"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &resolvers.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	_ = r.Run()
}

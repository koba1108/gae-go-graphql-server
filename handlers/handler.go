package handlers

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/koba1108/gae-go-graphql-server/gql"
	"github.com/koba1108/gae-go-graphql-server/gql/resolvers"
	"github.com/vektah/gqlparser/gqlerror"
	"log"
)

func UserGraphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(
		gql.NewExecutableSchema(gql.Config{
			Resolvers: &resolvers.Resolver{},
		}),
		handler.RequestMiddleware(func(ctx context.Context, next func(ctx context.Context) []byte) []byte {
			// リクエスト〜リゾルバを呼ぶ手前？
			log.Printf("RequestMiddleware")
			return next(ctx)
		}),
		handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			// リゾルバが動く手前？
			log.Printf("ResolverMiddleware")
			return next(ctx)
		}),
		/*
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			return fmt.Errorf("internal server error %v", err) // 何かあれば
		}),
		*/
		handler.ErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
			// エラーの振る舞い変えたいモノはここに書けばいけそう？
			log.Printf("ErrorPresenter %v", err)
			return graphql.DefaultErrorPresenter(ctx, err)
		}),
	)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func AdminGraphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(
		gql.NewExecutableSchema(gql.Config{Resolvers: &resolvers.Resolver{}}),
	)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

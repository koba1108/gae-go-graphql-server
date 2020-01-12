package handlers

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/koba1108/gae-go-graphql-server/gql"
	"github.com/koba1108/gae-go-graphql-server/gql/models"
	"github.com/koba1108/gae-go-graphql-server/gql/resolvers"
	"github.com/koba1108/gae-go-graphql-server/middleware"
	"github.com/vektah/gqlparser/gqlerror"
	"log"
)

func UserGraphqlHandler() gin.HandlerFunc {
	gqlConfig := gql.Config{}
	gqlConfig.Resolvers = &resolvers.Resolver{}
	gqlConfig.Directives.HasRole = hasRoleHandler
	h := handler.GraphQL(
		gql.NewExecutableSchema(gqlConfig),
		handler.RequestMiddleware(requestMiddlewareHandler),
		handler.ResolverMiddleware(resolverMiddlewareHandler),
		handler.RecoverFunc(recoverHandler),
		handler.ErrorPresenter(errorPresenter),
	)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func AdminGraphqlHandler() gin.HandlerFunc {
	// とりあえず一緒なやつ
	return UserGraphqlHandler()
}

func PlaygroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func requestMiddlewareHandler(ctx context.Context, next func(ctx context.Context) []byte) []byte {
	// リクエスト〜リゾルバを呼ぶ手前？
	log.Printf("RequestMiddleware")
	return next(ctx)
}

func resolverMiddlewareHandler(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	// リゾルバが動く手前？
	log.Printf("ResolverMiddleware")
	return next(ctx)
}

func recoverHandler(ctx context.Context, err interface{}) error {
	// 何かあれば
	return fmt.Errorf("internal server error %v", err)
}

func errorPresenter(ctx context.Context, err error) *gqlerror.Error {
	// エラーの振る舞い変えたいモノはここに書けばいけそう？
	log.Printf("ErrorPresenter %v", err)
	return graphql.DefaultErrorPresenter(ctx, err)
}

func hasRoleHandler(ctx context.Context, obj interface{}, next graphql.Resolver, role models.Role) (res interface{}, err error) {
	switch role {
	case models.RoleAdmin:
		if middleware.GetAdminIdFromContext(ctx) == "" {
			return nil, fmt.Errorf("access denied")
		}
	case models.RoleUser:
		if middleware.GetUserIdFromContext(ctx) == "" {
			return nil, fmt.Errorf("access denied")
		}
	}
	return next(ctx)
}

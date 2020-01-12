package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koba1108/gae-go-graphql-server/external"
	"github.com/koba1108/gae-go-graphql-server/handlers"
	"github.com/koba1108/gae-go-graphql-server/middleware"
	"os"
)

func init() {
	external.InitExternal()
}

func main() {
	r := gin.Default()
	setContextMiddleware(r)
	setUserGql(r)
	setAdminGql(r)
	if os.Getenv("ENV") != "production" {
		setPlayground(r)
	}
	_ = r.Run()
}

func setContextMiddleware(r *gin.Engine) {
	// @see https://gqlgen.com/recipes/gin/
	r.Use(middleware.GinContextToContext())
}

func setUserGql(r *gin.Engine) {
	r.Use(middleware.Auth())
	r.POST("/graphql", handlers.UserGraphqlHandler())
}

func setAdminGql(r *gin.Engine) {
	r.Use(middleware.AdminAuth())
	r.POST("/admin/graphql", handlers.AdminGraphqlHandler())
}

func setPlayground(r *gin.Engine) {
	r.GET("/", handlers.PlaygroundHandler())
}

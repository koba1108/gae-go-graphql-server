package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koba1108/gae-go-graphql-server/internal/graphql-server/external"
	"github.com/koba1108/gae-go-graphql-server/internal/graphql-server/handlers"
	"github.com/koba1108/gae-go-graphql-server/internal/graphql-server/middleware"
	"os"
)

func init() {
	external.InitExternal()
}

func main() {
	r := getGin()
	setCORSMiddleware(r)
	setContextMiddleware(r)
	setUserGql(r)
	setAdminGql(r)
	if os.Getenv("ENV") != "production" {
		setPlayground(r)
	}
	_ = r.Run()
}

func getGin() *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.HandleMethodNotAllowed = true
	return r
}

func setCORSMiddleware(r *gin.Engine) {
	r.Use(middleware.CORS())
}

// @see https://gqlgen.com/recipes/gin/
func setContextMiddleware(r *gin.Engine) {
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

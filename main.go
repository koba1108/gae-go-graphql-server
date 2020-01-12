package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koba1108/gae-go-graphql-server/external"
	"github.com/koba1108/gae-go-graphql-server/handlers"
	"github.com/koba1108/gae-go-graphql-server/middlewares"
	"os"
)

func init() {
	external.InitExternal()
}

func main() {
	r := gin.Default()
	setUserGql(r)
	setAdminGql(r)
	if os.Getenv("ENV") != "production" {
		setPlayground(r)
	}
	_ = r.Run()
}

func setUserGql(r *gin.Engine) {
	r.POST("/graphql", handlers.UserGraphqlHandler())
}

func setAdminGql(r *gin.Engine) {
	r.Use(middlewares.Auth())
	r.POST("/admin/graphql", handlers.AdminGraphqlHandler())
}

func setPlayground(r *gin.Engine) {
	r.GET("/", handlers.PlaygroundHandler())
}

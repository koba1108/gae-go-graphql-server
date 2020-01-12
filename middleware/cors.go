package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CORS() gin.HandlerFunc {
	var (
		domains []string
		origins []string
	)

	domains = []string{
		"127.0.0.1",
		"localhost",
		"localhost:3000",
		"127.0.0.1:8000",
		"localhost:8000",
		"127.0.0.1:8080",
		"localhost:8080",
		"ci-dot-nacodo-app-dev.appspot.com",
		"nacodo-app-dev.appspot.com",
		"nacodo-app-dev.firebaseapp.com",
		"appdev.naco-do.com",
		"naco-do.com",
		"www.naco-do.com",
		"app.naco-do.com",
		"dev.naco-do.com",
		"admin-dot-nacodo-app-dev.appspot.com",
		"admin-v3-dot-nacodo-app-dev.appspot.com",
		"nacodo-app-230107.appspot.com",
		"admin-dot-nacodo-app-230107.appspot.com",
	}

	for _, p := range []string{"http://", "https://"} {
		for _, d := range domains {
			origins = append(origins, p+d)
		}
	}

	return cors.New(cors.Config{
		AllowOrigins: origins,
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

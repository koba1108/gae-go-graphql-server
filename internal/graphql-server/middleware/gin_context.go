package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GinContextToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContext", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetGinContextFromContext(ctx context.Context) (*gin.Context, error) {
	gc, ok := ctx.Value("GinContext").(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func GetUserIdFromContext(ctx context.Context) string {
	if gc, _ := GetGinContextFromContext(ctx); gc != nil {
		if v, exists := gc.Get("UserId"); exists {
			if userId, ok := v.(string); ok {
				return userId
			}
		}
	}
	return ""
}

func GetAdminIdFromContext(ctx context.Context) string {
	if gc, _ := GetGinContextFromContext(ctx); gc != nil {
		if v, exists := gc.Get("AdminId"); exists {
			if userId, ok := v.(string); ok {
				return userId
			}
		}
	}
	return ""
}

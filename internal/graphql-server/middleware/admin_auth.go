package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken, err := findFirebaseIdToken(c)
		if err != nil {
			renderErrorResponse(c, http.StatusUnauthorized, "authorization header doesn't exist")
			return
		}
		authToken, err := verifyIdToken(c, idToken)
		if err != nil {
			renderErrorResponse(c, http.StatusUnauthorized, "authorization token is not valid")
			return
		}
		admin := AdminUserIdentification{ID: "koba"}
		adminUserIdentification, err := admin.findByUid(authToken.UID)
		if err != nil {
			renderErrorResponse(c, http.StatusUnauthorized, "user not found")
			return
		}
		c.Set("IdToken", idToken)
		c.Set("AdminId", adminUserIdentification.ID)
		c.Set("UID", authToken.UID)
		c.Next()
	}
}

// todo: modelファイル
type AdminUserIdentification struct {
	ID string
}

func (a *AdminUserIdentification) findByUid(uid string) (*AdminUserIdentification, error) {
	return &AdminUserIdentification{
		ID: "ADMIN_ID_KOBA",
	}, nil
}

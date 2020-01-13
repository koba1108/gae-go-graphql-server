package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken, err := findFirebaseIdToken(c)
		if err != nil {
			renderErrorResponse(c, http.StatusUnauthorized,"authorization header doesn't exist")
			return
		}
		authToken, err := verifyIdToken(c, idToken)
		if err != nil {
			renderErrorResponse(c, http.StatusUnauthorized, "authorization token is not valid")
			return
		}
		u := UserIdentification{ID: "koba"}
		userIdentification, err := u.findByUid(authToken.UID)
		if err != nil {
			renderErrorResponse(c, http.StatusUnauthorized, "user not found")
			return
		}
		c.Set("IdToken", idToken)
		c.Set("UserId", userIdentification.ID)
		c.Set("UID", authToken.UID)
		c.Next()
	}
}

// todo: modelファイル
type UserIdentification struct {
	ID string
}

func (u *UserIdentification) findByUid(uid string) (*UserIdentification, error) {
	return &UserIdentification{
		ID: "USER_ID_KOBA",
	}, nil
}

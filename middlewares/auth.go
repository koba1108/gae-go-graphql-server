package middlewares

import (
	"errors"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/koba1108/gae-go-graphql-server/external"
	"github.com/koba1108/gae-go-graphql-server/gql/models"
	"net/http"
	"regexp"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken, err := findFirebaseIdToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "authorization header doesn't exist")
			c.Abort()
			return
		}
		authToken, err := verifyIdToken(c, idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "authorization token is not valid")
			c.Abort()
			return
		}
		u := UserIdentification{ID: "koba"}
		userIdentification, err := u.findByUid(authToken.UID)
		if err != nil {
			//log.Fatalln(err)
			// response.Respond(c, 1, "authorization token is not valid", http.StatusUnauthorized)
			// c.Abort()
			c.JSON(http.StatusUnauthorized, "user not found")
			renderErrorResponse(c, http.StatusUnauthorized, "1", "user not found")
			c.Abort()
			return
		}
		c.Set("IdToken", idToken)
		c.Set("ID", userIdentification.ID)
		c.Set("UID", authToken.UID)
		c.Next()
	}
}

func findFirebaseIdToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 {
		return "", errors.New("authorization header doesn't exist")
	}
	r := regexp.MustCompile(`Bearer\s+(.+)`)
	matches := r.FindStringSubmatch(authHeader)
	if len(matches) < 2 {
		return "", errors.New("authorization token doesn't exist")
	}
	return matches[1], nil
}

func verifyIdToken(c *gin.Context, idToken string) (*auth.Token, error) {
	authClient, err := external.GetFirebaseApp().Auth(c)
	if err != nil {
		return nil, err
	}
	return authClient.VerifyIDToken(c, idToken)
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

func renderErrorResponse(c *gin.Context, httpStatus int,nacodoResponseCode models.NacodoResponseCode, message string) {
	data := models.NacodoResponse{
		Code:    &nacodoResponseCode,
		Message: &message,
	}
	c.IndentedJSON(httpStatus, data)
}

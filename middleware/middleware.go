package middleware

import (
	"errors"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/koba1108/gae-go-graphql-server/external"
	"github.com/koba1108/gae-go-graphql-server/gql/models"
	"regexp"
)

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

func renderErrorResponse(c *gin.Context, httpStatus int, message string) {
	code := models.NacodoResponseCodeNg
	data := models.NacodoResponse{
		Code:    &code,
		Message: &message,
	}
	c.IndentedJSON(httpStatus, data)
	c.Abort()
}

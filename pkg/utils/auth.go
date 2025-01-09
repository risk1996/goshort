package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseAuthBearerToken(c *gin.Context) (string, error) {
	auths, ok := c.Request.Header[http.CanonicalHeaderKey("Authorization")]
	if !ok || len(auths) < 1 {
		return "", errors.New("authorization header not found")
	}

	parts := strings.Split(auths[0], " ")
	if len(parts) < 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid token found")
	}

	return parts[1], nil
}

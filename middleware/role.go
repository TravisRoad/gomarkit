package middleware

import (
	"fmt"
	"net/http"

	"github.com/TravisRoad/gomarkit/errcode"
	"github.com/gin-gonic/gin"
)

// Role is a middleware function that checks if the user has the required roles.
//
// The function takes in a slice of strings, 'roles', which represents the roles that are allowed to access the endpoint.
// It returns a gin.HandlerFunc, which is a function that handles HTTP requests and responses.
func Role(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		flag := false
		r, _ := c.Get("role")
		role := r.(string)
		for _, r := range roles {
			if r == role {
				flag = true
				break
			}
		}
		if !flag {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": errcode.RoleMismatch,
				"msg":  fmt.Sprintf("role mismatch: %v", roles),
			})
			return
		}
		c.Next()
	}
}

package middlewares

import (
	"api/internal/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetInt("role") != 1 {
			responses.Fail(c, http.StatusForbidden, "forbidden", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

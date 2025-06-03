package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		if strings.Contains(token, "Bearer") {
			token = strings.TrimSpace(strings.Replace(token, "Bearer", "", -1))
		}
		// 解析token
		j := NewJwt()
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}

		// 将claims放入上下文
		c.Set("claim", claims)
		c.Next()
	}
}

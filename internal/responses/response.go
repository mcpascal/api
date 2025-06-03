package responses

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(200, Response[interface{}]{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

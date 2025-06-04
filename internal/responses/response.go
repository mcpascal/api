package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IResponse interface {
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(200, Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Fail(c *gin.Context, code int64, message string, err interface{}) {
	if gin.Mode() == gin.ReleaseMode {
		if _, ok := err.(error); ok {
			message = err.(error).Error()
		}
	}

	c.JSON(http.StatusOK, Response{
		code,
		message,
		nil,
	})
}

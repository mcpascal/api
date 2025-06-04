package services

import (
	"api/internal/requests"
	"api/internal/responses"

	"github.com/gin-gonic/gin"
)

type IService interface {
	Index(*gin.Context, *requests.IRequest) (*responses.IResponse, error)
	Show(*gin.Context, *requests.IRequest)
	Store(*gin.Context, *requests.IRequest)
	Update(*gin.Context, *requests.IRequest)
	Destroy(*gin.Context, *requests.IRequest)
}

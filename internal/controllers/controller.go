package controllers

import (
	"api/internal/requests"
	"api/internal/responses"
	"api/internal/services"
	"api/pkg/validator"

	"github.com/gin-gonic/gin"
)

type IController interface {
	Store(*gin.Context, *requests.IRequest, *services.IService)
	Show(*gin.Context, *requests.IRequest, *services.IService)
	Update(*gin.Context, *requests.IRequest, *services.IService)
	Destroy(*gin.Context, *requests.IRequest, *services.IService)
	Index(*gin.Context, *requests.IRequest, *services.IService)
}

type Controller struct {
}

func (c *Controller) Index(ctx *gin.Context, req requests.IRequest, s services.IService) {
	// req := &requests.Login{}
	if err := ctx.ShouldBind(&req); err != nil {
		validator.HandleValidatorError(ctx, err)
		return
	}
	resp, err := s.Index(ctx, &req)
	if err != nil {
		return
	}
	responses.Success(ctx, "login success", resp)
}

func (c *Controller) Store(ctx *gin.Context) {

}

func (c *Controller) Show(ctx *gin.Context) {

}

func (c *Controller) Update(ctx *gin.Context) {

}

func (c *Controller) Destroy(ctx *gin.Context) {

}

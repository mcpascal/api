package controllers

import (
	"api/internal/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IController interface {
	// Store(*gin.Context, *requests.IRequest, *services.IService)
	// Show(*gin.Context, *requests.IRequest, *services.IService)
	// Update(*gin.Context, *requests.IRequest, *services.IService)
	// Destroy(*gin.Context, *requests.IRequest, *services.IService)
	// Index(*gin.Context, *requests.IRequest, *services.IService)
}

type Controller struct {
}

func (c *Controller) CheckParams(ctx *gin.Context, req interface{}) {
	if err := ctx.ShouldBind(&req); err != nil {
		responses.Fail(ctx, 400, "params error", err)
		return
	}
}

func (c *Controller) GetParamId(ctx *gin.Context) int {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		responses.Fail(ctx, 400, "id error", err)
		return 0
	}
	return id
}

// func NewController(opts ...ControllerOptionFunc) *Controller {
// 	c := &Controller{}
// 	for _, opt := range opts {
// 		opt(c)
// 	}
// 	return c
// }

// type ControllerOptionFunc func(c *Controller)

// type Controller struct {
// 	Service *services.IService
// 	// IndexRequest   requests.IRequest
// 	// ShowRequest    requests.IRequest
// 	// StoreRequest   requests.IRequest
// 	// UpdateRequest  requests.IRequest
// 	// DestroyRequest requests.IRequest
// }

// func WithService(s services.IService) ControllerOptionFunc {
// 	return func(c *Controller) {
// 		c.Service = &s
// 	}
// }

// func (c *Controller) Index(ctx *gin.Context) {
// 	// req := c.IndexRequest
// 	// s := *c.Service
// 	// if err := ctx.ShouldBind(&req); err != nil {
// 	// 	validator.HandleValidatorError(ctx, err)
// 	// 	return
// 	// }
// 	// resp, err := s.Index(ctx, &req)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// responses.Success(ctx, "login success", resp)
// }

// func (c *Controller) Store(ctx *gin.Context) {

// }

// func (c *Controller) Show(ctx *gin.Context) {

// }

// func (c *Controller) Update(ctx *gin.Context) {

// }

// func (c *Controller) Destroy(ctx *gin.Context) {

// }

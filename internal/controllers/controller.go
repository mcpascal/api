package controllers

import (
	"api/internal/responses"
	validate "api/pkg/validator"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IController interface {
	Store(*gin.Context)
	Show(*gin.Context)
	Update(*gin.Context)
	Destroy(*gin.Context)
	Index(*gin.Context)
}

type Controller struct{}

// HandleValidatorError 处理字段校验异常
func (c *Controller) HandleValidatorError(ctx *gin.Context, err error) {
	//如果返回错误信息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		responses.Fail(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}
	errMap := c.removeTopStruct(errs.Translate(validate.Trans))
	var errstr string
	for k, v := range errMap {
		if len(errstr) > 0 {
			errstr += ";"
		}
		errstr += k + ":" + v
	}
	responses.Fail(ctx, http.StatusBadRequest, errstr, errors.New(errstr))
	return
}

// removeTopStruct 定义一个去掉结构体名称前缀的自定义方法：
func (c *Controller) removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
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

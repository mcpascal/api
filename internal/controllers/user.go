package controllers

import (
	"api/internal/requests"
	"api/internal/responses"
	"api/internal/services"

	"github.com/gin-gonic/gin"
)

type User struct {
	Controller
	service *services.User
}

func NewUser() *User {
	return &User{
		service: services.NewUser(),
	}
}

/**
 * Index
 * @Description: 用户列表
 * @receiver u
 * @param c
 */
func (u *User) Index(c *gin.Context) {
	req := requests.Paginator{}
	if err := c.ShouldBind(&req); err != nil {
		u.HandleValidatorError(c, err)
		return
	}
	users, err := u.service.Index(&req)
	if err != nil {
		responses.Fail(c, 500, "index failed", err)
		return
	}
	responses.Success(c, "index success", users)
}

func (u *User) Show(c *gin.Context) {
	r := requests.Id{}
	if err := c.ShouldBind(&r); err != nil {
		u.HandleValidatorError(c, err)
	}
	data, err := u.service.Show(r.Id)
	if err != nil {
		responses.Fail(c, 500, "show failed", err)
		return
	}
	responses.Success(c, "show success", data)
}

func (u *User) Store(c *gin.Context) {
	var r requests.User
	if err := c.ShouldBind(&r); err != nil {
		u.HandleValidatorError(c, err)
		return
	}
	user, err := u.service.Store(&r)
	if err != nil {
		responses.Fail(c, 500, "store failed", err)
		return
	}
	responses.Success(c, "store success", user)
}

func (u *User) Update(c *gin.Context) {
	id := requests.Id{}
	if err := c.ShouldBind(&id); err != nil {
		u.HandleValidatorError(c, err)
	}
	r := requests.User{}
	if err := c.ShouldBind(&r); err != nil {
		u.HandleValidatorError(c, err)
		return
	}
	user, err := u.service.Update(id.Id, &r)
	if err != nil {
		responses.Fail(c, 500, "update failed", err)
		return
	}
	responses.Success(c, "update success", user)
}

func (u *User) Destroy(c *gin.Context) {
	id := requests.Id{}
	if err := c.ShouldBind(&id); err != nil {
		u.HandleValidatorError(c, err)
	}
	if err := u.service.Destroy(id.Id); err != nil {
		responses.Fail(c, 500, "delete failed", err)
		return
	}
	responses.Success(c, "delete success", nil)
}

func (u *User) Search(c *gin.Context) {
	req := requests.Search{}
	if err := c.ShouldBind(&req); err != nil {
		u.HandleValidatorError(c, err)
		return
	}
	users, err := u.service.Search(&req)
	if err != nil {
		responses.Fail(c, 500, "index failed", err)
		return
	}
	responses.Success(c, "index success", users)
}

func (u *User) List(c *gin.Context) {
	req := requests.Search{}
	if err := c.ShouldBind(&req); err != nil {
		u.HandleValidatorError(c, err)
		return
	}
	users, err := u.service.Search(&req)
	if err != nil {
		responses.Fail(c, 500, "list failed", err)
		return
	}
	responses.Success(c, "list success", users)
}

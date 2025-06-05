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
	u.CheckParams(c, &req)
	users, err := u.service.Index(&req)
	if err != nil {
		responses.Fail(c, 500, "index failed", err)
		return
	}
	responses.Success(c, "index success", users)
}

func (u *User) Show(c *gin.Context) {
	id := u.GetParamId(c)
	data, err := u.service.Show(id)
	if err != nil {
		responses.Fail(c, 500, "show failed", err)
		return
	}
	responses.Success(c, "show success", data)
}

func (u *User) Store(c *gin.Context) {
	var req requests.User
	u.CheckParams(c, &req)
	user, err := u.service.Store(&req)
	if err != nil {
		responses.Fail(c, 500, "store failed", err)
		return
	}
	responses.Success(c, "store success", user)
}

func (u *User) Update(c *gin.Context) {
	req := requests.User{}
	id := u.GetParamId(c)
	u.CheckParams(c, &req)
	user, err := u.service.Update(id, &req)
	if err != nil {
		responses.Fail(c, 500, "update failed", err)
		return
	}
	responses.Success(c, "update success", user)
}

func (u *User) Destroy(c *gin.Context) {
	id := u.GetParamId(c)
	if err := u.service.Destroy(id); err != nil {
		responses.Fail(c, 500, "delete failed", err)
		return
	}
	responses.Success(c, "delete success", nil)
}

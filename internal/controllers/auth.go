package controllers

import (
	"api/internal/requests"
	"api/internal/responses"
	"api/internal/services"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	service *services.Auth
}

func NewAuth() *Auth {
	return &Auth{
		service: services.NewAuth(),
	}
}

func (a *Auth) Login(c *gin.Context) {
	var req *requests.Login
	if err := c.ShouldBind(req); err != nil {
		return
	}
	token, refreshToken, err := a.service.Login(c, req)
	if err != nil {
		return
	}
	resp := &responses.Login{
		Token:        token,
		RefreshToken: refreshToken,
	}
	responses.Success(c, "login success", resp)
}

// func (a *Auth) Register(c *gin.Context) {
// 	var req *requests.Register
// 	if err := c.ShouldBind(req); err != nil {
// 		return
// 	}
// 	err := a.service.Register(c, req)
// 	if err != nil {
// 		return
// 	}
// 	responses.Success(c, "register success", nil)
// }

// func (a *Auth) Logout(c *gin.Context) {
// 	a.service.Logout(c)
// 	responses.Success(c, "logout success", nil)
// }

// func (a *Auth) RefreshToken(c *gin.Context) {
// 	var req *requests.RefreshToken
// 	if err := c.ShouldBind(req); err != nil {
// 		return
// 	}
// 	token, err := a.service.RefreshToken(c)
// 	if err != nil {
// 		return
// 	}
// 	resp := &responses.Login{
// 		Token: token,
// 	}
// 	responses.Success(c, "refresh token success", resp)
// }

// func (a *Auth) ResetPassword(c *gin.Context) {
// 	var req *requests.ResetPassword
// 	if err := c.ShouldBind(req); err != nil {
// 		return
// 	}
// 	err := a.service.ResetPassword(c, req)
// 	if err != nil {
// 		return
// 	}
// 	responses.Success(c, "reset password success", nil)
// }

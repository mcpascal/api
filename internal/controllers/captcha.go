package controllers

import "github.com/gin-gonic/gin"

type Captcha struct {
}

func NewCaptcha() *Captcha {
	return &Captcha{}
}

func (c *Captcha) Generate(ctx *gin.Context) {

}

func (c *Captcha) Verify(ctx *gin.Context) {

}

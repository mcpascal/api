package controllers

import "github.com/gin-gonic/gin"

type Sms struct {
}

func NewSms() *Sms {
	return &Sms{}
}

func (s *Sms) Generate(c *gin.Context) {

}
func (s *Sms) Verify(c *gin.Context) {

}

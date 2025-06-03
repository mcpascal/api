package services

import (
	"api/internal/repositories"
	"api/internal/requests"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	defaultDuration = time.Minute * 5
	defaultTimes    = 5
)

type Auth struct {
	r repositories.User
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Login(c *gin.Context, req *requests.Login) (string, string, error) {
	var err error
	// // check captcha
	// err = a.VarifyCaptcha(c, req.Captcha)
	// if err != nil {
	// 	return "", err
	// }
	// // check login attempt fail times
	// ip := c.ClientIP()
	// err = a.CheckLoginAttempts(ip, string(user.ID))
	// if err != nil {
	// 	return "", err
	// }
	// check user exist
	user, err := a.r.FindByEmail(req.Email)
	fmt.Println(user)
	if err != nil {
		return "", "", err
	}
	return "", "", nil
	// check password
	// err = a.VarifyPassword(req.Password, user.Password)
	// if err != nil {
	// 	return "", "", err
	// }

	// // generate token
	// jwt := utils.NewJwt()
	// return jwt.GenerateToken(utils.CustomClaims{
	// 	Id:       uint(user.ID),
	// 	Username: user.Username,
	// 	Email:    user.Email,
	// })
}

func (a *Auth) CheckLoginAttempts(ip string, user string) error {
	// count, ok := loginCache.Get(ip)
	//
	//	if ok && count >= defaultTimes {
	//		common.ErrorStrResp(c, "Too many unsuccessful sign-in attempts have been made using an incorrect username or password, Try again later.", 429)
	//		loginCache.Expire(ip, defaultDuration)
	//		return
	return nil
}

// func (a *Auth) Register(c *gin.Context, req *requests.Register) error {
// 	user := &models.User{
// 		Email: req.Email,
// 	}
// 	user, err := a.r.FindOneByEmail(req.Email)
// 	if user.ID != 0 {
// 		return errors.New("email already exists")
// 	}
// 	// hash password
// 	hash, err := a.r.HashPassword(req.Password)
// 	if err != nil {
// 		return err
// 	}
// 	// create user
// 	user.Password = hash
// 	_, err = a.r.Create(&user)
// 	return err
// }

// ip := c.ClientIP()
// 	count, ok := loginCache.Get(ip)
// 	if ok && count >= defaultTimes {
// 		common.ErrorStrResp(c, "Too many unsuccessful sign-in attempts have been made using an incorrect username or password, Try again later.", 429)
// 		loginCache.Expire(ip, defaultDuration)
// 		return

// func (a *Auth) VarifyCaptcha(c *gin.Context, captcha string) error {
// 	// 从缓存中获取验证码
// 	key := fmt.Sprintf("captcha:%s", c.ClientIP())
// 	savedCaptcha, err := cache.Get(key)
// 	if err != nil {
// 		return errors.New("验证码已过期")
// 	}

// 	if !strings.EqualFold(savedCaptcha, captcha) {
// 		return errors.New("验证码错误")
// 	}

// 	// 验证成功后删除验证码
// 	cache.Del(key)
// 	return nil
// }

// func (a *Auth) VarifyPassword(password, hashedPassword string) error {
// 	return nil
// }

// func (a *Auth) GenerateToken(user models.User) (string, error) {
// 	return "", nil
// }

// func (a *Auth) recordLoginAttempt(userID uint, ip string, success bool, reason string) error {
// 	log := &models.LoginLog{
// 		UserID:    userID,
// 		IP:        ip,
// 		UserAgent: c.Request.UserAgent(),
// 		Success:   success,
// 		Reason:    reason,
// 		CreatedAt: time.Now(),
// 	}
// 	return a.loginLogRepo.Create(log)
// }

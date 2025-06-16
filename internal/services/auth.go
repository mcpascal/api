package services

import (
	"api/internal/repositories"
	"api/internal/requests"
	"api/internal/responses"
	"api/pkg/cache"
	"api/pkg/jwt"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	defaultDuration = time.Minute * 5
	defaultTimes    = 5
)

type Auth struct {
	repository *repositories.Auth
}

func NewAuth() *Auth {
	return &Auth{
		repository: repositories.NewAuth(),
	}
}

func (a *Auth) Login(c *gin.Context, req *requests.Login) (*responses.Login, error) {
	var err error
	resp := &responses.Login{}
	// check captcha
	err = a.VarifyCaptcha(c, req.Captcha)
	if err != nil {
		return nil, err
	}
	// check login attempt fail times
	ip := c.ClientIP()
	err = a.CheckLoginAttempts(ip)
	if err != nil {
		return nil, err
	}
	// check user exist
	userRepo := repositories.NewUser()
	user, err := userRepo.FindByEmail(req.Email)
	fmt.Println(user)
	if err != nil {
		return resp, err
	}
	// check password
	err = a.VarifyPassword(req.Password, user.Password)
	if err != nil {
		return nil, err
	}

	// generate token
	j := jwt.NewJwt()
	accessToken, refreshToken, err := j.Create(jwt.CustomClaims{
		Id:       uint(user.ID),
		Username: user.Username,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(accessToken, refreshToken)
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken
	return resp, nil
}

func (a *Auth) CheckLoginAttempts(ip string) error {
	// count, ok := loginCache.Get(ip)
	//
	//	if ok && count >= defaultTimes {
	//		common.ErrorStrResp(c, "Too many unsuccessful sign-in attempts have been made using an incorrect username or password, Try again later.", 429)
	//		loginCache.Expire(ip, defaultDuration)
	//		return
	return nil
}

func (a *Auth) Register(c *gin.Context, req *requests.Register) (responses.Register, error) {
	resp := responses.Register{}
	// user := &models.User{
	// 	Email: req.Email,
	// }
	// user, err := a.r.FindOneByEmail(req.Email)
	// if user.ID != 0 {
	// 	return errors.New("email already exists")
	// }
	// // hash password
	// hash, err := a.r.HashPassword(req.Password)
	// if err != nil {
	// 	return err
	// }
	// // create user
	// user.Password = hash
	// _, err = a.r.Create(&user)
	// return err
	return resp, nil
}

// ip := c.ClientIP()
// 	count, ok := loginCache.Get(ip)
// 	if ok && count >= defaultTimes {
// 		common.ErrorStrResp(c, "Too many unsuccessful sign-in attempts have been made using an incorrect username or password, Try again later.", 429)
// 		loginCache.Expire(ip, defaultDuration)
// 		return

func (a *Auth) VarifyCaptcha(c *gin.Context, captcha string) error {
	// 从缓存中获取验证码
	key := fmt.Sprintf("captcha:%s", c.ClientIP())
	if value, ok := cache.Cache.Get(key); ok {
		fmt.Println(value)
	}
	// savedCaptcha, err := cache.Get(key)
	// if err != nil {
	// 	return errors.New("验证码已过期")
	// }

	// if !strings.EqualFold(savedCaptcha, captcha) {
	// 	return errors.New("验证码错误")
	// }

	// // 验证成功后删除验证码
	// // cache.Del(key)
	return nil
}

func (a *Auth) VarifyPassword(password, hashedPassword string) error {
	return nil
}

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

package utils

import (
	"api/configs"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义Claims
type CustomClaims struct {
	Id uint `json:"id"`
	jwt.RegisteredClaims
}

type Jwt struct {
	Secret []byte
	Expire int64
}

// 定义错误
var (
	TokenExpired     = "Token已过期"
	TokenNotValidYet = "Token尚未生效"
	TokenMalformed   = "Token格式错误"
	TokenInvalid     = "无效的Token"
)

// NewJWT 新建JWT实例
func NewJwt() *Jwt {
	config := configs.App.JwtInfo
	return &Jwt{
		Secret: []byte(config.Secret), // 设置签名密钥
		Expire: int64(config.Expire),  // 设置过期时间
	}
}

// CreateToken 生成Token
func (j *Jwt) GenerateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.Secret)
}

// ParseToken 解析Token
func (j *Jwt) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	} else {
		return nil, jwt.ErrTokenInvalidClaims
	}
}

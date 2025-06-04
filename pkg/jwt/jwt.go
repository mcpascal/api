package jwt

import (
	"api/configs"
	"errors"

	v5 "github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义Claims
type CustomClaims struct {
	Id uint `json:"id"`
	v5.RegisteredClaims
}

type Jwt struct {
	Secret []byte
	Expire int64
}

// 定义错误
var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
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
func (j *Jwt) Create(claims CustomClaims) (string, error) {
	token := v5.NewWithClaims(v5.SigningMethodHS256, claims)
	return token.SignedString(j.Secret)
}

// ParseToken 解析Token
func (j *Jwt) Parse(tokenString string) (*CustomClaims, error) {
	token, err := v5.ParseWithClaims(tokenString, &CustomClaims{}, func(token *v5.Token) (interface{}, error) {
		return j.Secret, nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	} else {
		return nil, v5.ErrTokenInvalidClaims
	}
}

// // 更新token
// func (j *Jwt) Refresh(tokenString string) (string, error) {
// 	v5.TimeFunc = func() time.Time {
// 		return time.Unix(0, 0)
// 	}
// 	token, err := v5.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return j.Secret, nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
// 		v5.TimeFunc = time.Now
// 		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
// 		return j.Create(*claims)
// 	}
// 	return "", TokenInvalid
// }

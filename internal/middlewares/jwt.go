package middlewares

import (
	"api/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义Claims
// type CustomClaims struct {
// 	ID       uint   `json:"id"`
// 	Username string `json:"username"`
// 	jwt.RegisteredClaims
// }

type JWT struct {
	SigningKey []byte
}

// 定义错误
var (
	TokenExpired     = "Token已过期"
	TokenNotValidYet = "Token尚未生效"
	TokenMalformed   = "Token格式错误"
	TokenInvalid     = "无效的Token"
)

// NewJWT 新建JWT实例
func NewJwt() *JWT {
	return &JWT{
		SigningKey: []byte(configs.App.JwtInfo.Secret), // 设置签名密钥
	}
}

// CreateToken 生成Token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	} else {
		return nil, jwt.ErrTokenInvalidClaims
	}
}

// 在 CustomClaims 中添加刷新令牌的字段
type CustomClaims struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	TokenType string `json:"token_type"` // access 或 refresh
	jwt.RegisteredClaims
}

// 创建访问令牌和刷新令牌
func (j *JWT) CreateTokenPair(claims CustomClaims) (accessToken, refreshToken string, err error) {
	// 访问令牌 - 短期有效
	claims.TokenType = "access"
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 2))
	accessToken, err = j.CreateToken(claims)
	if err != nil {
		return "", "", err
	}

	// 刷新令牌 - 长期有效
	claims.TokenType = "refresh"
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7))
	refreshToken, err = j.CreateToken(claims)
	return
}

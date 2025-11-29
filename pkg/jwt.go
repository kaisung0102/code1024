package pkg

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const (
	TokenExpireDuration = time.Hour * 24 // 过期时间
	// jachow
	TokenIssuer = "jachow"
)

// 用于签名的字符串
var SigningKey = []byte("awslhgfrtyxp")

// GenToken 使用默认声明创建jwt
func GenToken(userID int64, username string) (string, error) {
	// 创建 Claims
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间
			Issuer:    TokenIssuer,                                             // 签发人
		},
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString(SigningKey)
}

// ParseToken 解析jwt
func ParseToken(tokenString string) (*Claims, error) {
	claims := new(Claims)
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})

	if err != nil { // 解析token失败
		return nil, err
	}

	if !token.Valid { // token无效
		return nil, errors.New("token is invalid")
	}

	return claims, nil
}

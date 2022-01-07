package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// UserClaims 用户信息类，作为生成token的参数
type UserClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	//自定义的token秘钥
	secret = []byte("eyJpZCI6IjAwMSIsIm5hbWUiOiJhZG1pbiIsImV4cCI6MTY0MTU0MzCwN30")
	//token有效时间
	effectTime = 24 * time.Hour
)

var (
	TokenInvalid = errors.New("Token is invalid")
)

// GenerateToken 生成token
func GenerateToken(claims *UserClaims) (string, error) {
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		return "", err
	}
	return sign, nil
}

// 解析Token
func parseToken(tokenString string) (*UserClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, TokenInvalid
	}
	claims := token.Claims.(*UserClaims)
	return claims, nil
}

// Refresh 更新token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims := token.Claims.(*UserClaims)
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	generateToken, _ := GenerateToken(claims)
	return generateToken
}

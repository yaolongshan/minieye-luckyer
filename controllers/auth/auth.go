package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 指定账号密码
var (
	UNM = "admin"
	PWD = "Minieye2022"
)

// Login 登录
func Login(c *gin.Context) {
	type req struct {
		UNM string `json:"unm"`
		PWD string `json:"pwd"`
	}
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "json参数错误",
			"Error":  err.Error()})
		return
	}
	if r.UNM != UNM && r.PWD != PWD {
		c.JSON(http.StatusForbidden, gin.H{
			"Status": false,
			"Msg":    "账号密码错误"})
		return
	}
	token, _ := GenerateToken(&UserClaims{
		ID:             "001",
		Name:           UNM,
		StandardClaims: jwt.StandardClaims{},
	})
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// JWTAuth 用户校验
func JWTAuth(c *gin.Context) {
	// 跳过/api/login路由
	if c.Request.RequestURI == "/api/login" {
		return
	}
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"Status": false,
			"Msg":    "无权访问，请求未携带token"})
		c.Abort()
		return
	}
	_, err := parseToken(token)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"Status": false,
			"Msg":    "无权访问，请检查token",
			"Error":  err.Error()})
		c.Abort()
		return
	}
}

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
	if r.UNM != UNM || r.PWD != PWD {
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
	_, err := c.Cookie("m5hbWUiOiJhZG1pbiIs")
	if err != nil {
		//c.SetCookie("m5hbWUiOiJhZG1pbiIs", token, 0, "/", "localhost", false, true)
		c.SetCookie("m5hbWUiOiJhZG1pbiIs", token, 0, "/", "annual-2022.minieye.tech", false, true)
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Msg":    "登录成功"})
}

// IsLogin 判断用户是否登录
func IsLogin(c *gin.Context) {
	token, err := c.Cookie("m5hbWUiOiJhZG1pbiIs")
	// cookie中是否携带token
	if err != nil || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status": false,
			"Msg":    "请登录后重试",
			"Error":  err.Error()})
		return
	}
	// 携带的token进行校验
	_, err = parseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status": false,
			"Msg":    "请登录后重试",
			"Error":  err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Msg":    "您已登录"})
}

// JWTAuth 用户校验
func JWTAuth(c *gin.Context) {
	// 跳过/api/login路由
	if c.Request.RequestURI == "/api/login" || c.Request.RequestURI == "/api/islogin"{
		return
	}
	token, err := c.Cookie("m5hbWUiOiJhZG1pbiIs")
	// cookie中是否携带token
	if err != nil || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status": false,
			"Msg":    "无权访问，请登录后重试",
			"Error":  err.Error()})
		c.Abort()
		return
	}
	// 携带的token进行校验
	_, err = parseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status": false,
			"Msg":    "无权访问，请登录后重试",
			"Error":  err.Error()})
		c.Abort()
		return
	}
	c.Next()
}

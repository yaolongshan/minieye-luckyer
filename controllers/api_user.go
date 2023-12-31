package controllers

import (
	"code/minieye-luckyer/models/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Number string `json:"number"`
	Type   string `json:"type"`
}

// ApiGetAllUser 获取所有员工
func ApiGetAllUser(c *gin.Context) {
	users := db.GetUserList()
	count := len(users)
	c.JSON(http.StatusOK, gin.H{"Status": true, "Count": count, "Users": users})
}

// ApiAddUser 添加员工
func ApiAddUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "json参数错误",
			"Error":  err.Error()})
		return
	}
	db.AddUser(user.Name, user.Phone, user.Number, user.Type)
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Msg":    "ok"})
}

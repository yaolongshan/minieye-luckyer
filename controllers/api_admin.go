package controllers

import (
	"code/minieye-luckyer/models/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Prize struct {
	Name string `json:"name"`
	Sum  int    `json:"sum"`
}

// ApiGetAllUser 获取所有员工
func ApiGetAllUser(c *gin.Context) {
	users := db.GetUserList()
	count := db.UserCount()
	c.JSON(http.StatusOK, gin.H{"Count": count, "Users": users})
}

// ApiAddUser 添加员工
func ApiAddUser(c *gin.Context) {
	type User struct {
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Type     string `json:"type"`
		Number   string `json:"number"`
		Contract string `json:"contract"`
		Mail     string `json:"mail"`
	}
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"msg":    err.Error()})
		return
	}
	db.AddUser(user.Name, user.Phone, user.Type, user.Number, user.Contract, user.Mail)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"msg":    "ok"})
}

// ApiAddPrize 添加一个奖项
func ApiAddPrize(c *gin.Context) {
	var p Prize
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"msg":    err.Error()})
		return
	}
	err := db.AddPrize(p.Name, p.Sum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"msg":    err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"msg":    "ok"})
}

// ApiGetAllPrize 奖项列表
func ApiGetAllPrize(c *gin.Context) {
	prizes := db.GetPrizeList()
	c.JSON(http.StatusOK, gin.H{"prizes": prizes})
}

// ApiUpdatePrize 修改奖项的数量
func ApiUpdatePrize(c *gin.Context) {
	var p Prize
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"msg":    err.Error()})
		return
	}
	err := db.UpdatePrize(p.Name, p.Sum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"msg":    err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"msg":    "ok"})
}

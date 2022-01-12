package controllers

import (
	"code/minieye-luckyer/models/db"
	"code/minieye-luckyer/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type req struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Content string `json:"content"`
}

func ApiSendSMS(c *gin.Context) {
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "json参数错误",
			"Error":  err.Error()})
		return
	}
	err := services.SendSMS(r.Name, r.Content, r.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "发送失败",
			"Error":  err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": true, "Msg": "发送成功"})
}

func ApiSendDingDing(c *gin.Context) {
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "json参数错误",
			"Error":  err.Error()})
		return
	}
	err := services.SendDingDingMsg(r.Phone, r.Name, r.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "发送失败",
			"Error":  err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": true, "Msg": "发送成功"})
}

// ApiNotifyGet 获取当前通知功能的开关状态
func ApiNotifyGet(c *gin.Context) {
	value := db.GetValue("notify").Value
	status, _ := strconv.ParseBool(value)
	c.JSON(http.StatusOK, gin.H{
		"NotifyStatus": status,
	})
}

// ApiNotifySet 设置发送短信、钉钉功能的开启或关闭
func ApiNotifySet(c *gin.Context) {
	status, err := strconv.ParseBool(c.Query("status"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "status参数错误",
			"Error":  err.Error(),
		})
		return
	}
	db.SetValue("notify", strconv.FormatBool(status))
	curr := db.GetValue("notify").Value
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Msg":    fmt.Sprintf("设置成功，当前状态为%v", curr),
	})
}
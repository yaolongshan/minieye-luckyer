package controllers

import (
	"code/minieye-luckyer/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiSendSMS(c *gin.Context) {
	type req struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Content string `json:"content"`
	}
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
	c.JSON(http.StatusOK, gin.H{"Status": true, "Msg":"发送成功"})
}

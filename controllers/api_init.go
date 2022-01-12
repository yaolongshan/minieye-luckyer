package controllers

import (
	"code/minieye-luckyer/models/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ApiDBInit 初始化数据表
func ApiDBInit(c *gin.Context) {
	err := db.InitTables()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "初始化失败",
			"Error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Msg":    "初始化成功",
	})
}

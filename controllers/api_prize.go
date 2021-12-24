package controllers

import (
	"code/minieye-luckyer/models/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type prize struct {
	Name string `json:"name"`
	Sum  int    `json:"sum"`
}

// ApiAddPrize 添加一个奖项
func ApiAddPrize(c *gin.Context) {
	var p prize
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    err.Error()})
		return
	}
	err := db.AddPrize(p.Name, p.Sum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Msg":    "ok"})
}

// ApiGetAllPrize 奖项列表
func ApiGetAllPrize(c *gin.Context) {
	prizes := db.GetPrizeList()
	c.JSON(http.StatusOK, gin.H{"Prizes": prizes})
}

// ApiUpdatePrize 修改奖项的数量
func ApiUpdatePrize(c *gin.Context) {
	var p prize
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    err.Error()})
		return
	}
	err := db.UpdatePrize(p.Name, p.Sum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Msg":    "ok"})
}

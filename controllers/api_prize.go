package controllers

import (
	"code/minieye-luckyer/comm"
	"code/minieye-luckyer/models/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type prize struct {
	Level       string `json:"level"`
	Name        string `json:"name"`
	Sum         int    `json:"sum"`
	ImageBase64 string `json:"image_base64"`
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
	b, urlOrMsg := comm.Base64SaveImage(p.ImageBase64)
	if !b {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    urlOrMsg})
		return
	}
	err := db.AddPrize(p.Level, p.Name, urlOrMsg, p.Sum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    err.Error()})
		return
	}
	prizeInfo := db.GetPrizeByLevel(p.Level)
	c.JSON(http.StatusOK, gin.H{
		"Status":    true,
		"Msg":       "ok",
		"Prize": prizeInfo})
}

// ApiGetAllPrize 奖项列表
func ApiGetAllPrize(c *gin.Context) {
	prizes := db.GetPrizeList()
	c.JSON(http.StatusOK, gin.H{"Prizes": prizes})
}

// ApiUpdatePrize 修改奖项的数量
func ApiUpdatePrize(c *gin.Context) {
	type req struct {
		Level string `json:"level"`
		Sum   int    `json:"sum"`
	}
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    err.Error()})
		return
	}
	err := db.UpdatePrize(r.Level, r.Sum)
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

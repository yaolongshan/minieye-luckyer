package controllers

import (
	"code/minieye-luckyer/models/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type prize struct {
	Level       string `json:"level"`
	Name        string `json:"name"`
	Sum         int    `json:"sum"`
}

// ApiAddPrize 添加一个奖项
func ApiAddPrize(c *gin.Context) {
	p := &prize{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "json参数错误",
			"Error":  err.Error()})
		return
	}
	//b, urlOrMsg := comm.Base64SaveImage(p.ImageBase64)
	//if !b {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"Status": false,
	//		"Msg":    "图片数据错误",
	//		"Error":  urlOrMsg})
	//	return
	//}
	err = db.AddPrize(p.Level, p.Name, p.Sum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "奖项级别不能重复",
			"Error":  err.Error()})
		return
	}
	prizeInfo := db.GetPrizeByLevel(p.Level)
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Msg":    "ok",
		"Prize":  prizeInfo})
}

// ApiGetAllPrize 奖项列表
func ApiGetAllPrize(c *gin.Context) {
	prizes := db.GetPrizeList()
	count := len(prizes)
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Count":  count,
		"Prizes": prizes})
}

// ApiUpdatePrizeSum 修改奖项的数量
func ApiUpdatePrizeSum(c *gin.Context) {
	type req struct {
		ID  int `json:"id"`
		Sum int `json:"sum"`
	}
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "json参数错误",
			"Error":  err.Error()})
		return
	}
	err := db.UpdatePrizeSum(r.ID, r.Sum)
	if err != nil {
		if err.Error() == "设置的数量不能低于已抽数量" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": false,
				"Msg":    "设置的数量不能低于已抽数量",
				"Error":  "error"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "修改失败",
			"Error":  err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Msg":    "ok"})
}

// ApiUpdatePrizeDrawNumber 修改奖项每次抽奖的数量
func ApiUpdatePrizeDrawNumber(c *gin.Context) {
	type req struct {
		ID     int `json:"id"`
		Number int `json:"number"`
	}
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "json参数错误",
			"Error":  err.Error()})
		return
	}
	if r.Number <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "number参数错误"})
		return
	}
	err := db.UpdatePrizeDrawNumber(r.ID, r.Number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "修改失败",
			"Error":  err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Msg":    "ok"})
}

// ApiDelPrize 删除一个奖项
func ApiDelPrize(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "id只能是非0正整数",
			"Error":  err.Error()})
		return
	}
	if err := db.PrizeDeleteByID(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "删除失败",
			"Error":  err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": true, "Msg": "删除成功"})
}

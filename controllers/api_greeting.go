package controllers

import (
	"code/minieye-luckyer/models/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ApiGreetingList 祝福语列表
func ApiGreetingList(c *gin.Context) {
	list := db.GetAllGreeting()
	count := len(list)
	c.JSON(http.StatusOK, gin.H{"Status": true, "Count": count, "Greetings": list})
}

// ApiAddGreeting 添加
func ApiAddGreeting(c *gin.Context) {
	type req struct {
		Name     string `json:"name"`
		Number   string `json:"number"`
		Greeting string `json:"greeting"`
	}
	var r req
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "json参数错误",
			"Error":  err.Error()})
		return
	}
	err = db.AddGreeting(r.Name, r.Number, r.Greeting)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "添加失败, 不能重复添加",
			"Error":  err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": true, "Msg": "添加成功"})
}

// ApiRandomGreeting 随机抽祝福语
func ApiRandomGreeting() {

}

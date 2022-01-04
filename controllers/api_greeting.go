package controllers

import (
	"code/minieye-luckyer/models/db"
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
	"sync"
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

var GMu sync.Mutex

// ApiRandomGreeting 随机抽祝福语
func ApiRandomGreeting(c *gin.Context) {
	GMu.Lock()
	count, err := strconv.Atoi(c.Query("count")) // 抽奖数量
	if err != nil || count <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "count参数错误",
			"Error":  err.Error()})
		GMu.Unlock()
		return
	}
	type result struct {
		Name     string
		Number   string
		Greeting string
	}
	var results []result
	// 本次抽祝福语的数量
	participants := len(db.GetNotLuckyGreeting())
	if participants == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "请检查有无未抽奖的祝福语"})
		GMu.Unlock()
		return
	}
	for i := 0; i < count; i++ {
		// 获取未中奖的祝福语
		greetings := db.GetNotLuckyGreeting()
		len_ := len(greetings)
		if len_ == 0 {
			break
		}
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len_)))
		greeting := greetings[index.Int64()]
		r := result{
			Name:     greeting.Name,
			Number:   greeting.Number,
			Greeting: greeting.Greeting,
		}
		results = append(results, r)
		//标记这条祝福语已中奖
		db.GreetingHasLucky(int(greeting.ID), true)
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":         true,
		"Count":          len(results),
		"Participants":   participants,
		"Results":        results,
	})
	GMu.Unlock()
}

package controllers

import (
	"code/minieye-luckyer/models/db"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

// ApiGetRandom 随机抽奖，根据每个奖项的可中奖数量，返回中奖人员
func ApiGetRandom(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id")) // 奖项id
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	count, err := strconv.Atoi(c.Query("count")) // 抽奖数量
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	prize := db.GetPrizeByID(id)
	if prize.Sum <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Msg": "该奖项已抽奖完毕"})
		return
	}
	type result struct {
		Name   string
		Phone  string
		Number string
		Mail   string
	}
	var results []result
	for i := 0; i < count; i++ {
		//拿到没中奖的小伙伴
		users := db.GetNotLuckyUserList()
		fmt.Println("www", len(users))
		prize = db.GetPrizeByID(id)
		if prize.Sum <= 0 {
			fmt.Println("抽完咯")
			break
		}
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(users))))
		user := users[index.Int64()]
		r := result{
			Name:   user.Name,
			Phone:  user.Phone,
			Number: user.Number,
			Mail:   user.Mail,
		}
		results = append(results, r)
		//保存到中奖信息
		db.AddLucky(int(user.ID), user.Name, prize.Name)
		//奖项数量递减一下
		db.PrizeDegressive(int(prize.ID))
		//标记一下用户表中的已中奖字段
		db.UserHasLucky(int(user.ID), true)
	}
	c.JSON(http.StatusOK, results)
}

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
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "id参数错误",
			"Error":  err.Error()})
		return
	}
	count, err := strconv.Atoi(c.Query("count")) // 抽奖数量
	if err != nil || count <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "count参数错误",
			"Error":  err.Error()})
		return
	}
	prize := db.GetPrizeByID(id)
	if prize.AlreadyUsed >= prize.Sum {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "该奖项已抽奖完毕"})
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
		prize = db.GetPrizeByID(id)
		if prize.AlreadyUsed >= prize.Sum {
			fmt.Println("抽完咯")
			break
		}
		//拿到没中奖的非实习生小伙伴
		users := db.GetNotLuckyFullTimeUserList()
		lenUser := len(users)
		if lenUser == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": false,
				"Msg":    "请检查有无未抽奖的用户"})
			return
		}
		fmt.Println(fmt.Sprintf("第%v轮共有%v人参与抽奖", i+1, lenUser))
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(lenUser)))
		user := users[index.Int64()]
		r := result{
			Name:   user.Name,
			Phone:  user.Phone,
			Number: user.Number,
			Mail:   user.Mail,
		}
		results = append(results, r)
		//保存到中奖信息
		db.AddLucky(int(user.ID), user.Name, user.Number, user.Phone, user.Mail, prize.Level, prize.Name)
		//奖项已抽数量递增
		db.PrizeIncrease(int(prize.ID))
		//标记一下用户表中的已中奖字段
		db.UserHasLucky(int(user.ID), true)
	}
	prize = db.GetPrizeByID(id)
	c.JSON(http.StatusOK, gin.H{
		"Status":         true,
		"Count":          len(results),
		"Results":        results,
		"PrizeRemaining": prize.Sum - prize.AlreadyUsed,
	})
}

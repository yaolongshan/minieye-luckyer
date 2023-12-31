package controllers

import (
	"code/minieye-luckyer/models/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var lMu sync.Mutex

type result struct {
	Name   string
	Phone  string
	Number string
}

// ApiGetRandom 随机抽奖，根据每个奖项的可中奖数量，返回中奖人员
func ApiGetRandom(c *gin.Context) {
	lMu.Lock()
	id, err := strconv.Atoi(c.Query("id")) // 奖项id
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "id参数错误",
			"Error":  err.Error()})
		lMu.Unlock()
		return
	}
	count, err := strconv.Atoi(c.Query("count")) // 抽奖数量
	if err != nil || count <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "count参数错误",
			"Error":  err.Error()})
		lMu.Unlock()
		return
	}
	prize := db.GetPrizeByID(id)
	if prize.AlreadyUsed >= prize.Sum {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "该奖项已抽奖完毕"})
		lMu.Unlock()
		return
	}
	var results []result
	// 参与抽奖的人数，这里要包括实习生的人数
	var participants = db.GetNotLuckyUserListCount()
	// 全职员工抽奖人数校验
	if db.GetNotLuckyFullTimeUserCount() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "请检查有无未抽奖的用户"})
		lMu.Unlock()
		return
	}
	// 抽奖过程
	for i := 0; i < count; i++ {
		// 剩余奖项数量
		prize = db.GetPrizeByID(id)
		if prize.AlreadyUsed >= prize.Sum {
			fmt.Println("抽完咯")
			break
		}
		//拿到没中奖的非实习生小伙伴
		users := db.GetNotLuckyFullTimeUserList()
		lenUser := len(users)
		if lenUser == 0 {
			break
		}
		fmt.Println(fmt.Sprintf("第%v轮共有%v人参与抽奖", i+1, lenUser))
		//index, _ := rand.Int(rand.Reader, big.NewInt(int64(lenUser)))
		index := rand.Intn(lenUser)
		user := users[index]
		r := result{
			Name:   user.Name,
			Phone:  user.Phone,
			Number: user.Number,
		}
		results = append(results, r)
		//保存到中奖信息
		db.AddLucky(int(user.ID), user.Name, user.Number, user.Phone, prize.Level, prize.Name)
		//奖项已抽数量递增
		db.PrizeIncrease(int(prize.ID))
		//标记一下用户表中的已中奖字段
		db.UserHasLucky(int(user.ID), true)
	}
	prize = db.GetPrizeByID(id)
	c.JSON(http.StatusOK, gin.H{
		"Status":         true,
		"Count":          len(results),
		"Participants":   participants,
		"Results":        results,
		"PrizeRemaining": prize.Sum - prize.AlreadyUsed,
	})
	lMu.Unlock()
}

// ApiGetRandomV2 优化 ApiGetRandom 后的抽奖
func ApiGetRandomV2(c *gin.Context) {
	lMu.Lock()
	id, err := strconv.Atoi(c.Query("id")) // 奖项id
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "id参数错误",
			"Error":  err.Error()})
		lMu.Unlock()
		return
	}
	count, err := strconv.Atoi(c.Query("count")) // 抽奖数量
	if err != nil || count <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "count参数错误",
			"Error":  err.Error()})
		lMu.Unlock()
		return
	}
	prize := db.GetPrizeByID(id)
	if prize.AlreadyUsed >= prize.Sum {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "该奖项已抽奖完毕"})
		lMu.Unlock()
		return
	}
	var results []result
	// 参与抽奖的人数，这里要包括实习生的人数
	var participants = db.GetNotLuckyUserListCount()
	// 全职员工抽奖人数校验
	if db.GetNotLuckyFullTimeUserCount() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "请检查有无未抽奖的用户"})
		lMu.Unlock()
		return
	}
	var lucks []db.TBLucky
	alreadyUsed := 0
	var ids []int
	//拿到没中奖的非实习生小伙伴
	users := db.GetNotLuckyFullTimeUserList()
	// 抽奖过程
	for i := 0; i < count; i++ {
		// 剩余奖项数量
		if prize.AlreadyUsed >= prize.Sum {
			fmt.Println("抽完咯")
			break
		}
		lenUser := len(users)
		if lenUser == 0 {
			break
		}
		fmt.Println(fmt.Sprintf("第%v轮共有%v人参与抽奖", i+1, lenUser))
		//index, _ := rand.Int(rand.Reader, big.NewInt(int64(lenUser)))
		index := rand.Intn(lenUser)
		user := users[index]
		//从数组中移除这位用户
		users = append(users[:index], users[index+1:]...)
		r := result{
			Name:   user.Name,
			Phone:  user.Phone,
			Number: user.Number,
		}
		results = append(results, r)
		//保存中奖记录
		l := db.TBLucky{
			UserID:     int(user.ID),
			Name:       user.Name,
			Number:     user.Number,
			Phone:      user.Phone,
			PrizeLevel: prize.Level,
			Content:    prize.Name,
		}
		lucks = append(lucks, l)
		//保存中奖的用户
		ids = append(ids, int(user.ID))
		//奖项已抽数量递增
		alreadyUsed++
		prize.AlreadyUsed++
	}
	//保存中奖记录
	db.AddLucks(lucks)
	//改变奖项已抽数量
	db.PrizeIncreaseBy(int(prize.ID), alreadyUsed)
	//标记一下用户表中的已中奖字段
	db.UsersHasLucky(ids, true)
	c.JSON(http.StatusOK, gin.H{
		"Status":         true,
		"Count":          len(results),
		"Participants":   participants,
		"Results":        results,
		"PrizeRemaining": prize.Sum - prize.AlreadyUsed,
	})
	lMu.Unlock()
}

// ApiGetRandomV3 根据设置的每次抽奖的数量进行抽奖
func ApiGetRandomV3(c *gin.Context) {
	lMu.Lock()
	id, err := strconv.Atoi(c.Query("id")) // 奖项id
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "id参数错误",
			"Error":  err.Error()})
		lMu.Unlock()
		return
	}
	prize := db.GetPrizeByID(id)
	if prize.AlreadyUsed >= prize.Sum {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "该奖项已抽奖完毕"})
		lMu.Unlock()
		return
	}
	var results []result
	// 参与抽奖的人数，这里要包括实习生的人数
	var participants = db.GetNotLuckyUserListCount()
	// 全职员工抽奖人数校验
	if db.GetNotLuckyFullTimeUserCount() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Msg":    "请检查有无未抽奖的用户"})
		lMu.Unlock()
		return
	}
	var lucks []db.TBLucky
	alreadyUsed := 0
	var ids []int
	//拿到没中奖的非实习生小伙伴
	users := db.GetNotLuckyFullTimeUserList()
	//随机种子
	rand.Seed(time.Now().UnixNano())
	// 抽奖过程
	for i := 0; i < prize.DrawNumber; i++ {
		// 剩余奖项数量
		if prize.AlreadyUsed >= prize.Sum {
			fmt.Println("抽完咯")
			break
		}
		lenUser := len(users)
		if lenUser == 0 {
			break
		}
		//打乱一波
		users = upsetUsers(users, lenUser)
		fmt.Println(fmt.Sprintf("第%v轮共有%v人参与抽奖", i+1, lenUser))
		index := rand.Intn(lenUser)
		user := users[index]
		//从数组中移除这位用户
		users = append(users[:index], users[index+1:]...)
		r := result{
			Name:   user.Name,
			Phone:  user.Phone,
			Number: user.Number,
		}
		results = append(results, r)
		//保存中奖记录
		l := db.TBLucky{
			UserID:     int(user.ID),
			Name:       user.Name,
			Number:     user.Number,
			Phone:      user.Phone,
			PrizeLevel: prize.Level,
			Content:    prize.Name,
		}
		lucks = append(lucks, l)
		//保存中奖的用户
		ids = append(ids, int(user.ID))
		//奖项已抽数量递增
		alreadyUsed++
		prize.AlreadyUsed++
	}
	//保存中奖记录
	db.AddLucks(lucks)
	//改变奖项已抽数量
	db.PrizeIncreaseBy(int(prize.ID), alreadyUsed)
	//标记一下用户表中的已中奖字段
	db.UsersHasLucky(ids, true)
	c.JSON(http.StatusOK, gin.H{
		"Status":         true,
		"Count":          len(results),
		"Participants":   participants,
		"Results":        results,
		"PrizeRemaining": prize.Sum - prize.AlreadyUsed,
	})
	lMu.Unlock()
}

// 打乱用户数组
func upsetUsers(users []db.TBUser, len int) []db.TBUser {
	rand.Shuffle(len, func(i, j int) {
		users[i], users[j] = users[j], users[i]
	})
	rand.Shuffle(len, func(i, j int) {
		users[i], users[j] = users[j], users[i]
	})
	return users
}

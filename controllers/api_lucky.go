package controllers

import (
	"code/minieye-luckyer/comm"
	"code/minieye-luckyer/conf"
	"code/minieye-luckyer/models/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

// ApiGetAllLucky 获取中奖名单
func ApiGetAllLucky(c *gin.Context) {
	list := db.GetLuckyList()
	count := len(list)
	c.JSON(http.StatusOK, gin.H{
		"Status":    true,
		"Count":     count,
		"LuckyList": list})
}

// ApiGetLuckyFile 下载中奖名单表格文件
func ApiGetLuckyFile(c *gin.Context) {
	comm.CreateLuckyXLSXFile()
	filePath := conf.Conf.RootPath + "/files/info.xlsx"
	//fileTmp, _ := os.Open(filePath)
	//defer fileTmp.Close()
	fileName := path.Base(filePath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
}

// ApiGetNotLucky 获取未中奖名单 -> 阳光普照奖
func ApiGetNotLucky(c *gin.Context) {
	users := db.GetNotLuckyUserList()
	count := db.GetNotLuckyUserListCount()
	usersLen := len(users)
	var luckyList = make([]db.TBLucky, usersLen)
	for i := 0; i < usersLen; i++ {
		luckyList[i].ID = uint(i + 1)
		luckyList[i].CreatedAt = users[i].CreatedAt
		luckyList[i].UpdatedAt = users[i].UpdatedAt
		luckyList[i].DeletedAt = users[i].DeletedAt
		luckyList[i].UserID = int(users[i].ID)
		luckyList[i].Name = users[i].Name
		luckyList[i].Number = users[i].Number
		luckyList[i].Phone = users[i].Phone
		luckyList[i].Mail = users[i].Mail
		luckyList[i].PrizeLevel = "阳光普照奖"
		luckyList[i].Content = "京东卡/沃尔玛购物卡"
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":    true,
		"Count":     count,
		"LuckyList": luckyList})
}

// ApiGetNotLuckyFile 下载未中奖名单表格文件
func ApiGetNotLuckyFile(c *gin.Context) {
	comm.CreateNotLuckyXLSXFile()
	filePath := conf.Conf.RootPath + "/files/not.xlsx"
	fileName := path.Base(filePath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
}

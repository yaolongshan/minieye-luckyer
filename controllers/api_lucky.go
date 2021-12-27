package controllers

import (
	"code/minieye-luckyer/conf"
	"code/minieye-luckyer/models/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

// ApiGetAllLucky 获取中奖名单
func ApiGetAllLucky(c *gin.Context) {
	list := db.GetLuckyList()
	count := db.LuckyCount()
	c.JSON(http.StatusOK, gin.H{
		"Status":    true,
		"Count":     count,
		"LuckyList": list})
}

// ApiGetLuckyFile 下载中奖名单表格文件
func ApiGetLuckyFile(c *gin.Context) {
	filePath := conf.Conf.RootPath + "/user.xlsx"
	fileTmp, _ := os.Open(filePath)
	defer fileTmp.Close()
	fileName := path.Base(filePath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
}

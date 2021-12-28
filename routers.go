package main

import (
	api "code/minieye-luckyer/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var app *gin.Engine

func init() {
	app = gin.Default()
	//app.Use(MiddleWare())
	v1 := app.Group("/api")
	{
		//员工列表
		v1.GET("/user/list", api.ApiGetAllUser)
		//添加员工
		v1.POST("/user/add", api.ApiAddUser)
		//添加奖项
		v1.POST("/prize/add", api.ApiAddPrize)
		//奖项列表
		v1.GET("/prize/list", api.ApiGetAllPrize)
		//设置奖项数量
		v1.POST("/prize/update", api.ApiUpdatePrize)
		//删除一个奖项
		v1.DELETE("/prize/delete", api.ApiDelPrize)
		//中奖名单列表
		v1.GET("/lucky/list", api.ApiGetAllLucky)
		//下载中奖名单表格文件
		v1.GET("/lucky/file", api.ApiGetLuckyFile)
		//未中奖(阳光普照奖)名单列表
		v1.GET("/lucky/notlist", api.ApiGetNotLucky)
		//下载未中奖(阳光普照奖)名单表格文件
		v1.GET("/lucky/notfile",api.ApiGetNotLuckyFile)
		//抽奖接口
		v1.GET("/lucky/random", api.ApiGetRandom)
		//图片预览
		v1.GET("/images/:name", api.ApiImagesPreview)
		//短信接口
		v1.POST("/sms/send", api.ApiSendSMS)
		//钉钉消息通知接口
		v1.POST("/ding/send", api.ApiSendDingDing)
		//抽祝福语 greetings
		//v1.GET("/greetings/random")
	}
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

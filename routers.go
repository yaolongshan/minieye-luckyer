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
	//前台接口
	web := app.Group("/api/web")
	{
		web.GET("/test", func(c *gin.Context) {

		})
	}
	//后台接口
	admin := app.Group("/api/admin")
	{
		admin.GET("/user/list", api.ApiGetAllUser)
		admin.POST("/user/add", api.ApiAddUser)
		admin.POST("/prize/add", api.ApiAddPrize)
		admin.GET("/prize/list", api.ApiGetAllPrize)
		admin.POST("/prize/update", api.ApiUpdatePrize)
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

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var app *gin.Engine

func init() {
	app = gin.Default()
	//app.Use(MiddleWare())
	//前台接口
	web := app.Group("/api/web")
	{
		web.GET("/add", func(c *gin.Context) {
			c.String(http.StatusOK, "hhh")
		})
	}
	//后台接口
	admin := app.Group("/api/admin")
	{
		admin.GET("/add", func(c *gin.Context) {
			c.String(http.StatusOK, "www")
		})
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

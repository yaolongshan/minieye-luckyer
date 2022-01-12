package router

import (
	api "code/minieye-luckyer/controllers"
	"code/minieye-luckyer/controllers/auth"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	app := gin.Default()
	v1 := app.Group("/api")
	v1.Use(auth.JWTAuth)
	{
		//登录
		v1.POST("/login", auth.Login)
		//判断用户是否登录
		v1.GET("/islogin", auth.IsLogin)
		//员工列表
		v1.GET("/user/list", api.ApiGetAllUser)
		//添加员工
		v1.POST("/user/add", api.ApiAddUser)
		//添加奖项
		v1.POST("/prize/add", api.ApiAddPrize)
		//奖项列表
		v1.GET("/prize/list", api.ApiGetAllPrize)
		//设置奖项数量
		v1.POST("/prize/update", api.ApiUpdatePrizeSum)
		//修改奖项每次抽奖的数量
		v1.POST("/prize/change", api.ApiUpdatePrizeDrawNumber)
		//删除一个奖项
		v1.DELETE("/prize/delete", api.ApiDelPrize)
		//中奖名单列表
		v1.GET("/lucky/list", api.ApiGetAllLucky)
		//下载中奖名单表格文件
		v1.GET("/lucky/file", api.ApiGetLuckyFile)
		//未中奖(阳光普照奖)名单列表
		v1.GET("/lucky/notlist", api.ApiGetNotLucky)
		//下载未中奖(阳光普照奖)名单表格文件
		v1.GET("/lucky/notfile", api.ApiGetNotLuckyFile)
		//抽奖接口
		v1.GET("/lucky/random", api.ApiGetRandomV3)
		//图片预览
		v1.GET("/images/:name", api.ApiImagesPreview)
		//短信接口
		v1.POST("/sms/send", api.ApiSendSMS)
		//钉钉消息通知接口
		v1.POST("/ding/send", api.ApiSendDingDing)
		//设置通知功能开关状态
		v1.GET("/notify/set", api.ApiNotifySet)
		//获取通知功能开关状态
		v1.GET("/notify/get", api.ApiNotifyGet)
		//祝福语列表
		v1.GET("/greeting/list", api.ApiGreetingList)
		//添加一条祝福语
		v1.POST("/greeting/add", api.ApiAddGreeting)
		//抽祝福语 greetings
		v1.GET("/greeting/random", api.ApiRandomGreeting)
		//获取中奖的祝福语
		v1.GET("/greeting/luckylist", api.ApiGreetingLuckyList)
		//下载中奖的祝福语表格文件
		v1.GET("/greeting/file", api.ApiLuckyGreetingFile)
		//测试用，初始化数据库数据
		v1.GET("/db/init", api.ApiDBInit)
	}
	return app
}

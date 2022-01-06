package main

import (
	"code/minieye-luckyer/conf"
	"code/minieye-luckyer/router"
	"fmt"
)

func main() {
	Init()
	app := router.SetupRouter()
	app.Run(fmt.Sprintf(":%v", conf.Conf.Port))
}

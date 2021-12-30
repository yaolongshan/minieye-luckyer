package main

import (
	"code/minieye-luckyer/conf"
	"fmt"
)

func main() {
	Init()
	app.Run(fmt.Sprintf(":%v", conf.Conf.Port))
}
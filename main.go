package main

import (
	"code/minieye-luckyer/conf"
	"fmt"
)

func main() {
	start()
}

func start(){
	app.Run(fmt.Sprintf(":%v", conf.Conf.Port))
}

func readUserFile(){

}

func init(){
	Init()
}

package main

import (
	"code/minieye-luckyer/conf"
	"code/minieye-luckyer/models/db"
)

func main() {
	conf.LoadLocalConf()
	db.InitDB()
	//db.ReadFile()
	app.Run()
}

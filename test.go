package main

import (
	"code/minieye-luckyer/comm"
	"code/minieye-luckyer/conf"
	"code/minieye-luckyer/models/db"
)

func main() {
	conf.LoadLocalConf()
	db.InitDB()
	comm.CreateXLSXFile()
}
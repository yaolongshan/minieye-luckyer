package main

import (
	"code/minieye-luckyer/comm"
	"code/minieye-luckyer/conf"
	"code/minieye-luckyer/models/db"
	"fmt"
	"os"
)

func Init() {
	conf.LoadLocalConf()
	db.InitDB()
	// images dir
	imagePath := fmt.Sprintf("%v/images", conf.Conf.RootPath)
	if !comm.IsFileExist(imagePath) {
		err := os.Mkdir(imagePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	//db.ReadFile()
}

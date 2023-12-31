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
	// dir
	imagePath := fmt.Sprintf("%v/images", conf.Conf.RootPath)
	if !comm.IsFileExist(imagePath) {
		err := os.Mkdir(imagePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	filesPath := fmt.Sprintf("%v/files", conf.Conf.RootPath)
	if !comm.IsFileExist(filesPath) {
		err := os.Mkdir(filesPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

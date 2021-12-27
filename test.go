package main

import (
	"code/minieye-luckyer/comm"
	"code/minieye-luckyer/conf"
)

func main() {
	conf.LoadLocalConf()
	comm.CreateXLSXFile()
}
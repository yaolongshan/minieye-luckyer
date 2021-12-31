package main

import (
	"code/minieye-luckyer/conf"
	"code/minieye-luckyer/models/db"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	conf.LoadLocalConf()
	db.InitDB()
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "初始化本地数据库data.db",
			Action: func(c *cli.Context) error {
				return db.InitTables()
			},
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "从xlsx文件中读取员工信息到数据库中",
			Action: func(c *cli.Context) error {
				path := c.Args().Get(0)
				return db.ReadUserFile(path)
			},
		},
		{
			Name:    "greeting",
			Aliases: []string{"g"},
			Usage:   "从xlsx文件中读取祝福语信息到数据库中",
			Action: func(c *cli.Context) error {
				path := c.Args().Get(0)
				return db.ReadGreetingFile(path)
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

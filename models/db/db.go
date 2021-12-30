package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tealeg/xlsx"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("sqlite3", "./data.db")
	if err != nil {
		panic(fmt.Sprintf("failed to connect database,error: %v", err))
	}
	db.AutoMigrate(&TBUser{})
	db.AutoMigrate(&TBPrize{})
	db.AutoMigrate(&TBLucky{})
	db.AutoMigrate(&TBGreeting{})
}

func ReadUserFile(path string) error {
	file, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Println("open failed:", err)
		return err
	}
	for i, row := range file.Sheets[0].Rows {
		if i == 0 {
			continue
		}
		var user TBUser
		for i, cell := range row.Cells {
			text := cell.Value
			switch i {
			case 0:
				//fmt.Print("姓名：", text," ")
				user.Name = text
			case 1:
				//fmt.Print("手机：", text," ")
				user.Phone = text
			case 2:
				//fmt.Print("类型：", text," ")
				user.Type = text
			case 3:
				//fmt.Print("工号：", text," ")
				user.Number = text
			case 4:
				//fmt.Print("公司：", text," ")
				user.Contract = text
			case 5:
				//fmt.Print("邮箱：", text," ")
				user.Mail = text
			}
			user.IsLucky = false
		}
		err := db.Create(&user).Error
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("成功添加用户: %v ", user.Name))
	}
	return nil
}

func ReadGreetingFile(path string) error {
	file, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Println("open failed:", err)
		return err
	}
	for i, row := range file.Sheets[0].Rows {
		if i == 0 {
			continue
		}
		var greet TBGreeting
		for i, cell := range row.Cells {
			text := cell.Value
			switch i {
			case 0:
				//fmt.Print("姓名：", text, " ")
				greet.Name = text
			case 1:
				//fmt.Print("工号：", text, " ")
				greet.Number = text
			case 2:
				//fmt.Print("祝福语：", text, " ")
				greet.Greeting = text
			}
		}
		err := db.Create(&greet).Error
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("成功添加%v的祝福语", greet.Name))
	}
	return nil
}

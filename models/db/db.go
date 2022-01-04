package db

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("./data.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database,error: %v", err))
	}
	_ = db.AutoMigrate(&TBUser{})
	_ = db.AutoMigrate(&TBPrize{})
	_ = db.AutoMigrate(&TBLucky{})
	_ = db.AutoMigrate(&TBGreeting{})
}

// InitTables 测试用，初始化表，删除数据
func InitTables() error {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&TBUser{}).Update("is_lucky", false)

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&TBGreeting{}).Update("is_lucky", false)

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&TBPrize{}).Update("already_used", 0)

	err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&TBLucky{}).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadUserFile(path string) error {
	file, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Println("open failed:", err)
		return err
	}
	var users []TBUser
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
		}
		user.IsLucky = false
		users = append(users, user)
	}
	err = db.Create(&users).Error
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("成功添加%v位用户", len(users)))
	return nil
}

func ReadGreetingFile(path string) error {
	file, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Println("open failed:", err)
		return err
	}
	var greets []TBGreeting
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
		greet.IsLucky = false
		greets = append(greets, greet)
	}
	err = db.Create(&greets).Error
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("成功添加%v条祝福语", len(greets)))
	return nil
}

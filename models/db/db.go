package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
}

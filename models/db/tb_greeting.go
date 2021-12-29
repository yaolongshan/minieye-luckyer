package db

import "github.com/jinzhu/gorm"

// TBGreeting 祝福语，跟其他表没有关联，是单独的一个抽奖
type TBGreeting struct {
	gorm.Model
	Name     string
	Number   string
	Greeting string
}



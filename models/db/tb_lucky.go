package db

import (
	"github.com/jinzhu/gorm"
)

// TBLucky 中奖名单
type TBLucky struct {
	gorm.Model
	UserID     int
	Name       string // 员工姓名
	PrizeLevel string // 中奖的奖项级别
}

// AddLucky 新加一个中奖名单
func AddLucky(userID int, name, prizeLevel string) {
	l := &TBLucky{
		UserID:     userID,
		Name:       name,
		PrizeLevel: prizeLevel,
	}
	db.Create(&l)
}

// QueryLucky 查询某人是否中奖
func QueryLucky(userID int) bool {
	var l TBLucky
	db.Where(&TBLucky{UserID: userID}).Find(&l)
	return l.Name != ""
}

// GetLuckyList 获取所有中奖名单
func GetLuckyList() (ls []TBLucky) {
	db.Find(&ls)
	return ls
}

// LuckyCount 中奖列表数量
func LuckyCount() (count int) {
	db.Model(&TBLucky{}).Count(&count)
	return count
}

func (TBLucky) TableName() string {
	return "tb_lucky"
}

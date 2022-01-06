package db

import (
	"gorm.io/gorm"
)

// TBLucky 中奖名单
type TBLucky struct {
	gorm.Model
	UserID     int
	Name       string // 员工姓名
	Number     string // 工号
	Phone      string
	Mail       string
	PrizeLevel string // 中奖的奖项级别
	Content    string // 奖品内容
}

// AddLucky 新加一个中奖名单
func AddLucky(userID int, name, number, phone, mail, prizeLevel, content string) {
	l := &TBLucky{
		UserID:     userID,
		Name:       name,
		Number:     number,
		Phone:      phone,
		Mail:       mail,
		PrizeLevel: prizeLevel,
		Content:    content,
	}
	db.Create(&l)
}

// AddLucks 添加多个中奖记录
func AddLucks(lucks []TBLucky) {
	db.Create(&lucks)
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
func LuckyCount() (count int64) {
	db.Model(&TBLucky{}).Count(&count)
	return count
}

func (TBLucky) TableName() string {
	return "tb_lucky"
}

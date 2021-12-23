package db

import (
	"github.com/jinzhu/gorm"
)

// TBLucky 中奖名单
type TBLucky struct {
	gorm.Model
	UserID    int
	Name      string // 员工姓名
	PrizeName string // 中奖的奖项名称
}

// AddLucky 新加一个中奖名单
func AddLucky(UserID int, Name, PrizeName string) {
	l := &TBLucky{
		UserID:    UserID,
		Name:      Name,
		PrizeName: PrizeName,
	}
	db.Create(&l)
}

// QueryLucky 查询某人是否中奖
func QueryLucky(UserID int) bool {
	var l TBLucky
	db.Where(&TBLucky{UserID: UserID}).Find(&l)
	return l.Name != ""
}

// GetLuckyList 获取所有中奖名单
func GetLuckyList() (ls []TBLucky) {
	db.Find(&ls)
	return ls
}

func (TBLucky) TableName() string {
	return "tb_lucky"
}

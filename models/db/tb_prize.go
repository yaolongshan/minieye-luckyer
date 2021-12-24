package db

import (
	"github.com/jinzhu/gorm"
)

// TBPrize 奖项设置表
type TBPrize struct {
	gorm.Model
	Level    string `gorm:"unique;not null"` // 奖项级别
	Name     string // 奖品名称
	Sum      int    // 奖项数量
	ImageUrl string
}

// GetPrizeList 奖项列表
func GetPrizeList() (prizes []TBPrize) {
	db.Find(&prizes)
	return prizes
}

// UpdatePrize 修改奖项的数量
func UpdatePrize(level string, sum int) error {
	err := db.Model(&TBPrize{}).Where("level = ?", level).Update("sum", sum).Error
	return err
}

// AddPrize 添加一个奖项
func AddPrize(level, name, url string, sum int) error {
	p := &TBPrize{
		Level:    level,
		Name:     name,
		Sum:      sum,
		ImageUrl: url,
	}
	err := db.Create(&p).Error
	return err
}

// GetPrizeByID 获取一个奖项的信息
func GetPrizeByID(id int) (p TBPrize) {
	db.Where("id = ?", id).Find(&p)
	return p
}

// GetPrizeByLevel 获取一个奖项的信息
func GetPrizeByLevel(level string) (p TBPrize) {
	db.Where("level = ?", level).Find(&p)
	return p
}

// PrizeDegressive 让这个奖项的数量递减
func PrizeDegressive(id int) {
	prize := GetPrizeByID(id)
	db.Model(&TBPrize{}).Where("id = ?", id).Update("sum", prize.Sum-1)
}

func (TBPrize) TableName() string {
	return "tb_prize"
}

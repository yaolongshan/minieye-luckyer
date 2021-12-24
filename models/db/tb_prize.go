package db

import (
	"github.com/jinzhu/gorm"
)

// TBPrize 奖项设置表
type TBPrize struct {
	gorm.Model
	Name string `gorm:"unique;not null"` // 奖项名称 特等奖、一等奖
	Sum  int    // 奖项数量
}

// GetPrizeList 奖项列表
func GetPrizeList() (prizes []TBPrize) {
	db.Find(&prizes)
	return prizes
}

// UpdatePrize 修改奖项的数量
func UpdatePrize(name string, sum int) error {
	err := db.Model(&TBPrize{}).Where("name = ?", name).Update("sum", sum).Error
	return err
}

// AddPrize 添加一个奖项
func AddPrize(name string, sum int) error {
	p := &TBPrize{
		Name: name,
		Sum:  sum,
	}
	err := db.Create(&p).Error
	return err
}

// GetPrizeByID 获取一个奖项的信息
func GetPrizeByID(id int) (p TBPrize) {
	db.Where("id = ?", id).Find(&p)
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

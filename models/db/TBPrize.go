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
func UpdatePrize(Name string, Sum int) error {
	err := db.Model(&TBPrize{}).Where("name = ?", Name).Update("sum", Sum).Error
	return err
}

// AddPrize 添加一个奖项
func AddPrize(Name string, Sum int) error {
	p := TBPrize{
		Name: Name,
		Sum:  Sum,
	}
	err := db.Create(&p).Error
	return err
}

func (TBPrize) TableName() string {
	return "tb_prize"
}

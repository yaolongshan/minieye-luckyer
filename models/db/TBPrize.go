package db

import "github.com/jinzhu/gorm"

// TBPrize 奖项表
type TBPrize struct {
	gorm.Model
}

func (TBPrize) TableName() string {
	return "tb_prize"
}

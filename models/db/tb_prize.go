package db

import (
	"errors"
	"gorm.io/gorm"
)

// TBPrize 奖项设置表
type TBPrize struct {
	gorm.Model
	Level       string `gorm:"unique;not null"` // 奖项级别
	Name        string // 奖品名称
	Sum         int    // 奖项数量
	AlreadyUsed int    // 已抽数量
	ImageUrl    string
}

// GetPrizeList 奖项列表
func GetPrizeList() (prizes []TBPrize) {
	db.Find(&prizes)
	return prizes
}

// PrizeCount 数量
func PrizeCount() (count int64) {
	db.Model(&TBPrize{}).Count(&count)
	return count
}

// UpdatePrize 修改奖项的总数量
func UpdatePrize(id, sum int) error {
	prize := GetPrizeByID(id)
	// 添加的情况
	if sum >= prize.Sum {
		err := db.Model(&TBPrize{}).Where("id = ?", id).Update("sum", sum).Error
		return err
	} else {
		// 减少的情况
		// 低于已抽数量时
		if sum < prize.AlreadyUsed {
			e := errors.New("设置的数量不能低于已抽数量")
			return e
		}
		err := db.Model(&TBPrize{}).Where("id = ?", id).Update("sum", sum).Error
		return err
	}
}

// AddPrize 添加一个奖项
func AddPrize(level, name, url string, sum int) error {
	p := &TBPrize{
		Level:       level,
		Name:        name,
		Sum:         sum,
		AlreadyUsed: 0,
		ImageUrl:    url,
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

//// PrizeDegressive 让这个奖项的剩余数量递减
//func PrizeDegressive(id int) {
//	prize := GetPrizeByID(id)
//	db.Model(&TBPrize{}).Where("id = ?", id).Update("remaining", prize.Remaining-1)
//}

// PrizeDeleteByID 删除一个
func PrizeDeleteByID(id int) error {
	err := db.Unscoped().Where("id = ?", id).Delete(&TBPrize{}).Error
	return err
}

// PrizeIncrease 奖项已抽数量递增
func PrizeIncrease(id int) {
	prize := GetPrizeByID(id)
	db.Model(&TBPrize{}).Where("id = ?", id).Update("already_used", prize.AlreadyUsed+1)
}

// PrizeIncreaseBy 根据数量对奖项已抽数量减少
func PrizeIncreaseBy(id, count int) {
	prize := GetPrizeByID(id)
	db.Model(&TBPrize{}).Where("id = ?", id).Update("already_used", prize.AlreadyUsed+count)
}

func (TBPrize) TableName() string {
	return "tb_prize"
}

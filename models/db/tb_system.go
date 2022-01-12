package db

import "gorm.io/gorm"

// TBSystem 系统设置
type TBSystem struct {
	gorm.Model
	Key   string `gorm:"unique"`
	Value string
}

func SetValue(key, value string) {
	db.Model(&TBSystem{}).Where("key = ?", key).Update("value", value)
}

func GetValue(key string) (sys TBSystem) {
	db.Where("key = ?", key).Find(&sys)
	return sys
}

func (TBSystem) TableName() string {
	return "tb_system"
}

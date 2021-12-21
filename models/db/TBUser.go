package db

import "github.com/jinzhu/gorm"

type TBUser struct {
	gorm.Model
}

func (TBUser) TableName() string {
	return "tb_user"
}

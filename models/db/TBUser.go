package db

import (
	"github.com/jinzhu/gorm"
)

type TBUser struct {
	gorm.Model
	Name     string // 姓名
	Phone    string // 手机号
	Type     string // 员工类型，实习、全职
	Number   string // 工号
	Contract string // 合同公司
	Mail     string // 邮箱
}

// UserCount 员工数量
func UserCount() (count int) {
	db.Model(&TBUser{}).Count(&count)
	return count
}

// GetUserList 所有员工
func GetUserList() (users []TBUser) {
	db.Find(&users)
	return users
}

// AddUser 添加员工
func AddUser(Name, Phone, Type, Number, Contract, Mail string) {
	user := &TBUser{
		Name:     Name,
		Phone:    Phone,
		Type:     Type,
		Number:   Number,
		Contract: Contract,
		Mail:     Mail,
	}
	db.Create(&user)
}

func (TBUser) TableName() string {
	return "tb_user"
}

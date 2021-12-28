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
	IsLucky  bool   // 是否中过奖
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

// GetNotLuckyUserList 获取未中奖的员工
func GetNotLuckyUserList() (users []TBUser) {
	db.Where("is_lucky = ?", false).Find(&users)
	return users
}

// AddUser 添加员工
func AddUser(name, phone, type_, number, contract, mail string) {
	user := &TBUser{
		Name:     name,
		Phone:    phone,
		Type:     type_,
		Number:   number,
		Contract: contract,
		Mail:     mail,
	}
	db.Create(&user)
}

// UserHasLucky 标记员工中过奖
func UserHasLucky(id int, is bool) {
	db.Model(&TBUser{}).Where("id = ?", id).Update(&TBUser{IsLucky: is})
}

func (TBUser) TableName() string {
	return "tb_user"
}

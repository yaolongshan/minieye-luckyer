package db

import (
	"github.com/jinzhu/gorm"
)

type TBUser struct {
	gorm.Model
	Name     string `gorm:"unique"` // 姓名
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
	db.Order("number").Find(&users)
	return users
}

// GetNotLuckyUserList 获取未中奖的员工
func GetNotLuckyUserList() (users []TBUser) {
	db.Where("is_lucky = ?", false).Find(&users)
	return users
}

// GetNotLuckyUserListCount 获取未中奖的员工数量
func GetNotLuckyUserListCount() (count int) {
	db.Model(&TBUser{}).Where("is_lucky = ?", false).Count(&count)
	return count
}

// GetNotLuckyFullTimeUserList 获取未中奖的全职员工
func GetNotLuckyFullTimeUserList() (users []TBUser) {
	db.Where("is_lucky = ?", false).Where("type LIKE ?", "%全职%").Find(&users)
	return users
}

// GetNotLuckyFullTimeUserCount 获取未中奖的全职员工数量
func GetNotLuckyFullTimeUserCount() (count int) {
	db.Model(&TBUser{}).Where("is_lucky = ?", false).Where("type LIKE ?", "%全职%").Count(&count)
	return count
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

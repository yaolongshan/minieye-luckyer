package db

import "github.com/jinzhu/gorm"

type TBUser struct {
	gorm.Model
	Name     string // 姓名
	Phone    string // 手机号
	Type     string // 员工类型，实习、全职
	Number   string // 工号
	Contract string // 合同公司
	Mail     string // 邮箱
}

func (TBUser) TableName() string {
	return "tb_user"
}

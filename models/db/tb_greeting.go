package db

import "github.com/jinzhu/gorm"

// TBGreeting 祝福语，跟其他表没有关联，是单独的一个抽奖
type TBGreeting struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Number   string
	Greeting string
	IsLucky  bool // 是否中奖
}

// GreetingCount 祝福语数量
func GreetingCount() (count int) {
	db.Model(&TBGreeting{}).Count(&count)
	return count
}

// GetAllGreeting 所有祝福语
func GetAllGreeting() (list []TBGreeting) {
	db.Order("number").Find(&list)
	return list
}

// GetNotLuckyGreeting 获取未中奖的祝福语列表
func GetNotLuckyGreeting() (greetings []TBGreeting) {
	db.Where("is_lucky = ?", false).Find(&greetings)
	return greetings
}

// GreetingHasLucky 标记祝福语中奖
func GreetingHasLucky(id int, is bool) {
	db.Model(&TBGreeting{}).Where("id = ?", id).Update(&TBGreeting{IsLucky: is})
}

// AddGreeting 添加一条祝福语
func AddGreeting(name, number, greeting string) error {
	g := &TBGreeting{
		Name:     name,
		Number:   number,
		Greeting: greeting,
		IsLucky:  false,
	}
	err := db.Create(&g).Error
	return err
}

func (TBGreeting) TableName() string {
	return "tb_greeting"
}

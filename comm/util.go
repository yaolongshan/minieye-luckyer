package comm

import (
	"code/minieye-luckyer/conf"
	"code/minieye-luckyer/models/db"
	"encoding/base64"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"time"
)

func RandName() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	s := string(result)
	return s
}

// CreateLuckyXLSXFile 中奖名单文件
func CreateLuckyXLSXFile() {
	list := db.GetLuckyList()
	savePath := fmt.Sprintf("%v/files/info.xlsx", conf.Conf.RootPath)
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	{
		row := sheet.AddRow()
		row.SetHeightCM(0.8) //设置每行的高度
		cell := row.AddCell()
		cell.Value = "姓名"
		cell = row.AddCell()
		cell.Value = "工号"
		cell = row.AddCell()
		cell.Value = "手机号"
		cell = row.AddCell()
		cell.Value = "邮箱"
		cell = row.AddCell()
		cell.Value = "奖项等级"
		cell = row.AddCell()
		cell.Value = "奖品内容"
	}

	for _, l := range list {
		row := sheet.AddRow()
		row.SetHeightCM(0.8)
		cell := row.AddCell()
		cell.Value = l.Name
		cell = row.AddCell()
		cell.Value = l.Number
		cell = row.AddCell()
		cell.Value = l.Phone
		cell = row.AddCell()
		cell.Value = l.Mail
		cell = row.AddCell()
		cell.Value = l.PrizeLevel
		cell = row.AddCell()
		cell.Value = l.Content
	}

	err := file.Save(savePath)
	if err != nil {
		panic(err)
	}
}

// CreateNotLuckyXLSXFile 未中奖名单文件
func CreateNotLuckyXLSXFile() {
	users := db.GetNotLuckyUserList()
	usersLen := len(users)
	var luckyList = make([]db.TBLucky, usersLen)
	for i := 0; i < usersLen; i++ {
		luckyList[i].ID = uint(i + 1)
		luckyList[i].CreatedAt = users[i].CreatedAt
		luckyList[i].UpdatedAt = users[i].UpdatedAt
		luckyList[i].DeletedAt = users[i].DeletedAt
		luckyList[i].UserID = int(users[i].ID)
		luckyList[i].Name = users[i].Name
		luckyList[i].Number = users[i].Number
		luckyList[i].Phone = users[i].Phone
		luckyList[i].Mail = users[i].Mail
		luckyList[i].PrizeLevel = "阳光普照奖"
		luckyList[i].Content = "京东卡/沃尔玛购物卡"
	}
	savePath := fmt.Sprintf("%v/files/not.xlsx", conf.Conf.RootPath)
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	{
		row := sheet.AddRow()
		row.SetHeightCM(0.8) //设置每行的高度
		cell := row.AddCell()
		cell.Value = "姓名"
		cell = row.AddCell()
		cell.Value = "工号"
		cell = row.AddCell()
		cell.Value = "手机号"
		cell = row.AddCell()
		cell.Value = "邮箱"
		cell = row.AddCell()
		cell.Value = "奖项等级"
		cell = row.AddCell()
		cell.Value = "奖品内容"
	}

	for _, l := range luckyList {
		row := sheet.AddRow()
		row.SetHeightCM(0.8)
		cell := row.AddCell()
		cell.Value = l.Name
		cell = row.AddCell()
		cell.Value = l.Number
		cell = row.AddCell()
		cell.Value = l.Phone
		cell = row.AddCell()
		cell.Value = l.Mail
		cell = row.AddCell()
		cell.Value = l.PrizeLevel
		cell = row.AddCell()
		cell.Value = l.Content
	}

	err := file.Save(savePath)
	if err != nil {
		panic(err)
	}
}

// CreateLuckyGreetingXLSXFile 祝福语中奖名单文件
func CreateLuckyGreetingXLSXFile() {
	greetings := db.GetLuckyGreeting()
	savePath := fmt.Sprintf("%v/files/greeting.xlsx", conf.Conf.RootPath)
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	{
		row := sheet.AddRow()
		row.SetHeightCM(0.8) //设置每行的高度
		cell := row.AddCell()
		cell.Value = "姓名"
		cell = row.AddCell()
		cell.Value = "工号"
		cell = row.AddCell()
		cell.Value = "祝福语"
	}
	for _, g := range greetings {
		row := sheet.AddRow()
		row.SetHeightCM(0.8)
		cell := row.AddCell()
		cell.Value = g.Name
		cell = row.AddCell()
		cell.Value = g.Number
		cell = row.AddCell()
		cell.Value = g.Greeting
	}
	err := file.Save(savePath)
	if err != nil {
		panic(err)
	}
}

// Base64SaveImage 图片保存
func Base64SaveImage(base64Content string) (bool, string) {
	// base64内容校验
	b, err := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64Content)
	if !b {
		return false, err.Error()
	}
	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	allData := re.FindAllSubmatch([]byte(base64Content), 2)
	//png ，jpeg 后缀获取
	fileType := string(allData[0][1])
	base64Str := re.ReplaceAllString(base64Content, "")
	//8ndvlpqndkqte0t7.jpg
	fileName := fmt.Sprintf("%v.%v", RandName(), fileType)
	//完整路径
	filePath := fmt.Sprintf("%v/images/%v", conf.Conf.RootPath, fileName)

	byteData, _ := base64.StdEncoding.DecodeString(base64Str)
	err = ioutil.WriteFile(filePath, byteData, 0666)

	//拼接访问图片的url
	imageUrl := fmt.Sprintf("/api/images/%v", fileName)

	if err != nil {
		return false, err.Error()
	}
	return true, imageUrl
}

func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

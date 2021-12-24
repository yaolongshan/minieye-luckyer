package comm

import (
	"code/minieye-luckyer/conf"
	"encoding/base64"
	"fmt"
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
	imageUrl := fmt.Sprintf("%v/api/images/%v", conf.Conf.AccessPath, fileName)

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

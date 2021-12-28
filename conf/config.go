package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	RootPath   string `json:"RootPath"`   // 项目根目录
	AccessPath string `json:"AccessPath"` // 服务访问url根路径
	SMS        sms    `json:"SMS"`        // 短信配置
}

type sms struct {
	AccessKeyId     string `json:"AccessKeyId"`     // 阿里云短信keyID
	AccessKeySecret string `json:"AccessKeySecret"` // KeySecret
	SignName        string `json:"SignName"`        // 短信签名
	TemplateCode    string `json:"TemplateCode"`    // 短信模板号
}

var Conf config

func LoadLocalConf() {
	file, err := os.Open("local_conf.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("local_conf.json decoder failed", err.Error())
		panic(err)
	}
}

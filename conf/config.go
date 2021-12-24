package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	RootPath string `json:"RootPath"`
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
		fmt.Println("Decoder failed", err.Error())
		panic(err)
	}
}

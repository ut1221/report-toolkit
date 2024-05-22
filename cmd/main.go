package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"report-toolkit/cmd/config"
	"report-toolkit/cmd/pkg"
)

func main() {
	dataBytes, err := os.ReadFile("cmd/config.yaml")
	if err != nil {
		log.Fatalln("读取配置文件失败:", err)
		return
	}
	conf := config.Config{}
	if err := yaml.Unmarshal(dataBytes, &conf); err != nil {
		log.Fatalln("解析配置文件失败:", err)
		return
	}
	log.Println("初始化翻译api连接需要时间...")
	pkg.Convert(&conf, &conf.BaiduTranslate)
}

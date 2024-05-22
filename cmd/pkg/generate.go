package pkg

import (
	"fmt"
	"log"
	"report-toolkit/cmd/config"
	"report-toolkit/cmd/utils"
	"strings"
)

// Convert
//
//	@Description: 对header进行翻译
//	@Author PTJ 2024-05-22 23:32:23
//	@param c
//	@param b
func Convert(c *config.Config, b *config.BaiduTranslate) {
	stopSpinner := make(chan bool)
	go utils.PrintSpinner(stopSpinner)
	list := make([][]string, 0)
	Generate(&c.Header, b, &list)
	// 发送信号告诉协程停止
	stopSpinner <- true
	<-stopSpinner
	var result []string
	generateCombinations(list, 0, []string{}, &result)
	//生成Java代码
	GenerateEntity(result, c.Entity)
	//生成json文件
	GenerateJson(result, c)
	//生成sql文件
	GenerateSql(result, c)
}

func Generate(c *config.Header, b *config.BaiduTranslate, list *[][]string) {
	translate, err := Translate(c.Name, "zh", "en", b)
	if err != nil || translate == "" {
		fmt.Printf("\r%s\r", strings.Repeat(" ", 10))
		log.Fatalln("翻译失败，请检查配置文件")
		return
	}
	words := strings.Split(translate, ",")
	var strList []string
	for _, word := range words {
		word = strings.TrimSpace(word)
		var str string
		if len(*list) == 0 {
			str = utils.ToCamelCase2(word)
		} else {
			str = utils.ToCamelCase1(word)
		}
		strList = append(strList, str)
	}
	*list = append(*list, strList)
	if c.Header != nil {
		Generate(c.Header, b, list)
	}
}

func generateCombinations(arrays [][]string, index int, current []string, result *[]string) {
	if index == len(arrays) {
		*result = append(*result, strings.Join(current, ""))
		return
	}
	for _, word := range arrays[index] {
		newCurrent := append([]string{}, current...)
		newCurrent = append(newCurrent, word)
		generateCombinations(arrays, index+1, newCurrent, result)
	}
}

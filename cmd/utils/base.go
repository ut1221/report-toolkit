package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// PostForm
//
//	@Description: 发送http post请求数据为form
//	@Author PTJ 2024-05-22 23:29:12
//	@param url
//	@param data
//	@return string
//	@return error
func PostForm(url string, data url.Values) (string, error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Md5
//
//	@Description: md5加密
//	@Author PTJ 2024-05-22 23:29:22
//	@param src
//	@return string
func Md5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// UrlDecode
//
//	@Description: url解码
//	@Author PTJ 2024-05-22 23:29:28
//	@param str
//	@return string
func UrlDecode(str string) string {
	res, err := url.QueryUnescape(str)
	if err != nil {
		return ""
	}
	return res
}

// ToCamelCase1
//
//	@Description: 转驼峰 首字母大写
//	@Author PTJ 2024-05-22 23:29:40
//	@param input
//	@return string
func ToCamelCase1(input string) string {
	input = strings.ReplaceAll(input, "_", " ")
	words := strings.Fields(input)
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

// ToCamelCase2
//
//	@Description: 转驼峰 首字母小写
//	@Author PTJ 2024-05-22 23:29:51
//	@param input
//	@return string
func ToCamelCase2(input string) string {
	// Replace underscores with spaces to handle both cases
	input = strings.ReplaceAll(input, "_", " ")
	words := strings.Fields(input)
	if len(words) == 0 {
		return ""
	}
	words[0] = strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

// ReadTpl
//
//	@Description: 读取模版文件
//	@Author PTJ 2024-05-22 23:30:01
//	@param tplName 模版文件的名称 如 "entity.tpl"
//	@return *template.Template
func ReadTpl(tplName string) *template.Template {
	tplPath := filepath.Join("template", tplName)
	content, err := os.ReadFile(tplPath)
	if err != nil {
		log.Fatalf("读取模板文件失败:%v\n", err)
	}
	tpl, err := template.New("entity").Parse(string(content))
	if err != nil {
		log.Fatalf("解析模板文件失败:%v\n", err)
	}
	return tpl
}

// WriteTpl
//
//	@Description: 将数据写入模版文件
//	@Author PTJ 2024-05-22 23:30:38
//	@param tpl
//	@param insData 对应模版生成的数据
//	@param name 生成的文件名
func WriteTpl(tpl *template.Template, insData interface{}, name string) {
	var data bytes.Buffer
	if err := tpl.Execute(&data, insData); err != nil {
		log.Fatalf("执行模板时出错: %v", err)
		return
	}
	//对于json文件 输出进行格式化
	if strings.Split(name, ".")[1] == "json" {
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, data.Bytes(), "", "  "); err != nil {
			log.Fatalf("格式化JSON时出错: %v", err)
		}
		data = prettyJSON
	}
	fileName := "gen/" + name
	// 创建文件夹（如果不存在）
	err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	if err != nil {
		log.Fatalf("创建文件夹时出错: %v", err)
		return
	}
	// 输出生成数据到文件
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("创建输出文件时出错: %v", err)
		return
	}
	defer func(outputFile *os.File) {
		_ = outputFile.Close()
	}(outputFile)
	_, err = outputFile.WriteString(data.String())
	if err != nil {
		log.Fatalf("写入输出文件时出错: %v", err)
		return
	}
}

func PrintSpinner(stop chan bool) {
	states := []string{"\033[31m|\033[0m", "\033[32m/\033[0m", "\033[33m-\033[0m", "\033[34m\\\033[0m"} // 使用ANSI转义码打印彩色的旋转符号
	for {
		select {
		case <-stop:
			fmt.Printf("\r%s\r", strings.Repeat(" ", 10))
			log.Println("翻译完成,代码生成中...")
			stop <- true
			return
		default:
			for _, state := range states {
				fmt.Printf("\r%s", state)
				time.Sleep(100 * time.Millisecond)
			}

		}

	}
}

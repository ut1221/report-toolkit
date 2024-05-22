package pkg

import (
	"log"
	"report-toolkit/cmd/config"
	"report-toolkit/cmd/utils"
	"strings"
)

// Column 代表表格的列结构
type Column struct {
	Title     string   `json:"title"`
	Align     string   `json:"align,omitempty"`
	DataIndex string   `json:"dataIndex,omitempty"`
	Key       string   `json:"key,omitempty"`
	Width     int      `json:"width,omitempty"`
	Children  []Column `json:"children,omitempty"`
}

// GenerateJson
//
//	@Description: 生成json文件
//	@Author PTJ 2024-05-22 23:32:11
//	@param res
//	@param c
func GenerateJson(res []string, c *config.Config) {
	var count = 0
	nestedMap := generateNestedMap(&c.Header, res, &count)
	utils.WriteTpl(utils.ReadTpl("json.tpl"), nestedMap, c.Json.Name+".json")
	log.Println("json生成成功")
}

// generateNestedMap
//
//	@Description: 构造json文件数据
//	@Author PTJ 2024-05-22 23:31:46
//	@param c
//	@param res
//	@param count
//	@return []Column
func generateNestedMap(c *config.Header, res []string, count *int) []Column {
	result := make([]Column, 0)
	s := strings.Split(c.Name, ",")
	for _, word := range s {
		column := Column{
			Title:     word,
			Align:     "center",
			DataIndex: "",
			Key:       "",
			Width:     0,
		}
		if c.Header != nil {
			column.Title = word
			column.Children = generateNestedMap(c.Header, res, count)
		} else {
			column.DataIndex = res[*count]
			column.Key = res[*count]
			column.Width = 200
			*count++
		}
		result = append(result, column)
	}
	return result
}

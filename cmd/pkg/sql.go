package pkg

import (
	"log"
	"report-toolkit/cmd/config"
	"report-toolkit/cmd/utils"
	"strings"
)

// QueryData 代表查询的结构
type QueryData struct {
	TableName string
	Columns   []TableColumn
}

type TableColumn struct {
	Field     string
	Condition string
}

// GenerateSql
//
//	@Description: 生成sql文件
//	@Author PTJ 2024-05-22 23:31:24
//	@param res
//	@param c
func GenerateSql(res []string, c *config.Config) {
	s := c.Sql
	l := len(s.Condition)
	columns := make([]TableColumn, 0)
	if len(s.Condition) == 0 {
		log.Fatalln("检查condition条件不能为空")
	}
	if l == 1 {
		for _, v := range res {
			columns = append(columns, TableColumn{
				Field:     v,
				Condition: s.Condition[1],
			})
		}
	} else {
		if l == len(strings.Split(c.Header.Name, ",")) {
			for i, v := range res {
				columns = append(columns, TableColumn{
					Field:     v,
					Condition: s.Condition[i%l],
				})
			}
		} else {
			log.Fatalln("condition条件必须与顶级header的name条数一致")
			return
		}
	}
	data := QueryData{
		TableName: s.TableName,
		Columns:   columns,
	}
	utils.WriteTpl(utils.ReadTpl("sql.tpl"), data, s.TableName+".sql")
	log.Println("sql生成成功")
}

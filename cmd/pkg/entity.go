package pkg

import (
	"log"
	"report-toolkit/cmd/config"
	"report-toolkit/cmd/utils"
)

// Field 代表类的字段
type Field struct {
	Name string
	Type string
}

// Entity 代表填充到模板中的数据
type Entity struct {
	PackageName string `yaml:"packageName"`
	ClassName   string `yaml:"className"`
	Fields      []Field
}

// GenerateEntity
//
//	@Description: 生成实体文件
//	@Author PTJ 2024-05-22 23:32:57
//	@param val
//	@param e
func GenerateEntity(val []string, e config.Entity) {
	tpl := utils.ReadTpl("entity.tpl")
	var fields []Field
	for i := range val {
		fields = append(fields, Field{
			Name: val[i],
			Type: "BigDecimal",
		})
	}
	entity := Entity{
		PackageName: e.PackageName,
		ClassName:   e.ClassName,
		Fields:      fields,
	}
	utils.WriteTpl(tpl, entity, e.ClassName+".java")
	log.Println("Java实体类生成成功")
}

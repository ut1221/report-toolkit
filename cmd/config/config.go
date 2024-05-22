package config

type Config struct {
	Header         Header         `yaml:"header"`
	BaiduTranslate BaiduTranslate `yaml:"baiduTranslate"`
	Entity         Entity         `yaml:"entity"`
	Json           Json           `yaml:"json"`
	Sql            Sql            `yaml:"sql"`
}

type Header struct {
	Name   string  `yaml:"name"`
	Header *Header `yaml:"header,omitempty"`
}

type BaiduTranslate struct {
	AppId  string `yaml:"appId"`
	Secure string `yaml:"secure"`
}

type Entity struct {
	PackageName string `yaml:"packageName"`
	ClassName   string `yaml:"className"`
}

type Json struct {
	Name string `yaml:"name"`
}

type Sql struct {
	TableName string   `yaml:"tableName"`
	Condition []string `yaml:"condition"`
}

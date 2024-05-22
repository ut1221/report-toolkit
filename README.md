# 报表工具
一个针对于大量重复表头的报表的java工具包，可多端使用，可动态生成报表的json头与sql字段和Java实体类，支持多级嵌套，免除报表开发人员手动维护报表字段，减少开发人员开发时间。
## 使用
直接终端运行 ./app-macos 或 ./app-linux 或 ./app.exe 即可
所生成文件在当前路径gen目录下
## 配置示例
```yaml
header: # 顶级表头配置
  name: "小红,小黑,小白,小绿" # 第一级表头名称，包含多个以逗号分隔的字段
  header: # 嵌套的表头配置
    name: "语文,数学,英语,历史,物理" # 第二级表头名称，包含多个以逗号分隔的字段
    header: # 更深层次的嵌套表头配置
      name: "早上,下午,晚上" # 第三级表头名称，包含多个以逗号分隔的字段

entity: # Java实体类的配置
  packageName: "com.example.entity" # 实体类的包名
  className: "Entity" # 实体类的类名
json: # JSON文件生成的配置
  name: "gen" # 生成的JSON文件的目录名称
sql: # SQL查询配置
  tableName: "sys_user" # 查询的表名
  condition: # SQL查询的条件列表
    # 这些条件将用于生成select语句中的过滤条件
    - "if(b.flag != 0, a.data, 0)" # 条件1，如果b.flag不等于0，则取a.data，否则取0
    - "if(c.flag != 0, a.data, 0)" # 条件2，如果c.flag不等于0，则取a.data，否则取0
    - "if(d.flag != 0, a.data, 0)" # 条件3，如果d.flag不等于0，则取a.data，否则取0
    - "if(e.flag != 0, a.data, 0)" # 条件4，如果e.flag不等于0，则取a.data，否则取0
    # 注意：condition的数量可以为1或者与顶级header的name字段数量一致

baiduTranslate: # 百度翻译API的配置
  appId: "20221112001449132" # 百度翻译API的应用ID
  secure: "f1IuAdaiFtZS0gOfNaed" # 百度翻译API的安全密钥
```
## 特殊sql配置
若需生成的sql不满足当前配置，可自行修改sql模版配置，sql模版配置解析
```tpl
SELECT
{{- range $index, $column := .Columns }}  // 遍历 .Columns（即根上下文中的列集合）
    {{- if $index }}, {{ end }}  // 如果不是第一个元素，添加逗号以分隔SUM函数调用
    SUM({{ $column.Condition }}) as `{{ $column.Field }}`  // 添加SUM函数调用，并将其结果命名为字段名
{{- end }}  // 结束遍历
FROM {{ .TableName }};  // 添加FROM子句，指定表名
```
例如需要将生成的sql进行保留两位小数
```tpl
SUM({{ $column.Condition }}) as `{{ $column.Field }}`修改为
ROUND(SUM({{ $column.Condition }}), 2) as `{{ $column.Field }}`
```
## 如何使用通用翻译API？
+ 使用您的百度账号登录[百度翻译开放平台](https://api.fanyi.baidu.com/)；
+ 注册成为开发者，获得 APPID ；
+ 进行开发者认证（如仅需标准版可跳过）；
+ 开通通用翻译API服务：[开通链接](https://fanyi-api.baidu.com/choose)；

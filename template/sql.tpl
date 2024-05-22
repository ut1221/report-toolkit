SELECT
{{- range $index, $column := .Columns }}
    {{- if $index }}, {{ end }}
    SUM({{ $column.Condition }}) as `{{ $column.Field }}`
{{- end }}
FROM {{ .TableName }};

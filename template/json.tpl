[
{{- range $index, $element := . }}
    {{- if $index }},{{ end }}
    {
        "title": "{{ $element.Title }}",
        "align": "{{ $element.Align }}"
        {{- if $element.DataIndex }}, "dataIndex": "{{ $element.DataIndex }}"{{ end }}
        {{- if $element.Key }}, "key": "{{ $element.Key }}"{{ end }}
        {{- if ne $element.Width 0 }}, "width": {{ $element.Width }}{{ end }}
        {{- if $element.Children | len }}, "children": [
            {{ template "children" $element.Children }}
        ]{{ end }}
    }
{{- end }}
]

{{ define "children" }}
    {{- range $index, $child := . }}
        {{- if $index }},{{ end }}
        {
            "title": "{{ $child.Title }}",
            "align": "{{ $child.Align }}"
            {{- if $child.DataIndex }}, "dataIndex": "{{ $child.DataIndex }}"{{ end }}
            {{- if $child.Key }}, "key": "{{ $child.Key }}"{{ end }}
            {{- if ne $child.Width 0 }}, "width": {{ $child.Width }}{{ end }}
            {{- if $child.Children | len }}, "children": [
                {{ template "children" $child.Children }}
            ]{{ end }}
        }
    {{- end }}
{{ end }}
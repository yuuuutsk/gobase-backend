{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}

func New{{$alias.UpSingular}}(
        {{- range $column := .Table.Columns -}}
        {{- $colAlias := $alias.Column $column.Name -}}
        {{- $orig_col_name := $column.Name -}}
        {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
        {{- else -}}
        {{$colAlias}} {{$column.Type}},
        {{- end -}}
        {{- end -}}
) *{{$alias.UpSingular}} {
    return &{{$alias.UpSingular}}{
{{- range $column := .Table.Columns -}}
        {{- $colAlias := $alias.Column $column.Name -}}
        {{- $orig_col_name := $column.Name -}}
        {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
        {{- else -}}
        {{$colAlias}}: {{$colAlias}},
        {{- end -}}
        {{- end -}}
    }
}
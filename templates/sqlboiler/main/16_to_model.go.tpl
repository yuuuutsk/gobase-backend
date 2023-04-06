{{- $alias := .Aliases.Table .Table.Name -}}
{{- $schemaTable := .Table.Name | .SchemaTable}}
{{- $orig_tbl_name := .Table.Name -}}

// dao to model
func (o *{{$alias.UpSingular}}) ToModel() *models.{{$alias.UpSingular}} {
    return models.Restore{{$alias.UpSingular}}(
        {{- range $column := .Table.Columns -}}
        {{- $colAlias := $alias.Column $column.Name -}}
        {{- $orig_col_name := $column.Name -}}
        {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
        {{- else -}}
        o.{{$colAlias}},
        {{- end -}}
        {{- end -}}
)
}

func NewFrom{{$alias.UpSingular}}(o *models.{{$alias.UpSingular}}) *{{$alias.UpSingular}} {
    return &{{$alias.UpSingular}}{
        {{- range $column := .Table.Columns -}}
        {{- $colAlias := $alias.Column $column.Name -}}
        {{- $orig_col_name := $column.Name -}}
        {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
        {{- else -}}
        {{$colAlias}}: o.{{$colAlias}},
        {{- end -}}
        {{- end -}}
    }
}
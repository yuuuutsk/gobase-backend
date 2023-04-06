{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}

// {{$alias.UpSingular}} is an object representing the database table.
type {{$alias.UpSingular}} struct {
	{{- range $column := .Table.Columns -}}
	{{- $colAlias := $alias.Column $column.Name -}}
	{{- $orig_col_name := $column.Name -}}
	{{- range $column.Comment | splitLines -}} // {{ . }}
	{{end -}}
	{{if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
	{{$colAlias}} {{$column.Type}}
	{{else if eq $.StructTagCasing "title" -}}
	{{$colAlias}} {{$column.Type}}
	{{else if eq $.StructTagCasing "camel" -}}
	{{$colAlias}} {{$column.Type}}
	{{else if eq $.StructTagCasing "alias" -}}
	{{$colAlias}} {{$column.Type}}
	{{else -}}
	{{$colAlias}} {{$column.Type}}
	{{end -}}
	{{end -}}
}
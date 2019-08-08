package template

var autoGenWarningTemplate = `
//
// Code generated by go-jet DO NOT EDIT.
// Generated at {{now}}
//
// WARNING: Changes to this file may cause incorrect behavior  
// and will be lost if the code is regenerated
//

`

var tableSQLBuilderTemplate = ` 
{{define "column-list" -}}
	{{- range $i, $c := . }}
		{{- if gt $i 0 }}, {{end}}{{ToGoIdentifier $c.Name}}Column
	{{- end}}
{{- end}}

package table

import (
	"github.com/go-jet/jet/{{dialect.PackageName}}"
)

var {{ToGoIdentifier .Name}} = new{{.GoStructName}}()

type {{.GoStructName}} struct {
	{{dialect.PackageName}}.Table
	
	//Columns
{{- range .Columns}}
	{{ToGoIdentifier .Name}} {{dialect.PackageName}}.Column{{.SqlBuilderColumnType}}
{{- end}}

	AllColumns     {{dialect.PackageName}}.IColumnList
	MutableColumns {{dialect.PackageName}}.IColumnList
}

// creates new {{.GoStructName}} with assigned alias
func (a *{{.GoStructName}}) AS(alias string) *{{.GoStructName}} {
	aliasTable := new{{.GoStructName}}()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func new{{.GoStructName}}() *{{.GoStructName}} {
	var (
	{{- range .Columns}}
		{{ToGoIdentifier .Name}}Column = {{dialect.PackageName}}.{{.SqlBuilderColumnType}}Column("{{.Name}}")
	{{- end}}
	)

	return &{{.GoStructName}}{
		Table: {{dialect.PackageName}}.NewTable("{{.SchemaName}}", "{{.Name}}", {{template "column-list" .Columns}}),

		//Columns
{{- range .Columns}}
		{{ToGoIdentifier .Name}}: {{ToGoIdentifier .Name}}Column,
{{- end}}

		AllColumns:     {{dialect.PackageName}}.ColumnList( {{template "column-list" .Columns}} ),
		MutableColumns: {{dialect.PackageName}}.ColumnList( {{template "column-list" .MutableColumns}} ),
	}
}

`

var tableModelTemplate = `package model

{{ if .GetImports }}
import (
{{- range .GetImports}}
	"{{.}}"
{{- end}}
)
{{end}}


type {{ToGoIdentifier .Name}} struct {
{{- range .Columns}}
	{{ToGoIdentifier .Name}} {{.GoModelType}} ` + "{{.GoModelTag ($.IsPrimaryKey .Name)}}" + `
{{- end}}
}


`
var enumSQLBuilderTemplate = `package enum

import "github.com/go-jet/jet/{{dialect.PackageName}}"

var {{ToGoIdentifier $.Name}} = &struct {
{{- range $index, $element := .Values}}
	{{ToGoIdentifier $element}} {{dialect.PackageName}}.StringExpression
{{- end}}
} {
{{- range $index, $element := .Values}}
	{{ToGoIdentifier $element}}: {{dialect.PackageName}}.NewEnumValue("{{$element}}"),
{{- end}}
}
`

var enumModelTemplate = `package model

import "errors"

type {{ToGoIdentifier $.Name}} string

const (
{{- range $index, $element := .Values}}
	{{ToGoIdentifier $.Name}}_{{ToGoIdentifier $element}} {{ToGoIdentifier $.Name}} = "{{$element}}"
{{- end}}
)

func (e *{{ToGoIdentifier $.Name}}) Scan(value interface{}) error {
	if v, ok := value.(string); !ok {
		return errors.New("jet: Invalid data for {{ToGoIdentifier $.Name}} enum")
	} else {
		switch string(v) {
{{- range $index, $element := .Values}}
		case "{{$element}}":
			*e = {{ToGoIdentifier $.Name}}_{{ToGoIdentifier $element}}
{{- end}}
		default:
			return errors.New("jet: Inavlid data " + string(v) + "for {{ToGoIdentifier $.Name}} enum")
		}

		return nil
	}
}

func (e {{ToGoIdentifier $.Name}}) String() string {
	return string(e)
}

`

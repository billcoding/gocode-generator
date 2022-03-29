package {{.Config.Model.PKG}}

// @author {{.Config.Global.Author}}
{{if .Config.Global.Date}}// @since {{.Extra.Date}}{{end}}
{{if .Config.Global.Copyright}}// @created by {{.Config.Global.CopyrightContent}}{{end}}
{{if .Config.Global.Website}}// @repo {{.Config.Global.WebsiteContent}}{{end}}

import (
	"encoding/json"
    "github.com/go-the-way/sg"{{if .Model.Orm}}
     _ "{{.Config.Module}}/{{.Config.Config.PKG}}"
    "github.com/go-the-way/anorm"{{end}}{{if .Model.ImportTime}}
    "time"{{end}}{{if .Model.ImportSql}}
    "database/sql"{{end}}
)
{{$mapperEnable := .Config.MapperEnable}}
{{if .Config.Model.Comment}}// {{.Model.Name}} struct {{.Model.Table.Comment}}{{end}}
type {{.Model.Name}} struct {
    {{range $i, $e := .Model.Ids}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} {{$e.Type}} `{{if $mapperEnable}}db:"{{$e.Column.Name}}" {{end}}{{if $e.Orm}}{{$e.Column.OrmTag}}{{end}}{{if $e.JSONTag}} json:"{{$e.JSONTagName}}"{{end}} generator:"DB_PRI"`
    {{end}}{{range $i, $e := .Model.Fields}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} {{$e.Type}} `{{if $mapperEnable}}db:"{{$e.Column.Name}}" {{end}}{{if $e.Orm}}{{$e.Column.OrmTag}}{{end}}{{if $e.JSONTag}} json:"{{$e.JSONTagName}}"{{end}}`
    {{end}}
}
{{if .Model.Orm}}
func (model *{{.Model.Name}}) MetaData() *anorm.ModelMeta {
	return &anorm.ModelMeta{
		Comment:          "{{.Model.Table.Comment}}",
		IndexDefinitions: []sg.Ge{ {{range $i, $e := .Model.OrmIndexDefinitions}}
			{{$e}},{{end}}
		},
	}
}

func init(){
	anorm.Register(new({{.Model.Name}}))
}
{{end}}

// FieldMap model to map named with fields
func (model *{{.Model.Name}}) FieldMap() map[string]interface{} {
	return map[string]interface{}{
	    {{range $i, $e := .Model.Ids}}
        "{{$e.Name}}": model.{{$e.Name}},
        {{end}}{{range $i, $e := .Model.Fields}}
        "{{$e.Name}}": model.{{$e.Name}},
        {{end}}
    }
}

// ColumnMap model to map named with columns
func (model *{{.Model.Name}}) ColumnMap() map[string]interface{} {
	return map[string]interface{}{
	    {{range $i, $e := .Model.Ids}}
        "{{$e.Column.Name}}": model.{{$e.Name}},
        {{end}}{{range $i, $e := .Model.Fields}}
        "{{$e.Column.Name}}": model.{{$e.Name}},
        {{end}}
    }
}

// JSON model to json
func (model *{{.Model.Name}}) JSON() string {
	bytes, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

var {{.Model.Name}}Columns = &struct{ {{range $i, $e := .Model.Ids}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} sg.C
    {{end}}{{range $i, $e := .Model.Fields}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} sg.C
    {{end}}
}{ {{range $i, $e := .Model.Ids}}
    {{$e.Name}} : "{{$e.Column.Name}}",
    {{end}}{{range $i, $e := .Model.Fields}}
    {{$e.Name}} : "{{$e.Column.Name}}",
{{end}}
}
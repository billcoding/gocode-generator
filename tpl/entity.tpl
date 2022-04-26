package {{.Config.Entity.PKG}}

// @author {{.Config.Global.Author}}
{{if .Config.Global.Date}}// @since {{.Extra.Date}}{{end}}
{{if .Config.Global.Copyright}}// @created by {{.Config.Global.CopyrightContent}}{{end}}
{{if .Config.Global.Website}}// @repo {{.Config.Global.WebsiteContent}}{{end}}

import (
	"encoding/json"
    "github.com/go-the-way/sg"{{if .Entity.Orm}}
     _ "{{.Config.Module}}/{{.Config.Config.PKG}}"
    "github.com/go-the-way/anorm"{{end}}{{if .Entity.ImportSql}}
    "database/sql"{{end}}
)
{{$mapperEnable := .Config.MapperEnable}}
{{if .Config.Entity.Comment}}// {{.Entity.Name}} struct {{.Entity.Table.Comment}}{{end}}
type {{.Entity.Name}} struct {
    {{range $i, $e := .Entity.Ids}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} {{$e.Type}} `{{if $mapperEnable}}db:"{{$e.Column.Name}}" {{end}}{{if $e.Orm}}{{$e.Column.OrmTag}}{{end}}{{if $e.JSONTag}} json:"{{$e.JSONTagName}}"{{end}} generator:"DB_PRI"`
    {{end}}{{range $i, $e := .Entity.Fields}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} {{$e.Type}} `{{if $mapperEnable}}db:"{{$e.Column.Name}}" {{end}}{{if $e.Orm}}{{$e.Column.OrmTag}}{{end}}{{if $e.JSONTag}} json:"{{$e.JSONTagName}}"{{end}}`
    {{end}}
}
{{if .Entity.Orm}}
func (model *{{.Entity.Name}}) Configure(c *anorm.EC) {
	c.Migrate = true
	c.Commented = true
	c.Comment = "{{.Entity.Table.Comment}}"
	c.IndexDefinitions = []sg.Ge{ {{range $i, $e := .Entity.OrmIndexDefinitions}}
		{{$e}},{{end}}
	}
}

func init(){
	anorm.Register(new({{.Entity.Name}}))
}
{{end}}

// FieldMap model to map named with fields
func (model *{{.Entity.Name}}) FieldMap() map[string]interface{} {
	return map[string]interface{}{
	    {{range $i, $e := .Entity.Ids}}
        "{{$e.Name}}": model.{{$e.Name}},
        {{end}}{{range $i, $e := .Entity.Fields}}
        "{{$e.Name}}": model.{{$e.Name}},
        {{end}}
    }
}

// ColumnMap model to map named with columns
func (model *{{.Entity.Name}}) ColumnMap() map[string]interface{} {
	return map[string]interface{}{
	    {{range $i, $e := .Entity.Ids}}
        "{{$e.Column.Name}}": model.{{$e.Name}},
        {{end}}{{range $i, $e := .Entity.Fields}}
        "{{$e.Column.Name}}": model.{{$e.Name}},
        {{end}}
    }
}

// JSON model to json
func (model *{{.Entity.Name}}) JSON() string {
	bytes, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

var {{.Entity.Name}}Columns = &struct{ {{range $i, $e := .Entity.Ids}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} sg.C
    {{end}}{{range $i, $e := .Entity.Fields}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} sg.C
    {{end}}
}{ {{range $i, $e := .Entity.Ids}}
    {{$e.Name}} : "{{$e.Column.Name}}",
    {{end}}{{range $i, $e := .Entity.Fields}}
    {{$e.Name}} : "{{$e.Column.Name}}",
{{end}}
}
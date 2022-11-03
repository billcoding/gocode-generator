package {{.Config.Entity.PKG}}

// @author {{.Config.Global.Author}}
{{if .Config.Global.Date}}// @since {{.Extra.Date}}{{end}}
{{if .Config.Global.Copyright}}// @created by {{.Config.Global.CopyrightContent}}{{end}}
{{if .Config.Global.Website}}// @repo {{.Config.Global.WebsiteContent}}{{end}}

import (
    "github.com/go-the-way/sg"
)

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

var {{.Entity.Name}}AliasColumns = &struct{ {{range $i, $e := .Entity.Ids}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} sg.C
    {{end}}{{range $i, $e := .Entity.Fields}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} sg.C
    {{end}}
}{ {{range $i, $e := .Entity.Ids}}
    {{$e.Name}} : "t.{{$e.Column.Name}}",
    {{end}}{{range $i, $e := .Entity.Fields}}
    {{$e.Name}} : "t.{{$e.Column.Name}}",
{{end}}
}
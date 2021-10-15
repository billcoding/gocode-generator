package {{.Config.Controller.PKG}}

// @author {{.Config.Global.Author}}
{{if .Config.Global.Date}}// @since {{.Extra.Date}}{{end}}
{{if .Config.Global.Copyright}}// @created by {{.Config.Global.CopyrightContent}}{{end}}
{{if .Config.Global.Website}}// @repo {{.Config.Global.WebsiteContent}}{{end}}

import (
    "github.com/billcoding/flygo"
    "github.com/billcoding/flygo/context"
)

{{if .Config.Controller.Comment}}// {{.Controller.Name}} {{.Controller.Model.Table.Comment}} Controller{{end}}
type {{.Controller.Name}} struct {
}

func init() {
    app := flygo.GetApp()
    {{.Controller.VarName}} := &{{.Controller.Name}}{}
    app.Controller({{.Controller.VarName}})
}

// Prefix returns route prefix
func (ctl *{{.Controller.Name}}) Prefix() string {
	return "{{.Controller.Route}}"
}

// Get return routed GET handler
func (ctl *{{.Controller.Name}}) Get() func(ctx *context.Context) {
	return func(ctx *context.Context) {
        ctx.Text("Get")
	}
}

// Gets return routed GETS handler
func (ctl *{{.Controller.Name}}) Gets() func(ctx *context.Context) {
	return func(ctx *context.Context) {
        ctx.Text("Gets")
	}
}

// Post return routed POST handler
func (ctl *{{.Controller.Name}}) Post() func(ctx *context.Context) {
	return func(ctx *context.Context) {
        ctx.Text("Post")
	}
}

// Put return routed PUT handler
func (ctl *{{.Controller.Name}}) Put() func(ctx *context.Context) {
	return func(ctx *context.Context) {
        ctx.Text("Put")
	}
}

// Delete return routed DELETE handler
func (ctl *{{.Controller.Name}}) Delete() func(ctx *context.Context) {
	return func(ctx *context.Context) {
        ctx.Text("Delete")
	}
}
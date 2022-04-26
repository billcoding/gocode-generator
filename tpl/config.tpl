package {{.Config.Config.PKG}}

// @author {{.Config.Global.Author}}
{{if .Config.Global.Date}}// @since {{.Extra.Date}}{{end}}
{{if .Config.Global.Copyright}}// @created by {{.Config.Global.CopyrightContent}}{{end}}
{{if .Config.Global.Website}}// @repo {{.Config.Global.WebsiteContent}}{{end}}

import ({{if .Config.Entity.Orm}}
	"database/sql"
	"github.com/go-the-way/anorm"
	{{end}}
	_ "github.com/go-sql-driver/mysql"
)

var (
	DSN = ""
)

func init() {
	{{if .Config.Entity.Orm}}db, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
	anorm.DS(db){{end}}
}

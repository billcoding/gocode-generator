package {{.Config.Config.PKG}}

// @author {{.Config.Global.Author}}
{{if .Config.Global.Date}}// @since {{.Extra.Date}}{{end}}
{{if .Config.Global.Copyright}}// @created by {{.Config.Global.CopyrightContent}}{{end}}
{{if .Config.Global.Website}}// @repo {{.Config.Global.WebsiteContent}}{{end}}

import ({{if .Config.Model.Orm}}
	"database/sql"
	"github.com/billcoding/gorm"
	{{end}}{{if .Config.MapperEnable}}ba "github.com/billcoding/gobatis"
	{{end}}_ "github.com/go-sql-driver/mysql"
)

var (
	{{if .Config.MapperEnable}}{{.Config.Mapper.Batis}} = ba.Default()
	{{end}}DSN = ""
)

func init() {
	{{if .Config.MapperEnable}}{{.Config.Mapper.Batis}}.DSN(DSN)
	{{end}}{{if .Config.Model.Orm}}db, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
	gorm.DS(db){{end}}
}

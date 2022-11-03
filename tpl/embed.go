package tpl

import "embed"

var (
	//go:embed config.tpl entity.tpl column.tpl
	FS               embed.FS
	configTpl        = `config.tpl`
	entityTpl        = `entity.tpl`
	columnTpl        = `column.tpl`
	configTplContent = ""
	entityTplContent = ""
	columnTplContent = ""
)

func ConfigTpl() string {
	if configTplContent == "" {
		file, err := FS.ReadFile(configTpl)
		if err != nil {
			panic(err)
		}
		configTplContent = string(file)
	}
	return configTplContent
}

func EntityTpl() string {
	if entityTplContent == "" {
		file, err := FS.ReadFile(entityTpl)
		if err != nil {
			panic(err)
		}
		entityTplContent = string(file)
	}
	return entityTplContent
}

func ColumnTpl() string {
	if columnTplContent == "" {
		file, err := FS.ReadFile(columnTpl)
		if err != nil {
			panic(err)
		}
		columnTplContent = string(file)
	}
	return columnTplContent
}

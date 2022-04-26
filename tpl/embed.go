package tpl

import "embed"

//go:embed entity.tpl config.tpl controller.tpl service.tpl
var FS embed.FS
var entityTpl = `entity.tpl`

var configTpl = `config.tpl`
var controllerTpl = `controller.tpl`
var serviceTpl = `service.tpl`
var entityTplContent = ""

var configTplContent = ""
var controllerTplContent = ""
var serviceTplContent = ""

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

func ControllerTpl() string {
	if controllerTplContent == "" {
		file, err := FS.ReadFile(controllerTpl)
		if err != nil {
			panic(err)
		}
		controllerTplContent = string(file)
	}
	return controllerTplContent
}

func ServiceTpl() string {
	if serviceTplContent == "" {
		file, err := FS.ReadFile(serviceTpl)
		if err != nil {
			panic(err)
		}
		serviceTplContent = string(file)
	}
	return serviceTplContent
}

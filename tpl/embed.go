package tpl

import "embed"

//go:embed model.tpl config.tpl controller.tpl service.tpl
var FS embed.FS
var modelTpl = `model.tpl`

var configTpl = `config.tpl`
var controllerTpl = `controller.tpl`
var serviceTpl = `service.tpl`
var modelTplContent = ""

var configTplContent = ""
var controllerTplContent = ""
var serviceTplContent = ""

func ModelTpl() string {
	if modelTplContent == "" {
		file, err := FS.ReadFile(modelTpl)
		if err != nil {
			panic(err)
		}
		modelTplContent = string(file)
	}
	return modelTplContent
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

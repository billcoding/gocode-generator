package bundle

import (
	. "github.com/billcoding/gocode-generator/config"
	. "github.com/billcoding/gocode-generator/generator"
	. "github.com/billcoding/gocode-generator/model"
)

func GetModelGenerators(CFG *Configuration, tableMap map[string]*Table) []Generator {
	egs := make([]Generator, 0)
	for _, v := range tableMap {
		eg := &ModelGenerator{
			C:     CFG,
			Table: v,
		}
		eg.Init()
		egs = append(egs, eg)
	}
	return egs
}

func GetCfgGenerator(CFG *Configuration) Generator {
	return &ConfigGenerator{
		C: CFG,
	}
}

func GetControllerGenerators(CFG *Configuration, modelGenerators []Generator) []Generator {
	cgs := make([]Generator, 0)
	for _, eg := range modelGenerators {
		cg := &ControllerGenerator{
			C: CFG,
		}
		cg.Init(eg.(*ModelGenerator).Model)
		cgs = append(cgs, cg)
	}
	return cgs
}

func GetServiceGenerators(CFG *Configuration, modelGenerators []Generator) []Generator {
	sgs := make([]Generator, 0)
	for _, eg := range modelGenerators {
		sg := &ServiceGenerator{
			C: CFG,
		}
		sg.Init(eg.(*ModelGenerator).Model)
		sgs = append(sgs, sg)
	}
	return sgs
}

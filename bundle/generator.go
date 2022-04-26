package bundle

import (
	. "github.com/billcoding/gocode-generator/config"
	. "github.com/billcoding/gocode-generator/generator"
	. "github.com/billcoding/gocode-generator/model"
)

func GetEntityGenerators(CFG *Configuration, tableMap map[string]*Table) []Generator {
	egs := make([]Generator, 0)
	for _, v := range tableMap {
		eg := &EntityGenerator{
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

func GetControllerGenerators(CFG *Configuration, entityGenerators []Generator) []Generator {
	cgs := make([]Generator, 0)
	for _, eg := range entityGenerators {
		cg := &ControllerGenerator{
			C: CFG,
		}
		cg.Init(eg.(*EntityGenerator).Entity)
		cgs = append(cgs, cg)
	}
	return cgs
}

func GetServiceGenerators(CFG *Configuration, entityGenerators []Generator) []Generator {
	sgs := make([]Generator, 0)
	for _, eg := range entityGenerators {
		sg := &ServiceGenerator{
			C: CFG,
		}
		sg.Init(eg.(*EntityGenerator).Entity)
		sgs = append(sgs, sg)
	}
	return sgs
}

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

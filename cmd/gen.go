package cmd

import (
	"fmt"
	. "github.com/billcoding/gocode-generator/bundle"
	. "github.com/billcoding/gocode-generator/config"
	. "github.com/billcoding/gocode-generator/generator"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var genCmd = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"g", "generate"},
	Short:   "Generate Go files",
	Long: `Generate Go files.
Simply type gocode-generator help gen for full details.`,
	Example: `gocode-generator gen -D "root:123@tcp(127.0.0.1:3306)/test" -d "database" -M "awesome"
gocode-generator gen -D "root:123@tcp(127.0.0.1:3306)/test" -d "database" -M "awesome" -o "/to/path" 
gocode-generator gen -D "root:123@tcp(127.0.0.1:3306)/test" -d "database" -M "awesome" -au "bigboss" -o "/to/path"`,
	Run: func(cmd *cobra.Command, args []string) {
		CFG.Verbose = verbose

		if dsn == "" {
			fmt.Fprintln(os.Stderr, "The DSN is required")
			return
		}

		if module == "" {
			fmt.Fprintln(os.Stderr, "The Module name is required")
			return
		}

		if database == "" {
			fmt.Fprintln(os.Stderr, "The Database name is required")
			return
		}

		if !entity {
			fmt.Fprintln(os.Stderr, "Nothing do...")
			return
		}

		Init(dsn)
		setCFG()

		tableList := Tables(database, CFG)
		columnList := Columns(database)
		indexList := Indexes(database)
		tableMap := TransformTables(tableList)
		columnMap := TransformColumns(columnList)
		SetTableColumns(tableMap, columnMap)
		indexMap := TransformIndexes(indexList)
		SetTableIndexes(tableMap, indexMap)
		generators := make([]Generator, 0)

		entityGenerators := GetEntityGenerators(CFG, tableMap)
		generators = append(generators, entityGenerators...)

		if config {
			generators = append(generators, GetCfgGenerator(CFG))
		}

		if controller {
			controllerGenerators := GetControllerGenerators(CFG, entityGenerators)
			generators = append(generators, controllerGenerators...)
		}

		if service {
			serviceGenerators := GetServiceGenerators(CFG, entityGenerators)
			generators = append(generators, serviceGenerators...)
		}

		for _, g := range generators {
			g.Generate()
		}
	},
}

func init() {
	genCmd.PersistentFlags().StringVarP(&outputDir, "output-dir", "o", "", "The output dir")
	genCmd.PersistentFlags().StringVarP(&dsn, "dsn", "D", "", "The MySQL DSN")
	genCmd.PersistentFlags().StringVarP(&database, "db", "d", "", "The Database name")
	genCmd.PersistentFlags().StringVarP(&module, "module", "M", "", "The Module name")
	genCmd.PersistentFlags().StringVarP(&includeTable, "include-table", "I", "", "The include table names[table_a,table_b]")
	genCmd.PersistentFlags().StringVarP(&excludeTable, "exclude-table", "E", "", "The exclude table names[table_a,table_b]")
	genCmd.PersistentFlags().StringVar(&author, "author", "bill", "The file copyright author")
	genCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "Print verbose output")

	genCmd.PersistentFlags().BoolVar(&entity, "entity", true, "Generate Entity Go file")
	genCmd.PersistentFlags().BoolVar(&orm, "orm", true, "Generate Entity Go anorm supports")
	genCmd.PersistentFlags().StringVarP(&entityPKG, "entity-pkg", "1", "entity", "The Entity package")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityDefault, "table2entity-default", false, "The Table to Entity name strategy: default")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityFirstLetterUpper, "table2entity-first-letter-upper", false, "The Table to Entity name strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityUnderlineToCamel, "table2entity-underline-to-camel", false, "The Table to Entity name strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityUnderlineToUpper, "table2entity-underline-to-upper", true, "The Table to Entity name strategy: UnderlineToUpper")
	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldDefault, "column2field-default", false, "The column to field name strategy: default")
	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldFirstLetterUpper, "column2field-first-letter-upper", false, "The column to field name strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldUnderlineToCamel, "column2field-underline-to-camel", false, "The column to field name strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldUnderlineToUpper, "column2field-underline-to-upper", true, "The column to field name strategy: UnderlineToUpper")
	genCmd.PersistentFlags().BoolVar(&entityFileNameDefault, "entity-filename-default", true, "The Entity file name strategy: default")
	genCmd.PersistentFlags().BoolVar(&entityFileNameFirstLetterUpper, "entity-filename-first-letter-upper", false, "The Entity file name strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&entityFileNameUnderlineToCamel, "entity-filename-underline-to-camel", false, "The Entity file name strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&entityFileNameUnderlineToUpper, "entity-filename-underline-to-upper", false, "The Entity file name strategy: UnderlineToUpper")
	genCmd.PersistentFlags().BoolVar(&entityComment, "entity-comment", true, "Generate Entity comment")
	genCmd.PersistentFlags().BoolVar(&entityFieldComment, "entity-field-comment", true, "Generate Entity field comment")
	genCmd.PersistentFlags().BoolVar(&entityJSONTag, "entity-json-tag", true, "Generate Entity field JSON tag")
	genCmd.PersistentFlags().BoolVar(&entityJSONTagKeyDefault, "entity-json-tag-key-default", true, "The Entity JSON Tag key strategy: default")
	genCmd.PersistentFlags().BoolVar(&entityJSONTagKeyFirstLetterUpper, "entity-json-tag-key-first-letter-upper", false, "The Entity JSON Tag key strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&entityJSONTagKeyUnderlineToCamel, "entity-json-tag-key-underline-to-camel", false, "The Entity JSON Tag key strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&entityJSONTagKeyUnderlineToUpper, "entity-json-tag-key-underline-to-upper", false, "The Entity JSON Tag key strategy: UnderlineToUpper")

	genCmd.PersistentFlags().BoolVarP(&config, "config", "C", true, "Generate Config")
	genCmd.PersistentFlags().StringVarP(&configPKG, "config-pkg", "3", "config", "The Config package")

	genCmd.PersistentFlags().BoolVarP(&controller, "controller", "c", false, "Generate Controller file")
	genCmd.PersistentFlags().StringVarP(&controllerPKG, "controller-pkg", "4", "controller", "The Controller package")
	genCmd.PersistentFlags().BoolVar(&controllerNameDefault, "controller-name-default", false, "The Controller name strategy: default")
	genCmd.PersistentFlags().BoolVar(&controllerNameFirstLetterUpper, "controller-name-first-letter-upper", false, "The Controller name strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&controllerNameUnderlineToCamel, "controller-name-underline-to-camel", false, "The Controller name strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&controllerNameUnderlineToUpper, "controller-name-underline-to-upper", true, "The Controller name strategy: UnderlineToUpper")

	genCmd.PersistentFlags().BoolVar(&controllerVarDefault, "controller-var-default", false, "The Controller var strategy: default")
	genCmd.PersistentFlags().BoolVar(&controllerVarFirstLetterUpper, "controller-var-first-letter-upper", false, "The Controller var strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&controllerVarUnderlineToCamel, "controller-var-underline-to-camel", true, "The Controller var strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&controllerVarUnderlineToUpper, "controller-var-underline-to-upper", false, "The Controller var strategy: UnderlineToUpper")

	genCmd.PersistentFlags().BoolVar(&controllerFileNameDefault, "controller-filename-default", true, "The Controller FileName strategy: default")
	genCmd.PersistentFlags().BoolVar(&controllerFileNameFirstLetterUpper, "controller-filename-first-letter-upper", false, "The Controller FileName strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&controllerFileNameUnderlineToCamel, "controller-filename-underline-to-camel", false, "The Controller FileName strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&controllerFileNameUnderlineToUpper, "controller-filename-underline-to-upper", false, "The Controller FileName strategy: UnderlineToUpper")

	genCmd.PersistentFlags().BoolVar(&controllerRouteDefault, "controller-route-default", true, "The Controller Route strategy: default")
	genCmd.PersistentFlags().BoolVar(&controllerRouteFirstLetterUpper, "controller-route-first-letter-upper", false, "The Controller Route strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&controllerRouteUnderlineToCamel, "controller-route-underline-to-camel", false, "The Controller Route strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&controllerRouteUnderlineToUpper, "controller-route-underline-to-upper", false, "The Controller Route strategy: UnderlineToUpper")

	genCmd.PersistentFlags().StringVar(&controllerNamePrefix, "controller-name-prefix", "", "The controller name prefix")
	genCmd.PersistentFlags().StringVar(&controllerNameSuffix, "controller-name-suffix", "Controller", "The controller name suffix")
	genCmd.PersistentFlags().StringVar(&controllerRoutePrefix, "controller-route-prefix", "/", "The controller route prefix")
	genCmd.PersistentFlags().StringVar(&controllerRouteSuffix, "controller-route-suffix", "", "The controller route suffix")
	genCmd.PersistentFlags().StringVar(&controllerVarPrefix, "controller-var-prefix", "", "The controller var prefix")
	genCmd.PersistentFlags().StringVar(&controllerVarSuffix, "controller-var-suffix", "Controller", "The controller var suffix")
	genCmd.PersistentFlags().BoolVar(&controllerComment, "controller-comment", true, "Generate Controller comment")

	genCmd.PersistentFlags().BoolVarP(&service, "service", "s", false, "Generate Service file")
	genCmd.PersistentFlags().StringVarP(&servicePKG, "service-pkg", "5", "service", "The Service package")
	genCmd.PersistentFlags().BoolVar(&serviceNameDefault, "service-name-default", false, "The Service name strategy: default")
	genCmd.PersistentFlags().BoolVar(&serviceNameFirstLetterUpper, "service-name-first-letter-upper", false, "The Service name strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&serviceNameUnderlineToCamel, "service-name-underline-to-camel", true, "The Service name strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&serviceNameUnderlineToUpper, "service-name-underline-to-upper", false, "The Service name strategy: UnderlineToUpper")

	genCmd.PersistentFlags().BoolVar(&serviceVarDefault, "service-var-default", true, "The Service var strategy: default")
	genCmd.PersistentFlags().BoolVar(&serviceVarFirstLetterUpper, "service-var-first-letter-upper", false, "The Service var strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&serviceVarUnderlineToCamel, "service-var-underline-to-camel", false, "The Service var strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&serviceVarUnderlineToUpper, "service-var-underline-to-upper", true, "The Service var strategy: UnderlineToUpper")

	genCmd.PersistentFlags().BoolVar(&serviceFileNameDefault, "service-filename-default", true, "The Service FileName strategy: default")
	genCmd.PersistentFlags().BoolVar(&serviceFileNameFirstLetterUpper, "service-filename-first-letter-upper", false, "The Service FileName strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&serviceFileNameUnderlineToCamel, "service-filename-underline-to-camel", false, "The Service FileName strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&serviceFileNameUnderlineToUpper, "service-filename-underline-to-upper", false, "The Service FileName strategy: UnderlineToUpper")

	genCmd.PersistentFlags().StringVar(&serviceNamePrefix, "service-name-prefix", "", "The Service name prefix")
	genCmd.PersistentFlags().StringVar(&serviceNameSuffix, "service-name-suffix", "Service", "The Service name suffix")
	genCmd.PersistentFlags().StringVar(&serviceVarPrefix, "service-var-prefix", "", "The Service var prefix")
	genCmd.PersistentFlags().StringVar(&serviceVarSuffix, "service-var-suffix", "Service", "The Service var suffix")
	genCmd.PersistentFlags().BoolVar(&serviceComment, "service-comment", true, "Generate Service comment")

	rootCmd.AddCommand(genCmd)
}

var (
	outputDir    = ""
	dsn          = ""
	module       = ""
	database     = ""
	includeTable = ""
	excludeTable = ""
	author       = ""
	verbose      = false

	orm                                = false
	entity                             = true
	entityPKG                          = "entity"
	entityTable2EntityDefault          = false
	entityTable2EntityFirstLetterUpper = false
	entityTable2EntityUnderlineToCamel = false
	entityTable2EntityUnderlineToUpper = true

	entityColumn2FieldDefault          = false
	entityColumn2FieldFirstLetterUpper = false
	entityColumn2FieldUnderlineToCamel = false
	entityColumn2FieldUnderlineToUpper = true

	entityFileNameDefault          = true
	entityFileNameFirstLetterUpper = false
	entityFileNameUnderlineToCamel = false
	entityFileNameUnderlineToUpper = false

	entityComment                    = true
	entityFieldComment               = true
	entityJSONTag                    = true
	entityJSONTagKeyDefault          = true
	entityJSONTagKeyFirstLetterUpper = false
	entityJSONTagKeyUnderlineToCamel = false
	entityJSONTagKeyUnderlineToUpper = false

	config    = true
	configPKG = "config"

	controller    = true
	controllerPKG = "controller"

	controllerNameDefault          = false
	controllerNameFirstLetterUpper = false
	controllerNameUnderlineToCamel = true
	controllerNameUnderlineToUpper = false

	controllerVarDefault          = false
	controllerVarFirstLetterUpper = false
	controllerVarUnderlineToCamel = false
	controllerVarUnderlineToUpper = true

	controllerRouteDefault          = true
	controllerRouteFirstLetterUpper = false
	controllerRouteUnderlineToCamel = false
	controllerRouteUnderlineToUpper = false

	controllerFileNameDefault          = true
	controllerFileNameFirstLetterUpper = false
	controllerFileNameUnderlineToCamel = false
	controllerFileNameUnderlineToUpper = false

	controllerNamePrefix  = ""
	controllerNameSuffix  = "Controller"
	controllerVarPrefix   = ""
	controllerVarSuffix   = "Controller"
	controllerRoutePrefix = "/"
	controllerRouteSuffix = ""
	controllerComment     = true

	service    = true
	servicePKG = "service"

	serviceNameDefault          = false
	serviceNameFirstLetterUpper = false
	serviceNameUnderlineToCamel = true
	serviceNameUnderlineToUpper = false

	serviceVarDefault          = false
	serviceVarFirstLetterUpper = false
	serviceVarUnderlineToCamel = false
	serviceVarUnderlineToUpper = true

	serviceFileNameDefault          = true
	serviceFileNameFirstLetterUpper = false
	serviceFileNameUnderlineToCamel = false
	serviceFileNameUnderlineToUpper = false

	serviceNamePrefix = ""
	serviceNameSuffix = "Service"
	serviceVarPrefix  = ""
	serviceVarSuffix  = "Service"
	serviceComment    = true
)

var CFG = &Configuration{
	Module:        "",
	OutputDir:     "",
	Verbose:       false,
	IncludeTables: make([]string, 0),
	ExcludeTables: make([]string, 0),
	Global: &GlobalConfiguration{
		Author:           "bill",
		Date:             true,
		DateLayout:       "2006-01-02",
		Copyright:        true,
		CopyrightContent: "A Go source code generator written in Golang",
		Website:          true,
		WebsiteContent:   "https://github.com/billcoding/gocode-generator",
	},
	Entity: &EntityConfiguration{
		PKG:                   "entity",
		TableToEntityStrategy: UnderlineToUpper,
		ColumnToFieldStrategy: UnderlineToUpper,
		FileNameStrategy:      Default,
		JSONTag:               true,
		JSONTagKeyStrategy:    Default,
		FieldIdUpper:          true,
		Comment:               true,
		FieldComment:          true,
		NamePrefix:            "",
		NameSuffix:            "",
		Orm:                   true,
	},
	Config: &CfgConfiguration{
		PKG:  "config",
		Name: "config",
	},
	Controller: &ControllerConfiguration{
		PKG:              "controller",
		NameStrategy:     UnderlineToCamel,
		VarNameStrategy:  UnderlineToUpper,
		RouteStrategy:    Default,
		FileNameStrategy: Default,
		NamePrefix:       "",
		NameSuffix:       "Controller",
		RoutePrefix:      "/",
		RouteSuffix:      "",
		VarNamePrefix:    "",
		VarNameSuffix:    "Controller",
		Comment:          true,
	},
	Service: &ServiceConfiguration{
		PKG:              "service",
		NameStrategy:     UnderlineToCamel,
		VarNameStrategy:  UnderlineToUpper,
		FileNameStrategy: Default,
		NamePrefix:       "",
		NameSuffix:       "Service",
		VarNamePrefix:    "",
		VarNameSuffix:    "Service",
		Comment:          true,
	},
}

func setCFG() {
	if outputDir != "" {
		CFG.OutputDir = outputDir
	}
	if CFG.OutputDir == "" {
		exec, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		CFG.OutputDir = exec
	}
	if module != "" {
		CFG.Module = module
	}
	if includeTable != "" {
		CFG.IncludeTables = strings.Split(includeTable, ",")
	} else if excludeTable != "" {
		CFG.ExcludeTables = strings.Split(excludeTable, ",")
	}

	if author != "" {
		CFG.Global.Author = author
	}

	{
		if entityPKG != "" {
			CFG.Entity.PKG = entityPKG
		}

		CFG.Entity.Orm = orm
		CFG.Entity.Comment = entityComment
		CFG.Entity.FieldComment = entityFieldComment
		CFG.Entity.JSONTag = entityJSONTag

		switch {
		case entityTable2EntityUnderlineToUpper:
			CFG.Entity.TableToEntityStrategy = UnderlineToUpper
		case entityTable2EntityUnderlineToCamel:
			CFG.Entity.TableToEntityStrategy = UnderlineToCamel
		case entityTable2EntityFirstLetterUpper:
			CFG.Entity.TableToEntityStrategy = FirstLetterUpper
		case entityTable2EntityDefault:
			CFG.Entity.TableToEntityStrategy = Default
		}

		switch {
		case entityColumn2FieldUnderlineToUpper:
			CFG.Entity.ColumnToFieldStrategy = UnderlineToUpper
		case entityColumn2FieldUnderlineToCamel:
			CFG.Entity.ColumnToFieldStrategy = UnderlineToCamel
		case entityColumn2FieldFirstLetterUpper:
			CFG.Entity.ColumnToFieldStrategy = FirstLetterUpper
		case entityColumn2FieldDefault:
			CFG.Entity.ColumnToFieldStrategy = Default
		}

		switch {
		case entityFileNameUnderlineToUpper:
			CFG.Entity.FileNameStrategy = UnderlineToUpper
		case entityFileNameUnderlineToCamel:
			CFG.Entity.FileNameStrategy = UnderlineToCamel
		case entityFileNameFirstLetterUpper:
			CFG.Entity.FileNameStrategy = FirstLetterUpper
		case entityFileNameDefault:
			CFG.Entity.FileNameStrategy = Default
		}

		switch {
		case entityJSONTagKeyUnderlineToUpper:
			CFG.Entity.JSONTagKeyStrategy = UnderlineToUpper
		case entityJSONTagKeyUnderlineToCamel:
			CFG.Entity.JSONTagKeyStrategy = UnderlineToCamel
		case entityJSONTagKeyFirstLetterUpper:
			CFG.Entity.JSONTagKeyStrategy = FirstLetterUpper
		case entityJSONTagKeyDefault:
			CFG.Entity.JSONTagKeyStrategy = Default
		}
	}

	{
		if configPKG != "" {
			CFG.Config.PKG = configPKG
		}
	}

	{
		if controllerPKG != "" {
			CFG.Controller.PKG = controllerPKG
		}
		CFG.Controller.NamePrefix = controllerNamePrefix
		CFG.Controller.NameSuffix = controllerNameSuffix
		CFG.Controller.RoutePrefix = controllerRoutePrefix
		CFG.Controller.RouteSuffix = controllerRouteSuffix
		CFG.Controller.VarNamePrefix = controllerVarPrefix
		CFG.Controller.VarNameSuffix = controllerVarSuffix
		CFG.Controller.Comment = controllerComment

		switch {
		case controllerNameUnderlineToUpper:
			CFG.Controller.NameStrategy = UnderlineToUpper
		case controllerNameUnderlineToCamel:
			CFG.Controller.NameStrategy = UnderlineToCamel
		case controllerNameFirstLetterUpper:
			CFG.Controller.NameStrategy = FirstLetterUpper
		case controllerNameDefault:
			CFG.Controller.NameStrategy = Default
		}

		switch {
		case controllerVarUnderlineToUpper:
			CFG.Controller.VarNameStrategy = UnderlineToUpper
		case controllerVarUnderlineToCamel:
			CFG.Controller.VarNameStrategy = UnderlineToCamel
		case controllerVarFirstLetterUpper:
			CFG.Controller.VarNameStrategy = FirstLetterUpper
		case controllerVarDefault:
			CFG.Controller.VarNameStrategy = Default
		}

		switch {
		case controllerRouteUnderlineToUpper:
			CFG.Controller.RouteStrategy = UnderlineToUpper
		case controllerRouteUnderlineToCamel:
			CFG.Controller.RouteStrategy = UnderlineToCamel
		case controllerRouteFirstLetterUpper:
			CFG.Controller.RouteStrategy = FirstLetterUpper
		case controllerRouteDefault:
			CFG.Controller.RouteStrategy = Default
		}

		switch {
		case controllerFileNameUnderlineToUpper:
			CFG.Controller.FileNameStrategy = UnderlineToUpper
		case controllerFileNameUnderlineToCamel:
			CFG.Controller.FileNameStrategy = UnderlineToCamel
		case controllerFileNameFirstLetterUpper:
			CFG.Controller.FileNameStrategy = FirstLetterUpper
		case controllerFileNameDefault:
			CFG.Controller.FileNameStrategy = Default
		}
	}

	{
		if servicePKG != "" {
			CFG.Service.PKG = servicePKG
		}
		CFG.Service.NamePrefix = serviceNamePrefix
		CFG.Service.NameSuffix = serviceNameSuffix
		CFG.Service.VarNamePrefix = serviceVarPrefix
		CFG.Service.VarNameSuffix = serviceVarSuffix
		CFG.Service.Comment = serviceComment

		switch {
		case serviceNameUnderlineToUpper:
			CFG.Service.NameStrategy = UnderlineToUpper
		case serviceNameUnderlineToCamel:
			CFG.Service.NameStrategy = UnderlineToCamel
		case serviceNameFirstLetterUpper:
			CFG.Service.NameStrategy = FirstLetterUpper
		case serviceNameDefault:
			CFG.Service.NameStrategy = Default
		}

		switch {
		case serviceVarUnderlineToUpper:
			CFG.Service.VarNameStrategy = UnderlineToUpper
		case serviceVarUnderlineToCamel:
			CFG.Service.VarNameStrategy = UnderlineToCamel
		case serviceVarFirstLetterUpper:
			CFG.Service.VarNameStrategy = FirstLetterUpper
		case serviceVarDefault:
			CFG.Service.VarNameStrategy = Default
		}
		switch {
		case serviceFileNameUnderlineToUpper:
			CFG.Service.FileNameStrategy = UnderlineToUpper
		case serviceFileNameUnderlineToCamel:
			CFG.Service.FileNameStrategy = UnderlineToCamel
		case serviceFileNameFirstLetterUpper:
			CFG.Service.FileNameStrategy = FirstLetterUpper
		case serviceFileNameDefault:
			CFG.Service.FileNameStrategy = Default
		}
	}
}

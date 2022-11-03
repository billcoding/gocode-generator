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
			_, _ = fmt.Fprintln(os.Stderr, "The DSN is required")
			return
		}

		if module == "" {
			_, _ = fmt.Fprintln(os.Stderr, "The Module name is required")
			return
		}

		if database == "" {
			_, _ = fmt.Fprintln(os.Stderr, "The Database name is required")
			return
		}

		if !entity {
			_, _ = fmt.Fprintln(os.Stderr, "Nothing do...")
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
	genCmd.PersistentFlags().BoolVar(&onlyColumnAlias, "only-column-alias", false, "Only generate entity's column alias")
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
	onlyColumnAlias                    = false
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
		OnlyColumnAlias:       false,
	},
	Config: &CfgConfiguration{
		PKG:  "config",
		Name: "config",
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
		CFG.Entity.OnlyColumnAlias = onlyColumnAlias
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

}

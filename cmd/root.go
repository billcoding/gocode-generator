package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocode-generator",
	Short: "A Go source code generator written in Golang",
	Long: `A Go source code generator for 
[gobatis](https://github.com/billcoding/gobatis)
[gorm](https://github.com/billcoding/gorm) 
[flygo](https://github.com/billcoding/flygo)
written in Golang.
It supports to generate:
1. global config
2. gobatis Model, Mapper and itself XML
3. gorm Model, Field's Column tag definition and metadata
4. flygo REST-ful Controller and Service`,
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			PrintVersion(false)
		}
	},
}

var versionFlag bool

// Execute executes the root command.
func Execute() error {
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "version")
	return rootCmd.Execute()
}

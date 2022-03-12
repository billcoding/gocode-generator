package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocode-generator",
	Short: "A Go source code generator written in Golang",
	Long: `A Go source code generator for 
[gobatis](https://github.com/billcoding/gobatis)
[anorm](https://github.com/go-the-way/anorm) 
[anoweb](https://github.com/go-the-way/anoweb)
written in Golang.`,
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

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocode-generator",
	Short: "A Go source code generator written in Golang",
	Long:  "A Go source code generator written in Golang",
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

// gocode-generator gen -V -C=F -D "root:wIiouqLR8v4vEFIxlHKw8ir9URc@tcp(222.186.173.3:3307)/zy_proj" -d "zy_proj" -M "system" --only-column-alias=T

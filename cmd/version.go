package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "ver"},
	Short:   "Print the version",
	Long: `Print the version of gocode-generator.
Simply type fly help version for full details.`,
	Example: `gocode-generator version [-a]`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintVersion(goVersionFlag)
	},
}

var goVersionFlag = false

func init() {
	versionCmd.PersistentFlags().BoolVarP(&goVersionFlag, "all", "a", false, "print Go SDK version")
	rootCmd.AddCommand(versionCmd)
}

const version = "1.1.0"

func PrintVersion(goVersion bool) {
	_, _ = fmt.Fprintln(os.Stdout, version)
	if goVersion {
		_, _ = fmt.Fprintln(os.Stdout, runtime.Version())
	}
}

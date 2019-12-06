package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

const VERSION = "3.0.2"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print secrets version",
	Long: `
Print secrets version and exit.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", VERSION)
	},
}

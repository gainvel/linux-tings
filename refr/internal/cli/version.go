package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of refr",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("refr v%s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

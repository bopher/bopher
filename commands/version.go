package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VersionCommand get cli version
var VersionCommand = &cobra.Command{
	Use:   "version",
	Short: "get bopher cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 1.1.0")
	},
}

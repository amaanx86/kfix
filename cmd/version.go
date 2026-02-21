package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\nVersion: %s\nSource: %s\nDocs: %s\n", banner, version, sourceURL, docsURL)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

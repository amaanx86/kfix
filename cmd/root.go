package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

const (
	sourceURL = "https://github.com/amaanx86/kfix"
	docsURL   = "https://kfix.readthedocs.io"
)

const banner = `
 _   __  _____   _____   __   __
| | / / |  ___| |_   _|  \ \ / /
| |/ /  | |_      | |     \ V / 
|    \  |  _|     | |     /   \ 
| |\  \ | |      _| |_   / / \ \
\_| \_/ \_|     |_____|  \/   \/
`

var rootCmd = &cobra.Command{
	Use:   "kfix",
	Short: "Kubernetes YAML formatter for clean, consistent manifests",
	Long: fmt.Sprintf(`%s
Version: %s
Source: %s
Docs: %s

kfix is an opinionated Kubernetes YAML formatter that understands 
K8s resource structure and applies consistent formatting rules.

It formats YAML files while maintaining proper field ordering for 
Kubernetes resources and applying context-aware indentation.`, banner, version, sourceURL, docsURL),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntP("indent", "i", 2, "number of spaces for indentation")
	rootCmd.PersistentFlags().BoolP("in-place", "w", false, "write result to file instead of stdout")
}

package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/amaanx86/kfix/pkg/formatter"
	"github.com/amaanx86/kfix/pkg/k8s"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format [file...]",
	Short: "Format Kubernetes YAML files",
	Long: `Format one or more Kubernetes YAML files with consistent styling.
If no files are specified, reads from stdin.`,
	Args: cobra.ArbitraryArgs,
	RunE: runFormat,
}

func init() {
	rootCmd.AddCommand(formatCmd)
}

func runFormat(cmd *cobra.Command, args []string) error {
	indent, _ := cmd.Flags().GetInt("indent")
	inPlace, _ := cmd.Flags().GetBool("in-place")

	opts := formatter.Options{
		Indent: indent,
	}
	f := formatter.New(opts)

	if len(args) == 0 {
		return formatStdin(f)
	}

	for _, file := range args {
		if err := formatFile(f, file, inPlace); err != nil {
			return fmt.Errorf("failed to format %s: %w", file, err)
		}
	}

	return nil
}

func formatStdin(f *formatter.Formatter) error {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %w", err)
	}

	if !k8s.IsK8sResource(input) {
		return fmt.Errorf("input does not appear to be valid Kubernetes resource(s)")
	}

	output, err := f.Format(input)
	if err != nil {
		return err
	}

	fmt.Print(string(output))
	return nil
}

func formatFile(f *formatter.Formatter, path string, inPlace bool) error {
	input, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if !k8s.IsK8sResource(input) {
		return fmt.Errorf("file does not appear to be valid Kubernetes resource(s)")
	}

	output, err := f.Format(input)
	if err != nil {
		return err
	}

	if inPlace {
		if err := os.WriteFile(path, output, 0644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	} else {
		fmt.Print(string(output))
	}

	return nil
}

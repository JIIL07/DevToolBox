package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devtoolbox",
	Short: "DevToolBox - Code generation tool from JSON schemas",
	Long: `DevToolBox is a powerful code generation tool that converts JSON schemas
into various programming language structures like Go structs, TypeScript interfaces, and more.

Features:
- Generate Go structs with JSON tags
- Generate TypeScript interfaces
- Plugin system for custom generators
- Web interface for easy usage
- CLI for advanced users`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(pluginCmd)
	rootCmd.AddCommand(serverCmd)
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

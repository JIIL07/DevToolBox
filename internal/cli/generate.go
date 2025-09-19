package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JIIL07/devtoolbox/internal/core"
	"github.com/JIIL07/devtoolbox/internal/plugins"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate [template] [input-file]",
	Short: "Generate code from JSON schema",
	Long: `Generate code from a JSON schema file using the specified template.

Examples:
  devtoolbox generate go-struct schema.json
  devtoolbox generate ts_interface_gen schema.json
  devtoolbox generate go-struct -i '{"name": "string", "age": "number"}'`,
	Args: cobra.RangeArgs(1, 2),
	Run:  runGenerate,
}

var inputInline string

func init() {
	generateCmd.Flags().StringVarP(&inputInline, "input", "i", "", "JSON input as string")
}

func runGenerate(cmd *cobra.Command, args []string) {
	template := args[0]
	
	var input string
	var err error
	
	if inputInline != "" {
		input = inputInline
	} else if len(args) > 1 {
		inputFile := args[1]
		if !filepath.IsAbs(inputFile) {
			wd, _ := os.Getwd()
			inputFile = filepath.Join(wd, inputFile)
		}
		
		data, err := os.ReadFile(inputFile)
		if err != nil {
			exitWithError(fmt.Errorf("failed to read input file: %v", err))
		}
		input = string(data)
	} else {
		exitWithError(fmt.Errorf("either input file or --input flag is required"))
	}
	
	registry := core.NewGeneratorRegistry()
	
	pluginManager := plugins.NewPluginManager()
	customPlugins, err := pluginManager.LoadPlugins()
	if err == nil {
		for _, pluginInfo := range customPlugins {
			if pluginInfo.Type == "python" {
				pythonPlugin := plugins.NewPythonPlugin(pluginInfo.Name, pluginInfo.Description, pluginInfo.Path)
				registry.Register(pythonPlugin)
			}
		}
	}
	
	generator, exists := registry.Get(template)
	if !exists {
		available := registry.GetNames()
		exitWithError(fmt.Errorf("template '%s' not found. Available templates: %v", template, available))
	}
	
	result, err := generator.Generate(input)
	if err != nil {
		exitWithError(fmt.Errorf("generation failed: %v", err))
	}
	
	fmt.Println(result)
}

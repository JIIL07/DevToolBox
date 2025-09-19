package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JIIL07/devtoolbox/internal/plugins"
	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Manage plugins",
	Long:  "Add, remove, and list custom plugins for code generation",
}

var pluginAddCmd = &cobra.Command{
	Use:   "add <plugin-path>",
	Short: "Add a new plugin",
	Long: `Add a new plugin to the system.

Examples:
  devtoolbox plugin add ./plugins/custom/my_plugin.py
  devtoolbox plugin add /path/to/plugin.py`,
	Args: cobra.ExactArgs(1),
	Run:  runPluginAdd,
}

var pluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all plugins",
	Long:  "List all available plugins (official and custom)",
	Run:   runPluginList,
}

var pluginRemoveCmd = &cobra.Command{
	Use:   "remove <plugin-name>",
	Short: "Remove a plugin",
	Long: `Remove a plugin from the system.

Examples:
  devtoolbox plugin remove my_plugin
  devtoolbox plugin remove ts_interface_gen`,
	Args: cobra.ExactArgs(1),
	Run:  runPluginRemove,
}

func init() {
	pluginCmd.AddCommand(pluginAddCmd)
	pluginCmd.AddCommand(pluginListCmd)
	pluginCmd.AddCommand(pluginRemoveCmd)
}

func runPluginAdd(cmd *cobra.Command, args []string) {
	pluginPath := args[0]
	
	if !filepath.IsAbs(pluginPath) {
		wd, _ := os.Getwd()
		pluginPath = filepath.Join(wd, pluginPath)
	}
	
	if _, err := os.Stat(pluginPath); os.IsNotExist(err) {
		exitWithError(fmt.Errorf("plugin file not found: %s", pluginPath))
	}
	
	manager := plugins.NewPluginManager()
	err := manager.AddPlugin(pluginPath)
	if err != nil {
		exitWithError(fmt.Errorf("failed to add plugin: %v", err))
	}
	
	fmt.Printf("Plugin added successfully: %s\n", filepath.Base(pluginPath))
}

func runPluginList(cmd *cobra.Command, args []string) {
	manager := plugins.NewPluginManager()
	pluginList, err := manager.ListPlugins()
	if err != nil {
		exitWithError(fmt.Errorf("failed to list plugins: %v", err))
	}
	
	if len(pluginList) == 0 {
		fmt.Println("No plugins found.")
		return
	}
	
	fmt.Println("Available plugins:")
	fmt.Println("==================")
	
	for _, plugin := range pluginList {
		fmt.Printf("Name: %s\n", plugin.Name)
		fmt.Printf("Description: %s\n", plugin.Description)
		fmt.Printf("Type: %s\n", plugin.Type)
		fmt.Printf("Path: %s\n", plugin.Path)
		fmt.Println("---")
	}
}

func runPluginRemove(cmd *cobra.Command, args []string) {
	pluginName := args[0]
	
	manager := plugins.NewPluginManager()
	err := manager.RemovePlugin(pluginName)
	if err != nil {
		exitWithError(fmt.Errorf("failed to remove plugin: %v", err))
	}
	
	fmt.Printf("Plugin removed successfully: %s\n", pluginName)
}

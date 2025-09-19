package plugins

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type PluginInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Path        string `json:"path"`
}

type PluginManager struct {
	configPath string
}

func NewPluginManager() *PluginManager {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".devtoolbox", "plugins.json")
	
	return &PluginManager{
		configPath: configPath,
	}
}

func (pm *PluginManager) AddPlugin(pluginPath string) error {
	plugins, err := pm.LoadPlugins()
	if err != nil {
		return err
	}
	
	pluginName := filepath.Base(pluginPath)
	pluginName = pluginName[:len(pluginName)-len(filepath.Ext(pluginName))]
	
	for _, plugin := range plugins {
		if plugin.Name == pluginName {
			return fmt.Errorf("plugin '%s' already exists", pluginName)
		}
	}
	
	newPlugin := PluginInfo{
		Name:        pluginName,
		Description: fmt.Sprintf("Custom plugin: %s", pluginName),
		Type:        "python",
		Path:        pluginPath,
	}
	
	plugins = append(plugins, newPlugin)
	
	return pm.savePlugins(plugins)
}

func (pm *PluginManager) RemovePlugin(pluginName string) error {
	plugins, err := pm.LoadPlugins()
	if err != nil {
		return err
	}
	
	found := false
	for i, plugin := range plugins {
		if plugin.Name == pluginName {
			plugins = append(plugins[:i], plugins[i+1:]...)
			found = true
			break
		}
	}
	
	if !found {
		return fmt.Errorf("plugin '%s' not found", pluginName)
	}
	
	return pm.savePlugins(plugins)
}

func (pm *PluginManager) ListPlugins() ([]PluginInfo, error) {
	plugins, err := pm.LoadPlugins()
	if err != nil {
		return nil, err
	}
	
	officialPlugins := []PluginInfo{
		{
			Name:        "go-struct",
			Description: "Генерирует Go структуры с JSON тегами из JSON схемы",
			Type:        "go",
			Path:        "builtin",
		},
		{
			Name:        "ts_interface_gen",
			Description: "Python plugin: ts_interface_gen",
			Type:        "python",
			Path:        "plugins/official/ts_interface_gen.py",
		},
	}
	
	allPlugins := append(officialPlugins, plugins...)
	return allPlugins, nil
}

func (pm *PluginManager) LoadPlugins() ([]PluginInfo, error) {
	if _, err := os.Stat(pm.configPath); os.IsNotExist(err) {
		return []PluginInfo{}, nil
	}
	
	data, err := os.ReadFile(pm.configPath)
	if err != nil {
		return nil, err
	}
	
	var plugins []PluginInfo
	err = json.Unmarshal(data, &plugins)
	return plugins, err
}

func (pm *PluginManager) savePlugins(plugins []PluginInfo) error {
	configDir := filepath.Dir(pm.configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}
	
	data, err := json.MarshalIndent(plugins, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(pm.configPath, data, 0644)
}

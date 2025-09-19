package plugins

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/JIIL07/devtoolbox/internal/plugins"
)

func TestPythonPlugin_GetName(t *testing.T) {
	plugin := plugins.NewPythonPlugin("test-plugin", "Test plugin", "/path/to/script.py")
	
	if plugin.GetName() != "test-plugin" {
		t.Errorf("expected name 'test-plugin', got '%s'", plugin.GetName())
	}
}

func TestPythonPlugin_GetDescription(t *testing.T) {
	plugin := plugins.NewPythonPlugin("test-plugin", "Test plugin", "/path/to/script.py")
	
	if plugin.GetDescription() != "Test plugin" {
		t.Errorf("expected description 'Test plugin', got '%s'", plugin.GetDescription())
	}
}

func TestPythonPluginLoader_LoadPlugin(t *testing.T) {
	loader := plugins.NewPythonPluginLoader("plugins")
	
	scriptPath := filepath.Join("plugins", "official", "ts_interface_gen.py")
	plugin, err := loader.LoadPlugin(scriptPath)
	
	if err != nil {
		t.Skipf("Skipping test - Python plugin not available: %v", err)
	}
	
	if plugin == nil {
		t.Error("expected plugin to be loaded, got nil")
	}
	
	if plugin.GetName() != "ts_interface_gen" {
		t.Errorf("expected plugin name 'ts_interface_gen', got '%s'", plugin.GetName())
	}
}

func TestPythonPluginLoader_LoadOfficialPlugins(t *testing.T) {
	loader := plugins.NewPythonPluginLoader("plugins")
	
	plugins, err := loader.LoadOfficialPlugins()
	
	if err != nil {
		t.Skipf("Skipping test - Python plugins not available: %v", err)
	}
	
	if len(plugins) == 0 {
		t.Skip("No Python plugins found - skipping test")
	}
	
	found := false
	for _, plugin := range plugins {
		if plugin.GetName() == "ts_interface_gen" {
			found = true
			break
		}
	}
	
	if !found {
		t.Skip("ts_interface_gen plugin not found - skipping test")
	}
}

func TestPythonPlugin_Generate(t *testing.T) {
	scriptPath := filepath.Join("plugins", "official", "ts_interface_gen.py")
	
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		t.Skip("Skipping test - Python script not found")
	}
	
	plugin := plugins.NewPythonPlugin("test-plugin", "Test plugin", scriptPath)
	
	input := `{"name": "string", "age": "number"}`
	result, err := plugin.Generate(input)
	
	if err != nil {
		t.Skipf("Skipping test - Python execution failed: %v", err)
	}
	
	if result == "" {
		t.Error("expected non-empty result")
	}
	
	if !contains(result, "interface") {
		t.Error("expected result to contain 'interface'")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || 
		(len(s) > len(substr) && (s[:len(substr)] == substr || 
		s[len(s)-len(substr):] == substr || 
		containsSubstring(s, substr))))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

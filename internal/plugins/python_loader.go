package plugins

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type PythonPlugin struct {
	name        string
	description string
	scriptPath  string
}

func NewPythonPlugin(name, description, scriptPath string) *PythonPlugin {
	return &PythonPlugin{
		name:        name,
		description: description,
		scriptPath:  scriptPath,
	}
}

func (p *PythonPlugin) GetName() string {
	return p.name
}

func (p *PythonPlugin) GetDescription() string {
	return p.description
}

func (p *PythonPlugin) Generate(input string) (string, error) {
	var cmd *exec.Cmd
	
	if _, err := exec.LookPath("/opt/venv/bin/python"); err == nil {
		cmd = exec.Command("/opt/venv/bin/python", p.scriptPath)
	} else if _, err := exec.LookPath("python3"); err == nil {
		cmd = exec.Command("python3", p.scriptPath)
	} else if _, err := exec.LookPath("python"); err == nil {
		cmd = exec.Command("python", p.scriptPath)
	} else {
		return "", fmt.Errorf("python executable not found in PATH")
	}
	
	cmd.Stdin = strings.NewReader(input)
	
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("python plugin execution failed: %w, stderr: %s", err, stderr.String())
	}
	
	result := strings.TrimSpace(stdout.String())
	if strings.HasPrefix(result, "Error:") {
		return "", fmt.Errorf("python plugin error: %s", result)
	}
	
	return result, nil
}

type PythonPluginLoader struct {
	pluginsDir string
}

func NewPythonPluginLoader(pluginsDir string) *PythonPluginLoader {
	return &PythonPluginLoader{
		pluginsDir: pluginsDir,
	}
}

func (l *PythonPluginLoader) LoadPlugin(scriptPath string) (*PythonPlugin, error) {
	absPath, err := filepath.Abs(scriptPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}
	
	pluginName := strings.TrimSuffix(filepath.Base(scriptPath), ".py")
	description := fmt.Sprintf("Python plugin: %s", pluginName)
	
	return NewPythonPlugin(pluginName, description, absPath), nil
}

func (l *PythonPluginLoader) LoadOfficialPlugins() ([]*PythonPlugin, error) {
	officialDir := filepath.Join(l.pluginsDir, "official")
	pattern := filepath.Join(officialDir, "*.py")
	
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to find python plugins: %w", err)
	}
	
	var plugins []*PythonPlugin
	for _, match := range matches {
		plugin, err := l.LoadPlugin(match)
		if err != nil {
			continue
		}
		plugins = append(plugins, plugin)
	}
	
	return plugins, nil
}

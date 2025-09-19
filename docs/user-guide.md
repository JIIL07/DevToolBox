# User Guide

Complete guide to using DevToolBox for code generation.

## 🎯 Overview

DevToolBox is a powerful code generation tool that helps developers create boilerplate code from JSON schemas. It supports multiple output formats and can be extended with custom plugins.

## 🚀 Getting Started

### Installation

Choose your preferred installation method:

```bash
# Docker (Recommended)
git clone https://github.com/JIIL07/devtoolbox.git
cd devtoolbox
make docker-up

# From Source
make install-deps
make build
```

### First Steps

1. **Start the application**:
   ```bash
   make docker-up
   ```

2. **Access the web interface**: http://localhost:3000

3. **Try your first generation**:
   - Paste JSON schema in the input field
   - Select a generator
   - Click "Generate Code"

## 🌐 Web Interface

### Main Features

#### Code Generation Form

- **JSON Input**: Paste your schema in the textarea
- **Generator Selection**: Choose from available generators
- **Generate Button**: Create code from your schema
- **Live Preview**: See results instantly

#### Code Preview

- **Syntax Highlighting**: Colored code for better readability
- **Copy to Clipboard**: One-click copying
- **Download**: Save generated code to file
- **Format**: Automatic code formatting

#### Example JSON

Click "Load Example" to see sample schemas:

```json
{
  "name": "string",
  "age": "int",
  "email": "string",
  "active": "bool"
}
```

### Navigation

- **Header**: Project logo and navigation
- **Main Area**: Generation form and preview
- **Footer**: Links and information

## 💻 Command Line Interface

### Basic Usage

```bash
# Show help
devtoolbox --help

# Generate code
devtoolbox generate --template go-struct --input '{"name": "string"}'

# List available generators
devtoolbox generate --help
```

### Generate Command

```bash
devtoolbox generate [options]
```

**Options:**
- `--template, -t`: Generator name (required)
- `--input, -i`: JSON input string
- `--input-file, -f`: Path to JSON file
- `--output, -o`: Output file path

**Examples:**
```bash
# From string
devtoolbox generate -t go-struct -i '{"user": "string", "age": "int"}'

# From file
devtoolbox generate -t ts-interface -f schema.json

# Save to file
devtoolbox generate -t go-struct -i '{"data": "string"}' -o model.go
```

### Plugin Management

```bash
# List plugins
devtoolbox plugin list

# Add plugin
devtoolbox plugin add ./my_plugin.py

# Remove plugin
devtoolbox plugin remove my_plugin
```

### Server Command

```bash
# Start web server
devtoolbox server

# Custom port
devtoolbox server --port 3000

# Custom host
devtoolbox server --host 0.0.0.0
```

## 🔌 Plugin System

### Built-in Generators

#### Go Struct Generator

Generates Go structs with JSON tags:

**Input:**
```json
{
  "name": "string",
  "age": "int",
  "email": "string"
}
```

**Output:**
```go
type GeneratedStruct struct {
    Age   string `json:"age"`
    Email string `json:"email"`
    Name  string `json:"name"`
}
```

#### TypeScript Interface Generator

Generates TypeScript interfaces:

**Input:**
```json
{
  "user": {
    "name": "string",
    "age": "int"
  },
  "active": "bool"
}
```

**Output:**
```typescript
interface User {
  age: number;
  name: string;
}

interface GeneratedInterface {
  active: boolean;
  user: User;
}
```

### Custom Plugins

#### Creating Plugins

1. **Create plugin file**:
   ```python
   # plugins/custom/my_plugin.py
   import json
   
   def generate(input_json: str) -> str:
       data = json.loads(input_json)
       # Your generation logic
       return generated_code
   ```

2. **Register plugin**:
   ```bash
   devtoolbox plugin add ./plugins/custom/my_plugin.py
   ```

3. **Use plugin**:
   ```bash
   devtoolbox generate --template my_plugin --input '{"data": "value"}'
   ```

#### Plugin Examples

**Java Class Generator:**
```python
def generate(input_json: str) -> str:
    data = json.loads(input_json)
    
    class_name = data.get('className', 'GeneratedClass')
    fields = data.get('fields', {})
    
    result = f"public class {class_name} {{\n"
    
    for field_name, field_type in fields.items():
        java_type = {
            'string': 'String',
            'int': 'int',
            'bool': 'boolean'
        }.get(field_type, 'Object')
        
        result += f"    private {java_type} {field_name};\n"
    
    result += "}\n"
    return result
```

**SQL Schema Generator:**
```python
def generate(input_json: str) -> str:
    data = json.loads(input_json)
    
    table_name = data.get('tableName', 'generated_table')
    columns = data.get('columns', {})
    
    result = f"CREATE TABLE {table_name} (\n"
    
    column_definitions = []
    for col_name, col_type in columns.items():
        sql_type = {
            'string': 'VARCHAR(255)',
            'int': 'INTEGER',
            'bool': 'BOOLEAN'
        }.get(col_type, 'TEXT')
        
        column_definitions.append(f"    {col_name} {sql_type}")
    
    result += ",\n".join(column_definitions)
    result += "\n);"
    
    return result
```

## 📊 JSON Schema Format

### Basic Structure

DevToolBox accepts JSON objects with type annotations:

```json
{
  "field_name": "type_annotation"
}
```

### Supported Types

- **string**: Text data
- **int**: Integer numbers
- **float**: Floating-point numbers
- **bool**: Boolean values
- **array**: Arrays (specify as `array<type>`)

### Complex Structures

#### Nested Objects

```json
{
  "user": {
    "name": "string",
    "profile": {
      "avatar": "string",
      "bio": "string"
    }
  },
  "settings": {
    "theme": "string",
    "notifications": "bool"
  }
}
```

#### Arrays

```json
{
  "users": "array<object>",
  "tags": "array<string>",
  "scores": "array<int>"
}
```

#### Mixed Types

```json
{
  "id": "int",
  "name": "string",
  "active": "bool",
  "metadata": {
    "created_at": "string",
    "updated_at": "string",
    "version": "int"
  },
  "permissions": "array<string>"
}
```

## 🎨 Customization

### Frontend Customization

#### Styling

Edit `frontend/src/index.css` for global styles:

```css
/* Custom theme */
:root {
  --primary-color: #your-color;
  --secondary-color: #your-color;
}

/* Custom component styles */
.generator-form {
  /* Your styles */
}
```

#### Configuration

Set environment variables in `frontend/.env`:

```env
VITE_API_URL=http://localhost:8080
VITE_APP_NAME=DevToolBox
```

### Backend Configuration

#### Environment Variables

```bash
export PORT=8080
export GIN_MODE=release
export LOG_LEVEL=info
export PLUGIN_DIR=./plugins
```

#### Custom Middleware

Add custom middleware in `internal/api/middleware.go`:

```go
func CustomMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Your middleware logic
        c.Next()
    }
}
```

## 🔧 Advanced Usage

### Batch Processing

#### Process Multiple Files

```bash
#!/bin/bash

# Generate Go models from multiple schemas
for schema in schemas/*.json; do
    name=$(basename "$schema" .json)
    devtoolbox generate --template go-struct \
        --input-file "$schema" \
        --output "models/${name}.go"
done
```

#### API Integration

```bash
# Generate from API response
curl -s "https://api.example.com/schema" | \
    devtoolbox generate --template go-struct \
        --input-file /dev/stdin \
        --output model.go
```

### Automation

#### CI/CD Integration

```yaml
# .github/workflows/generate-models.yml
name: Generate Models

on:
  push:
    paths: ['schemas/**']

jobs:
  generate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Generate models
        run: |
          for schema in schemas/*.json; do
            name=$(basename "$schema" .json)
            devtoolbox generate --template go-struct \
                --input-file "$schema" \
                --output "models/${name}.go"
          done
```

#### Pre-commit Hooks

```bash
#!/bin/bash
# .git/hooks/pre-commit

# Generate models before commit
for schema in schemas/*.json; do
    name=$(basename "$schema" .json)
    devtoolbox generate --template go-struct \
        --input-file "$schema" \
        --output "models/${name}.go"
done

git add models/
```

## 🚨 Troubleshooting

### Common Issues

#### Plugin Not Found

```bash
# Check plugin registration
devtoolbox plugin list

# Verify file exists
ls -la plugins/custom/my_plugin.py

# Re-register plugin
devtoolbox plugin remove my_plugin
devtoolbox plugin add ./plugins/custom/my_plugin.py
```

#### Invalid JSON

```bash
# Validate JSON
echo '{"test": "data"}' | python -m json.tool

# Test with simple input
devtoolbox generate --template go-struct --input '{"name": "string"}'
```

#### Generation Errors

```bash
# Test plugin directly
echo '{"test": "data"}' | python plugins/custom/my_plugin.py

# Check Python version
python3 --version

# Install dependencies
pip install -r requirements.txt
```

#### Web Interface Issues

```bash
# Check backend
curl http://localhost:8080/health

# Check frontend
curl http://localhost:3000

# Restart services
make docker-down
make docker-up
```

### Performance Issues

#### Large Schemas

For large JSON schemas:

1. **Split schemas**: Break into smaller parts
2. **Use files**: Use `--input-file` instead of `--input`
3. **Optimize plugins**: Improve plugin performance

#### Memory Usage

```bash
# Monitor memory usage
docker stats

# Restart if needed
make docker-restart
```

## 📚 Best Practices

### Schema Design

1. **Use descriptive names**: `user_name` instead of `n`
2. **Consistent types**: Use the same type for similar data
3. **Logical grouping**: Group related fields together
4. **Documentation**: Add comments for complex schemas

### Plugin Development

1. **Error handling**: Always handle invalid input
2. **Validation**: Validate input data
3. **Testing**: Write comprehensive tests
4. **Documentation**: Document plugin behavior

### Code Organization

1. **Modular plugins**: Create focused, single-purpose plugins
2. **Reusable components**: Share common functionality
3. **Version control**: Track plugin versions
4. **Backup**: Keep backups of important schemas

## 🎯 Use Cases

### API Development

Generate models for REST APIs:

```json
{
  "user": {
    "id": "int",
    "username": "string",
    "email": "string",
    "created_at": "string"
  },
  "pagination": {
    "page": "int",
    "limit": "int",
    "total": "int"
  }
}
```

### Database Design

Generate database schemas:

```json
{
  "tableName": "users",
  "columns": {
    "id": "int",
    "username": "string",
    "email": "string",
    "password_hash": "string",
    "created_at": "string",
    "updated_at": "string"
  }
}
```

### Configuration Management

Generate configuration classes:

```json
{
  "database": {
    "host": "string",
    "port": "int",
    "name": "string",
    "ssl": "bool"
  },
  "redis": {
    "host": "string",
    "port": "int",
    "password": "string"
  },
  "features": {
    "debug": "bool",
    "analytics": "bool",
    "notifications": "bool"
  }
}
```

## 📞 Support

### Getting Help

- 📖 [Documentation](docs/)
- 🐛 [Report Issues](https://github.com/JIIL07/devtoolbox/issues)
- 💬 [Discussions](https://github.com/JIIL07/devtoolbox/discussions)
- 📧 Email: romana9059@gmail.com

### Community

- ⭐ [Star the project](https://github.com/JIIL07/devtoolbox)
- 🍴 [Fork and contribute](https://github.com/JIIL07/devtoolbox/fork)
- 💬 [Join discussions](https://github.com/JIIL07/devtoolbox/discussions)
- 📢 [Share your plugins](https://github.com/JIIL07/devtoolbox/discussions/categories/plugins)

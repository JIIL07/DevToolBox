# Plugins Guide

DevToolBox supports a plugin architecture that allows you to extend functionality with custom code generators.

## Plugin Types

### Official Plugins

Built-in plugins that come with DevToolBox:

- **go-struct**: Generates Go structs with JSON tags
- **ts-interface**: Generates TypeScript interfaces

### Custom Plugins

User-created plugins for specific needs.

## Creating Python Plugins

### Plugin Structure

A Python plugin must have a `generate` function:

```python
def generate(input_json: str) -> str:
    """
    Generate code from JSON input.
    
    Args:
        input_json: JSON string containing the schema
        
    Returns:
        Generated code as string
    """
    # Your generation logic here
    return generated_code
```

### Example Plugin

Create `plugins/custom/my_plugin.py`:

```python
import json

def generate(input_json: str) -> str:
    try:
        data = json.loads(input_json)
    except json.JSONDecodeError:
        return "Error: Invalid JSON input"
    
    result = "// Generated code\n"
    result += "class MyClass {\n"
    
    for key, value in data.items():
        if isinstance(value, str):
            result += f"    {key}: string;\n"
        elif isinstance(value, int):
            result += f"    {key}: number;\n"
        elif isinstance(value, bool):
            result += f"    {key}: boolean;\n"
    
    result += "}\n"
    return result
```

### Plugin Naming Conventions

- Use descriptive names: `user_model_gen.py`
- Include language: `java_class_gen.py`
- Add purpose: `database_schema_gen.py`

## Managing Plugins

### CLI Commands

```bash
# List all plugins
devtoolbox plugin list

# Add a plugin
devtoolbox plugin add ./plugins/custom/my_plugin.py

# Remove a plugin
devtoolbox plugin remove my_plugin

# Test a plugin
echo '{"name": "John"}' | python plugins/custom/my_plugin.py
```

### Plugin Configuration

Plugins are stored in `~/.devtoolbox/plugins.json`:

```json
{
  "plugins": [
    {
      "name": "my_plugin",
      "path": "./plugins/custom/my_plugin.py",
      "description": "My custom plugin"
    }
  ]
}
```

## Plugin Development Best Practices

### 1. Error Handling

Always handle errors gracefully:

```python
def generate(input_json: str) -> str:
    try:
        data = json.loads(input_json)
    except json.JSONDecodeError as e:
        return f"Error: Invalid JSON - {str(e)}"
    
    try:
        # Your generation logic
        return generated_code
    except Exception as e:
        return f"Error: Generation failed - {str(e)}"
```

### 2. Input Validation

Validate input data:

```python
def generate(input_json: str) -> str:
    if not input_json.strip():
        return "Error: Empty input"
    
    try:
        data = json.loads(input_json)
    except json.JSONDecodeError:
        return "Error: Invalid JSON format"
    
    if not isinstance(data, dict):
        return "Error: Expected JSON object"
    
    # Continue with generation
```

### 3. Consistent Output

Maintain consistent code style:

```python
def generate(input_json: str) -> str:
    # Use consistent indentation
    # Add proper comments
    # Follow language conventions
    pass
```

## Advanced Plugin Features

### Using External Libraries

You can use external Python libraries in plugins:

```python
import json
import re
from typing import Dict, Any

def to_camel_case(snake_str: str) -> str:
    """Convert snake_case to camelCase"""
    components = snake_str.split('_')
    return components[0] + ''.join(x.title() for x in components[1:])

def generate(input_json: str) -> str:
    data = json.loads(input_json)
    
    result = "interface GeneratedInterface {\n"
    for key, value in data.items():
        camel_key = to_camel_case(key)
        result += f"  {camel_key}: {get_type(value)};\n"
    result += "}\n"
    
    return result

def get_type(value: Any) -> str:
    if isinstance(value, str):
        return "string"
    elif isinstance(value, int):
        return "number"
    elif isinstance(value, bool):
        return "boolean"
    else:
        return "any"
```

### Complex Data Structures

Handle nested objects and arrays:

```python
def generate(input_json: str) -> str:
    data = json.loads(input_json)
    
    def generate_interface(name: str, obj: dict, level: int = 0) -> str:
        indent = "  " * level
        result = f"{indent}interface {name} {{\n"
        
        for key, value in obj.items():
            field_type = get_field_type(key, value, level + 1)
            result += f"{indent}  {key}: {field_type};\n"
        
        result += f"{indent}}}\n"
        return result
    
    return generate_interface("Root", data)
```

## Testing Plugins

### Unit Tests

Create tests for your plugins:

```python
# tests/python/test_my_plugin.py
import pytest
from plugins.custom.my_plugin import generate

def test_simple_object():
    input_json = '{"name": "John", "age": 30}'
    result = generate(input_json)
    
    assert "class MyClass" in result
    assert "name: string" in result
    assert "age: number" in result

def test_invalid_json():
    result = generate("invalid json")
    assert "Error:" in result

def test_empty_input():
    result = generate("")
    assert "Error:" in result
```

### Integration Tests

Test plugins with the CLI:

```bash
# Test plugin directly
echo '{"test": "data"}' | python plugins/custom/my_plugin.py

# Test through CLI
devtoolbox generate --template my_plugin --input '{"test": "data"}'
```

## Plugin Examples

### Java Class Generator

```python
def generate(input_json: str) -> str:
    data = json.loads(input_json)
    
    class_name = data.get('className', 'GeneratedClass')
    fields = data.get('fields', {})
    
    result = f"public class {class_name} {{\n"
    
    for field_name, field_type in fields.items():
        java_type = map_to_java_type(field_type)
        result += f"    private {java_type} {field_name};\n"
    
    result += "}\n"
    return result

def map_to_java_type(type_str: str) -> str:
    mapping = {
        'string': 'String',
        'int': 'int',
        'float': 'float',
        'bool': 'boolean'
    }
    return mapping.get(type_str, 'Object')
```

### SQL Schema Generator

```python
def generate(input_json: str) -> str:
    data = json.loads(input_json)
    
    table_name = data.get('tableName', 'generated_table')
    columns = data.get('columns', {})
    
    result = f"CREATE TABLE {table_name} (\n"
    
    column_definitions = []
    for col_name, col_type in columns.items():
        sql_type = map_to_sql_type(col_type)
        column_definitions.append(f"    {col_name} {sql_type}")
    
    result += ",\n".join(column_definitions)
    result += "\n);"
    
    return result

def map_to_sql_type(type_str: str) -> str:
    mapping = {
        'string': 'VARCHAR(255)',
        'int': 'INTEGER',
        'float': 'FLOAT',
        'bool': 'BOOLEAN'
    }
    return mapping.get(type_str, 'TEXT')
```

## Contributing Plugins

### Submitting to Official Repository

1. Fork the repository
2. Create your plugin in `plugins/community/`
3. Add tests in `tests/python/`
4. Update documentation
5. Submit a pull request

### Plugin Guidelines

- Follow the plugin structure
- Include comprehensive tests
- Add documentation
- Handle errors gracefully
- Use consistent naming conventions

## Troubleshooting

### Common Issues

#### Plugin Not Found

```bash
# Check plugin path
devtoolbox plugin list

# Verify file exists
ls -la plugins/custom/my_plugin.py
```

#### Python Execution Errors

```bash
# Test Python directly
python3 --version

# Test plugin manually
echo '{"test": "data"}' | python3 plugins/custom/my_plugin.py
```

#### Import Errors

Ensure all dependencies are available:

```python
# Add error handling for imports
try:
    import external_library
except ImportError:
    print("Error: external_library not found")
    exit(1)
```

## Resources

- 📖 [Plugin API Reference](api-reference.md#plugins)
- 🧪 [Testing Guide](testing-guide.md)
- 💡 [Plugin Examples](examples/)
- 🐛 [Report Issues](https://github.com/JIIL07/devtoolbox/issues)

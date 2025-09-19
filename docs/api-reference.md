# API Reference

DevToolBox provides both REST API and CLI interfaces for code generation.

## REST API

### Base URL

- **Development**: `http://localhost:8080`
- **Production**: `https://api.devtoolbox.com`

### Authentication

Currently, the API is open and doesn't require authentication. Future versions may include API keys.

### Content Type

All API requests and responses use `application/json`.

## Endpoints

### Health Check

Check if the API is running.

```http
GET /health
```

**Response:**
```json
{
  "service": "devtoolbox-api",
  "status": "ok"
}
```

### List Generators

Get a list of available code generators.

```http
GET /generators
```

**Response:**
```json
{
  "generators": [
    {
      "name": "go-struct",
      "description": "Генерирует Go структуры с JSON тегами из JSON схемы"
    },
    {
      "name": "ts-interface",
      "description": "Генерирует TypeScript интерфейсы из JSON схемы"
    }
  ]
}
```

### Generate Code

Generate code using a specific template.

```http
POST /generate
```

**Request Body:**
```json
{
  "template": "go-struct",
  "input": "{\"name\": \"string\", \"age\": \"int\"}"
}
```

**Response:**
```json
{
  "code": "type GeneratedStruct struct {\n\tAge string `json:\"age\"`\n\tName string `json:\"name\"`\n}"
}
```

**Error Response:**
```json
{
  "error": "Invalid template: unknown-template"
}
```

## CLI Reference

### Global Options

```bash
devtoolbox [command] [options]
```

**Options:**
- `--help, -h`: Show help information
- `--version, -v`: Show version information

### Commands

#### Generate

Generate code from JSON input.

```bash
devtoolbox generate [options]
```

**Options:**
- `--template, -t`: Template name (required)
- `--input, -i`: JSON input string
- `--input-file, -f`: Path to JSON input file
- `--output, -o`: Output file path (optional)

**Examples:**
```bash
# Generate from string
devtoolbox generate --template go-struct --input '{"name": "string", "age": "int"}'

# Generate from file
devtoolbox generate --template ts-interface --input-file schema.json

# Save to file
devtoolbox generate --template go-struct --input '{"user": "string"}' --output user.go
```

#### Plugin Management

Manage custom plugins.

```bash
devtoolbox plugin [command]
```

**Commands:**
- `add <path>`: Add a new plugin
- `list`: List all plugins
- `remove <name>`: Remove a plugin

**Examples:**
```bash
# Add plugin
devtoolbox plugin add ./plugins/custom/my_plugin.py

# List plugins
devtoolbox plugin list

# Remove plugin
devtoolbox plugin remove my_plugin
```

#### Server

Start the web server.

```bash
devtoolbox server [options]
```

**Options:**
- `--port, -p`: Port number (default: 8080)
- `--host`: Host address (default: localhost)

**Examples:**
```bash
# Start server on default port
devtoolbox server

# Start on custom port
devtoolbox server --port 3000

# Start on all interfaces
devtoolbox server --host 0.0.0.0
```

## Data Types

### JSON Schema Format

DevToolBox accepts JSON objects with type annotations:

```json
{
  "field_name": "type_annotation"
}
```

**Supported Types:**
- `string`: Text data
- `int`: Integer numbers
- `float`: Floating-point numbers
- `bool`: Boolean values
- `array`: Arrays (specify as `array<type>`)

**Examples:**
```json
{
  "name": "string",
  "age": "int",
  "salary": "float",
  "active": "bool",
  "tags": "array<string>"
}
```

### Generated Code Examples

#### Go Struct

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

#### TypeScript Interface

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

## Error Handling

### HTTP Status Codes

- `200 OK`: Request successful
- `400 Bad Request`: Invalid request data
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

### Error Response Format

```json
{
  "error": "Error description",
  "details": "Additional error information"
}
```

### Common Errors

#### Invalid Template

```json
{
  "error": "Invalid template: unknown-template"
}
```

#### Invalid JSON

```json
{
  "error": "Invalid JSON input"
}
```

#### Generation Failed

```json
{
  "error": "Generation failed: plugin execution error"
}
```

## Rate Limiting

Currently, there are no rate limits. Future versions may implement rate limiting for production use.

## CORS

The API supports Cross-Origin Resource Sharing (CORS) for web applications:

- **Allowed Origins**: `*` (all origins)
- **Allowed Methods**: `GET`, `POST`, `OPTIONS`
- **Allowed Headers**: `Content-Type`, `Authorization`

## WebSocket Support

WebSocket support is planned for real-time code generation and collaboration features.

## SDKs and Libraries

### JavaScript/TypeScript

```typescript
import { DevToolBoxClient } from 'devtoolbox-client';

const client = new DevToolBoxClient('http://localhost:8080');

// Generate code
const result = await client.generate('go-struct', {
  name: 'string',
  age: 'int'
});

console.log(result.code);
```

### Python

```python
import requests

class DevToolBoxClient:
    def __init__(self, base_url='http://localhost:8080'):
        self.base_url = base_url
    
    def generate(self, template, data):
        response = requests.post(
            f'{self.base_url}/generate',
            json={
                'template': template,
                'input': json.dumps(data)
            }
        )
        return response.json()

# Usage
client = DevToolBoxClient()
result = client.generate('go-struct', {'name': 'string', 'age': 'int'})
print(result['code'])
```

### Go

```go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type DevToolBoxClient struct {
    BaseURL string
}

func (c *DevToolBoxClient) Generate(template string, data map[string]string) (string, error) {
    input, _ := json.Marshal(data)
    
    reqBody := map[string]string{
        "template": template,
        "input":    string(input),
    }
    
    jsonData, _ := json.Marshal(reqBody)
    
    resp, err := http.Post(
        c.BaseURL+"/generate",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    
    var result map[string]string
    json.NewDecoder(resp.Body).Decode(&result)
    
    return result["code"], nil
}
```

## Examples

### Complete Workflow

1. **List available generators:**
```bash
curl http://localhost:8080/generators
```

2. **Generate Go struct:**
```bash
curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d '{
    "template": "go-struct",
    "input": "{\"name\": \"string\", \"age\": \"int\"}"
  }'
```

3. **Save to file:**
```bash
devtoolbox generate --template go-struct \
  --input '{"name": "string", "age": "int"}' \
  --output user.go
```

### Batch Processing

```bash
#!/bin/bash

# Process multiple schemas
for file in schemas/*.json; do
  name=$(basename "$file" .json)
  devtoolbox generate --template go-struct \
    --input-file "$file" \
    --output "generated/${name}.go"
done
```

## Changelog

### v0.1.0
- Initial API release
- Go struct generation
- TypeScript interface generation
- Plugin system
- CLI interface

## Support

- 📖 [Documentation](docs/)
- 🐛 [Report Issues](https://github.com/JIIL07/devtoolbox/issues)
- 💬 [Discussions](https://github.com/JIIL07/devtoolbox/discussions)

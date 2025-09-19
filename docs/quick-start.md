# Quick Start Guide

Get up and running with DevToolBox in 5 minutes!

## 🚀 Quick Setup

### Option 1: Docker (Fastest)

```bash
# Clone and start
git clone https://github.com/JIIL07/devtoolbox.git
cd devtoolbox
make docker-up

# Access the application
open http://localhost:3000
```

### Option 2: Local Development

```bash
# Clone repository
git clone https://github.com/JIIL07/devtoolbox.git
cd devtoolbox

# Install dependencies
make install-deps

# Start services
make run
```

## 🎯 Your First Code Generation

### 1. Using the Web Interface

1. Open http://localhost:3000
2. Paste this JSON in the input field:
```json
{
  "name": "string",
  "age": "int",
  "email": "string",
  "active": "bool"
}
```
3. Select "Go Struct" from the dropdown
4. Click "Generate Code"
5. See your generated Go struct!

### 2. Using the CLI

```bash
# Generate Go struct
devtoolbox generate --template go-struct \
  --input '{"name": "string", "age": "int", "email": "string"}'

# Generate TypeScript interface
devtoolbox generate --template ts-interface \
  --input '{"user": {"name": "string", "age": "int"}, "active": "bool"}'
```

### 3. Using the API

```bash
# Generate via API
curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d '{
    "template": "go-struct",
    "input": "{\"name\": \"string\", \"age\": \"int\"}"
  }'
```

## 📝 Common Use Cases

### Generate Go Models

```bash
# User model
devtoolbox generate --template go-struct \
  --input '{
    "id": "int",
    "username": "string",
    "email": "string",
    "created_at": "string",
    "is_active": "bool"
  }' \
  --output models/user.go
```

### Generate TypeScript Interfaces

```bash
# API response interface
devtoolbox generate --template ts-interface \
  --input '{
    "data": {
      "users": "array<object>",
      "total": "int",
      "page": "int"
    },
    "status": "string",
    "message": "string"
  }' \
  --output types/api.ts
```

### Batch Generation

```bash
# Generate multiple files
for schema in schemas/*.json; do
  name=$(basename "$schema" .json)
  devtoolbox generate --template go-struct \
    --input-file "$schema" \
    --output "models/${name}.go"
done
```

## 🔌 Adding Custom Plugins

### 1. Create a Plugin

Create `plugins/custom/java_class.py`:

```python
import json

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

### 2. Register the Plugin

```bash
# Add plugin
devtoolbox plugin add ./plugins/custom/java_class.py

# List plugins
devtoolbox plugin list
```

### 3. Use the Plugin

```bash
# Generate Java class
devtoolbox generate --template java_class \
  --input '{
    "className": "User",
    "fields": {
      "name": "string",
      "age": "int",
      "active": "bool"
    }
  }'
```

## 🛠️ Development Workflow

### 1. Start Development Environment

```bash
# Start all services
make docker-dev-up

# Or start individually
go run ./cmd/web &          # Backend on :8080
cd frontend && npm run dev  # Frontend on :3000
```

### 2. Make Changes

- **Backend**: Edit Go files in `internal/`
- **Frontend**: Edit React files in `frontend/src/`
- **Plugins**: Edit Python files in `plugins/`

### 3. Test Changes

```bash
# Run tests
make test

# Test specific components
go test ./internal/core/...
cd frontend && npm test
python -m pytest tests/python/
```

### 4. Build and Deploy

```bash
# Build for production
make build

# Build Docker images
make docker-build

# Deploy
make docker-up
```

## 📊 Monitoring and Debugging

### Check Service Health

```bash
# Backend health
curl http://localhost:8080/health

# Frontend
curl http://localhost:3000

# Docker services
docker-compose ps
```

### View Logs

```bash
# Docker logs
make docker-logs

# Individual service logs
docker-compose logs backend
docker-compose logs frontend
```

### Debug Mode

```bash
# Start with debug logging
GIN_MODE=debug go run ./cmd/web

# Frontend with debug
cd frontend && npm run dev -- --debug
```

## 🎨 Customization

### Frontend Styling

Edit `frontend/src/index.css` for global styles or component-specific CSS files.

### Backend Configuration

Set environment variables:

```bash
export PORT=3000
export GIN_MODE=release
export LOG_LEVEL=info
```

### Plugin Configuration

Edit `~/.devtoolbox/plugins.json`:

```json
{
  "plugins": [
    {
      "name": "my_plugin",
      "path": "./plugins/custom/my_plugin.py",
      "description": "My custom generator"
    }
  ]
}
```

## 🚨 Troubleshooting

### Common Issues

#### Port Already in Use

```bash
# Find and kill process
lsof -i :8080
kill -9 <PID>

# Use different port
export PORT=8081
go run ./cmd/web
```

#### Docker Issues

```bash
# Clean up Docker
docker system prune -f

# Rebuild images
make docker-build
```

#### Python Plugin Errors

```bash
# Test plugin directly
echo '{"test": "data"}' | python3 plugins/custom/my_plugin.py

# Check Python version
python3 --version
```

#### Frontend Build Issues

```bash
# Clear cache and reinstall
cd frontend
rm -rf node_modules package-lock.json
npm install
npm run build
```

### Getting Help

- 📖 Check the [full documentation](docs/)
- 🐛 [Report issues](https://github.com/JIIL07/devtoolbox/issues)
- 💬 [Join discussions](https://github.com/JIIL07/devtoolbox/discussions)
- 📧 Contact: your-email@example.com

## 🎉 Next Steps

Now that you're up and running:

1. 📖 Explore the [full documentation](docs/)
2. 🔌 Create your own [plugins](plugins-guide.md)
3. 🚀 Learn about [advanced features](api-reference.md)
4. 🤝 [Contribute](contributing.md) to the project
5. ⭐ [Star the repository](https://github.com/JIIL07/devtoolbox) if you like it!

## 📚 Additional Resources

- [Installation Guide](installation.md)
- [Plugins Guide](plugins-guide.md)
- [API Reference](api-reference.md)
- [Contributing Guide](contributing.md)
- [Examples Repository](https://github.com/JIIL07/devtoolbox-examples)

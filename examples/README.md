# DevToolBox Examples

This directory contains example JSON schemas that demonstrate various use cases for DevToolBox code generation.

## 📁 Example Files

### Basic Examples

- **`simple-schema.json`** - Basic user schema with common fields
- **`user-schema.json`** - Comprehensive user model with nested objects and arrays

### API Examples

- **`api-response-schema.json`** - Standard API response format with pagination

### Database Examples

- **`database-schema.json`** - Database table schema for products

### Configuration Examples

- **`config-schema.json`** - Application configuration with multiple sections

### Complex Examples

- **`nested-objects-schema.json`** - Complex nested objects with multiple levels

## 🚀 How to Use

### Using the Web Interface

1. Open http://localhost:3000
2. Copy any example JSON from the files below
3. Paste into the input field
4. Select a generator (Go Struct, TypeScript Interface, etc.)
5. Click "Generate Code"

### Using the CLI

```bash
# Generate Go struct from simple schema
devtoolbox generate --template go-struct --input-file examples/simple-schema.json

# Generate TypeScript interface from user schema
devtoolbox generate --template ts-interface --input-file examples/user-schema.json

# Save to file
devtoolbox generate --template go-struct --input-file examples/api-response-schema.json --output api_response.go
```

### Using the API

```bash
# Generate via API
curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d '{
    "template": "go-struct",
    "input": "'"$(cat examples/simple-schema.json | tr -d '\n')"'"
  }'
```

## 📋 Example Schemas

### Simple Schema

```json
{
  "name": "string",
  "age": "int",
  "email": "string",
  "active": "bool"
}
```

**Generated Go Struct:**
```go
type GeneratedStruct struct {
    Active bool   `json:"active"`
    Age    string `json:"age"`
    Email  string `json:"email"`
    Name   string `json:"name"`
}
```

**Generated TypeScript Interface:**
```typescript
interface GeneratedInterface {
  active: boolean;
  age: number;
  email: string;
  name: string;
}
```

### User Schema

```json
{
  "id": "int",
  "username": "string",
  "email": "string",
  "password_hash": "string",
  "first_name": "string",
  "last_name": "string",
  "avatar_url": "string",
  "is_active": "bool",
  "is_verified": "bool",
  "created_at": "string",
  "updated_at": "string",
  "last_login": "string",
  "preferences": {
    "theme": "string",
    "language": "string",
    "notifications": "bool",
    "email_notifications": "bool"
  },
  "roles": "array<string>",
  "permissions": "array<string>"
}
```

### API Response Schema

```json
{
  "data": {
    "users": "array<object>",
    "total": "int",
    "page": "int",
    "limit": "int"
  },
  "status": "string",
  "message": "string",
  "timestamp": "string",
  "request_id": "string"
}
```

### Database Schema

```json
{
  "tableName": "products",
  "columns": {
    "id": "int",
    "name": "string",
    "description": "string",
    "price": "float",
    "category_id": "int",
    "in_stock": "bool",
    "created_at": "string",
    "updated_at": "string"
  },
  "indexes": "array<string>",
  "constraints": "array<string>"
}
```

### Configuration Schema

```json
{
  "app": {
    "name": "string",
    "version": "string",
    "debug": "bool",
    "port": "int"
  },
  "database": {
    "host": "string",
    "port": "int",
    "name": "string",
    "username": "string",
    "password": "string",
    "ssl": "bool",
    "max_connections": "int"
  },
  "redis": {
    "host": "string",
    "port": "int",
    "password": "string",
    "db": "int"
  },
  "features": {
    "analytics": "bool",
    "notifications": "bool",
    "caching": "bool",
    "rate_limiting": "bool"
  },
  "logging": {
    "level": "string",
    "file": "string",
    "max_size": "int",
    "max_backups": "int"
  }
}
```

### Nested Objects Schema

```json
{
  "company": {
    "id": "int",
    "name": "string",
    "address": {
      "street": "string",
      "city": "string",
      "country": "string",
      "postal_code": "string"
    },
    "contact": {
      "email": "string",
      "phone": "string",
      "website": "string"
    }
  },
  "employees": "array<object>",
  "departments": {
    "engineering": {
      "head": "string",
      "size": "int",
      "budget": "float"
    },
    "marketing": {
      "head": "string",
      "size": "int",
      "budget": "float"
    }
  },
  "settings": {
    "timezone": "string",
    "currency": "string",
    "language": "string",
    "features": "array<string>"
  }
}
```

## 🎯 Use Cases

### API Development
Use `api-response-schema.json` to generate response models for REST APIs.

### Database Design
Use `database-schema.json` to generate database models and schemas.

### Configuration Management
Use `config-schema.json` to generate configuration classes.

### User Management
Use `user-schema.json` to generate user models with authentication.

### Complex Data Structures
Use `nested-objects-schema.json` to generate complex nested models.

## 🔧 Customization

You can modify these examples or create your own schemas:

1. **Copy an existing schema** that's similar to your needs
2. **Modify the fields** to match your requirements
3. **Add or remove nested objects** as needed
4. **Test with different generators** to see various outputs

## 📚 Learn More

- [User Guide](../docs/user-guide.md) - Complete usage guide
- [Plugins Guide](../docs/plugins-guide.md) - Create custom generators
- [API Reference](../docs/api-reference.md) - API documentation
- [Quick Start](../docs/quick-start.md) - Get started quickly

## 🤝 Contributing

Have a great example schema? Submit a pull request!

1. Create a new JSON file with a descriptive name
2. Add it to this README with a brief description
3. Include example generated code
4. Submit a pull request

---

**Happy coding with DevToolBox!** 🎉

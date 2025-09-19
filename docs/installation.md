# Installation Guide

## Prerequisites

Before installing DevToolBox, ensure you have the following installed:

- **Go 1.23+** - [Download](https://golang.org/dl/)
- **Node.js 18+** - [Download](https://nodejs.org/)
- **Python 3.10+** - [Download](https://python.org/downloads/)
- **Docker & Docker Compose** - [Download](https://docker.com/get-started/)

## Installation Methods

### Method 1: Docker (Recommended)

The easiest way to get started with DevToolBox is using Docker:

```bash
# Clone the repository
git clone https://github.com/JIIL07/devtoolbox.git
cd devtoolbox

# Start all services
make docker-up

# Or for development
make docker-dev-up
```

### Method 2: From Source

#### 1. Clone Repository

```bash
git clone https://github.com/JIIL07/devtoolbox.git
cd devtoolbox
```

#### 2. Install Dependencies

```bash
# Install Go dependencies
go mod download

# Install frontend dependencies
cd frontend
npm install
cd ..

# Install Python dependencies
pip install -r requirements.txt
```

#### 3. Build Project

```bash
# Build CLI
make build

# Build frontend
cd frontend
npm run build
cd ..
```

#### 4. Run Services

```bash
# Start backend
go run ./cmd/web

# Start frontend (in another terminal)
cd frontend
npm run dev
```

### Method 3: Pre-built Binaries

Download pre-built binaries from the [Releases](https://github.com/JIIL07/devtoolbox/releases) page:

```bash
# Download and extract
wget https://github.com/JIIL07/devtoolbox/releases/latest/download/devtoolbox-linux-amd64.tar.gz
tar -xzf devtoolbox-linux-amd64.tar.gz

# Make executable
chmod +x devtoolbox

# Run
./devtoolbox --help
```

## Verification

After installation, verify everything works:

```bash
# Test CLI
./bin/devtoolbox --help

# Test web interface
curl http://localhost:8080/health

# Test frontend
curl http://localhost:3000
```

## Troubleshooting

### Common Issues

#### Port Already in Use

If you get "port already in use" errors:

```bash
# Find process using port 8080
lsof -i :8080

# Kill process
kill -9 <PID>

# Or use different ports
export PORT=8081
go run ./cmd/web
```

#### Python Plugin Issues

If Python plugins fail:

```bash
# Check Python version
python3 --version

# Install dependencies
pip3 install -r requirements.txt

# Test plugin directly
echo '{"name": "string"}' | python3 plugins/official/ts_interface_gen.py
```

#### Docker Issues

If Docker fails to start:

```bash
# Check Docker status
docker --version
docker-compose --version

# Clean up
docker system prune -f

# Rebuild
make docker-build
```

### Getting Help

- 📖 [Documentation](docs/)
- 🐛 [Report Issues](https://github.com/JIIL07/devtoolbox/issues)
- 💬 [Discussions](https://github.com/JIIL07/devtoolbox/discussions)
- 📧 [Contact](mailto:your-email@example.com)

## Next Steps

After installation:

1. 📖 Read the [User Guide](user-guide.md)
2. 🔌 Learn about [Plugins](plugins-guide.md)
3. 🚀 Try the [Quick Start](quick-start.md)
4. 🛠️ Explore the [API Reference](api-reference.md)

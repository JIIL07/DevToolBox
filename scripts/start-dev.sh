#!/bin/bash

echo "🚀 Starting DevToolBox Development Environment..."

echo "📦 Building Docker images..."
docker-compose -f docker-compose.dev.yml build

echo "🔄 Starting services..."
docker-compose -f docker-compose.dev.yml up -d

echo "⏳ Waiting for services to be ready..."
sleep 10

echo "✅ DevToolBox is running!"
echo ""
echo "🌐 Frontend: http://localhost:3000"
echo "🔧 Backend API: http://localhost:8080"
echo "🗄️  PostgreSQL: localhost:5433"
echo ""
echo "📋 Available commands:"
echo "  make docker-logs    - View logs"
echo "  make docker-dev-down - Stop services"
echo "  make docker-clean   - Clean up everything"
echo ""
echo "🎉 Happy coding!"

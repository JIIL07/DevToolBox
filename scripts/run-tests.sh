#!/bin/bash

set -e

echo "🧪 Running all tests for DevToolBox..."

echo ""
echo "📦 Testing Go backend..."
cd "$(dirname "$0")/.."
go test ./... -v

echo ""
echo "🎨 Testing TypeScript frontend..."
cd frontend
npm test -- --run

echo ""
echo "🐍 Testing Python plugins..."
cd ../tests/python
python -m pytest -v

echo ""
echo "✅ All tests completed successfully!"

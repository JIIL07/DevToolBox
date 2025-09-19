# DevToolBox Makefile
# Универсальные команды для сборки, тестирования и запуска проекта

.PHONY: help build test run clean docker-up docker-down release-build

# Переменные
BINARY_NAME=devtoolbox
WEB_BINARY_NAME=devtoolbox-web
GO_VERSION=1.21
NODE_VERSION=18

# Цвета для вывода
GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m # No Color

help: ## Показать справку по командам
	@echo "$(GREEN)DevToolBox - Генератор кода$(NC)"
	@echo ""
	@echo "$(YELLOW)Доступные команды:$(NC)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(GREEN)%-15s$(NC) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Собрать все бинарники
	@echo "$(GREEN)Сборка CLI...$(NC)"
	@go build -o bin/$(BINARY_NAME) ./cmd/cli
	@echo "$(GREEN)Сборка веб-сервера...$(NC)"
	@go build -o bin/$(WEB_BINARY_NAME) ./cmd/web
	@echo "$(GREEN)Сборка завершена!$(NC)"

build-cli: ## Собрать только CLI
	@echo "$(GREEN)Сборка CLI...$(NC)"
	@go build -o bin/$(BINARY_NAME) ./cmd/cli

build-web: ## Собрать только веб-сервер
	@echo "$(GREEN)Сборка веб-сервера...$(NC)"
	@go build -o bin/$(WEB_BINARY_NAME) ./cmd/web

test: ## Запустить все тесты
	@echo "$(GREEN)Запуск всех тестов...$(NC)"
	@bash scripts/run-tests.sh

test-all: ## Запустить все тесты (альтернативная команда)
	@echo "$(GREEN)Запуск Go тестов...$(NC)"
	@go test -v ./...
	@echo "$(GREEN)Запуск фронтенд тестов...$(NC)"
	@cd frontend && npm test
	@echo "$(GREEN)Запуск Python тестов...$(NC)"
	@python -m pytest tests/python/ -v

test-go: ## Запустить только Go тесты
	@echo "$(GREEN)Запуск Go тестов...$(NC)"
	@go test -v ./...

test-frontend: ## Запустить только фронтенд тесты
	@echo "$(GREEN)Запуск фронтенд тестов...$(NC)"
	@cd frontend && npm test

test-python: ## Запустить только Python тесты
	@echo "$(GREEN)Запуск Python тестов...$(NC)"
	@python -m pytest tests/python/ -v

run: ## Запустить веб-сервер
	@echo "$(GREEN)Запуск веб-сервера...$(NC)"
	@go run ./cmd/web

run-cli: ## Запустить CLI
	@echo "$(GREEN)Запуск CLI...$(NC)"
	@go run ./cmd/cli

dev: ## Запустить в режиме разработки (веб-сервер + фронтенд)
	@echo "$(GREEN)Запуск в режиме разработки...$(NC)"
	@echo "$(YELLOW)Запустите в отдельных терминалах:$(NC)"
	@echo "  make run-web"
	@echo "  make run-frontend"

run-web: ## Запустить только веб-сервер
	@echo "$(GREEN)Запуск веб-сервера на порту 8080...$(NC)"
	@go run ./cmd/web

run-frontend: ## Запустить только фронтенд
	@echo "$(GREEN)Запуск фронтенда на порту 5173...$(NC)"
	@cd frontend && npm run dev

install-deps: ## Установить все зависимости
	@echo "$(GREEN)Установка Go зависимостей...$(NC)"
	@go mod tidy
	@go mod download
	@echo "$(GREEN)Установка фронтенд зависимостей...$(NC)"
	@cd frontend && npm install
	@echo "$(GREEN)Установка Python зависимостей...$(NC)"
	@pip install -r requirements.txt

clean: ## Очистить собранные файлы
	@echo "$(GREEN)Очистка...$(NC)"
	@rm -rf bin/
	@rm -rf frontend/dist/
	@rm -rf frontend/node_modules/
	@go clean

docker-up: ## Запустить все сервисы через Docker Compose
	@echo "$(GREEN)Запуск Docker Compose...$(NC)"
	@docker-compose -f configs/docker-compose.yml up -d

docker-down: ## Остановить все сервисы Docker Compose
	@echo "$(GREEN)Остановка Docker Compose...$(NC)"
	@docker-compose -f configs/docker-compose.yml down

docker-build: ## Собрать Docker образы
	@echo "$(GREEN)Сборка Docker образов...$(NC)"
	@docker-compose -f configs/docker-compose.yml build

release-build: ## Собрать релизные бинарники для всех платформ
	@echo "$(GREEN)Сборка релизных бинарников...$(NC)"
	@mkdir -p dist
	@echo "$(YELLOW)Сборка для Linux...$(NC)"
	@GOOS=linux GOARCH=amd64 go build -o dist/$(BINARY_NAME)-linux-amd64 ./cmd/cli
	@echo "$(YELLOW)Сборка для macOS...$(NC)"
	@GOOS=darwin GOARCH=amd64 go build -o dist/$(BINARY_NAME)-darwin-amd64 ./cmd/cli
	@GOOS=darwin GOARCH=arm64 go build -o dist/$(BINARY_NAME)-darwin-arm64 ./cmd/cli
	@echo "$(YELLOW)Сборка для Windows...$(NC)"
	@GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY_NAME)-windows-amd64.exe ./cmd/cli
	@echo "$(GREEN)Релизные бинарники готовы в папке dist/$(NC)"

lint: ## Запустить линтеры
	@echo "$(GREEN)Запуск Go линтера...$(NC)"
	@golangci-lint run
	@echo "$(GREEN)Запуск фронтенд линтера...$(NC)"
	@cd frontend && npm run lint

format: ## Форматировать код
	@echo "$(GREEN)Форматирование Go кода...$(NC)"
	@go fmt ./...
	@echo "$(GREEN)Форматирование фронтенд кода...$(NC)"
	@cd frontend && npm run format

# Команды по умолчанию
.DEFAULT_GOAL := help

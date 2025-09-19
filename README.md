# 🧰 DevToolBox

> Универсальный генератор кода для разработчиков

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![Node.js Version](https://img.shields.io/badge/Node.js-18+-green.svg)](https://nodejs.org/)
[![Python Version](https://img.shields.io/badge/Python-3.8+-yellow.svg)](https://python.org/)
[![License](https://img.shields.io/badge/License-MIT-red.svg)](LICENSE)
[![Build Status](https://github.com/JIIL07/devtoolbox/workflows/CI/badge.svg)](https://github.com/JIIL07/devtoolbox/actions)

DevToolBox — это мощный инструмент для генерации кода из JSON-схем. Поддерживает множество языков программирования и форматов через систему плагинов.

## ✨ Особенности

- 🚀 **Быстрая генерация** — мгновенное создание кода из JSON
- 🔌 **Система плагинов** — поддержка Python, C++, Go плагинов
- 🌐 **Веб-интерфейс** — удобный UI для генерации кода
- 💻 **CLI** — командная строка для автоматизации
- 🧪 **Тестирование** — полное покрытие тестами
- 🐳 **Docker** — готовые контейнеры для развертывания

## 🎯 Поддерживаемые языки

- **Go** — структуры с JSON тегами
- **TypeScript** — интерфейсы и типы
- **Python** — dataclasses и Pydantic модели
- **Java** — POJO классы
- **C#** — классы с атрибутами
- **Rust** — структуры с serde
- **И многое другое...**

## 📦 Installation

### Через Go (рекомендуется)

```bash
go install github.com/JIIL07/devtoolbox/cmd/cli@latest
```

### Через Docker

```bash
docker pull jiil/devtoolbox:latest
```

### Сборка из исходников

```bash
git clone https://github.com/JIIL07/devtoolbox.git
cd devtoolbox
make install-deps
make build
```

## 🚀 Usage

### CLI

```bash
# Генерация Go структуры
devtoolbox generate --template go-struct --input '{"name":"string","age":"int"}'

# Генерация TypeScript интерфейса
devtoolbox generate --template ts-interface --input '{"name":"string","age":"number"}'

# Использование файла
devtoolbox generate --template go-struct --file schema.json
```

### Веб-интерфейс

```bash
# Запуск веб-сервера
make run-web

# Откройте http://localhost:8080
```

### API

```bash
# POST запрос к API
curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d '{
    "template": "go-struct",
    "input": "{\"name\":\"string\",\"age\":\"int\"}"
  }'
```

## 🔌 Plugins

DevToolBox поддерживает плагины на различных языках:

### Python плагины

```python
# plugins/custom/my_plugin.py
def generate(input_json: str) -> str:
    # Ваша логика генерации
    return generated_code
```

### Go плагины

```go
// plugins/custom/my_plugin.go
package main

func Generate(input string) (string, error) {
    // Ваша логика генерации
    return generatedCode, nil
}
```

### Регистрация плагинов

```bash
# Добавление плагина
devtoolbox plugin add ./plugins/custom/my_plugin.py

# Список плагинов
devtoolbox plugin list

# Удаление плагина
devtoolbox plugin remove my_plugin
```

## 🎬 Demo

### CLI Demo

```bash
$ devtoolbox generate --template go-struct --input '{"user":{"name":"string","email":"string"},"posts":["string"]}'

type Data struct {
    User  User     `json:"user"`
    Posts []string `json:"posts"`
}

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

### Веб-интерфейс

![Web Interface Screenshot](docs/screenshots/web-interface.png)

## 🛠️ Development

### Локальная разработка

```bash
# Клонирование репозитория
git clone https://github.com/yourname/devtoolbox.git
cd devtoolbox

# Установка зависимостей
make install-deps

# Запуск в режиме разработки
make dev
```

### Docker разработка

```bash
# Запуск всех сервисов
make docker-up

# Остановка сервисов
make docker-down
```

### Тестирование

```bash
# Все тесты
make test

# Только Go тесты
make test-go

# Только фронтенд тесты
make test-frontend

# Только Python тесты
make test-python
```

## 📁 Структура проекта

```
devtoolbox/
├── cmd/                    # CLI и веб-приложения
│   ├── cli/               # CLI приложение
│   └── web/               # Веб-сервер
├── internal/              # Внутренние пакеты Go
│   ├── core/              # Ядро генератора
│   ├── api/               # HTTP API handlers
│   └── plugins/           # Загрузчики плагинов
├── frontend/              # React + TypeScript фронтенд
├── plugins/               # Плагины генерации
│   ├── official/          # Официальные плагины
│   └── community/         # Пользовательские плагины
├── tests/                 # Тесты
├── docs/                  # Документация
└── scripts/               # Скрипты сборки
```

## 🤝 Contributing

Мы приветствуем вклад в развитие проекта! Пожалуйста, ознакомьтесь с [руководством по участию](CONTRIBUTING.md).

### Как внести вклад

1. Форкните репозиторий
2. Создайте ветку для новой функции (`git checkout -b feature/amazing-feature`)
3. Зафиксируйте изменения (`git commit -m 'Add amazing feature'`)
4. Отправьте в ветку (`git push origin feature/amazing-feature`)
5. Откройте Pull Request

### Создание плагинов

См. [руководство по созданию плагинов](docs/plugins-guide.md).

## 📄 License

Этот проект лицензирован под MIT License - см. файл [LICENSE](LICENSE) для деталей.

## 🙏 Acknowledgments

- [Go](https://golang.org/) — язык программирования
- [React](https://reactjs.org/) — UI библиотека
- [TypeScript](https://www.typescriptlang.org/) — типизированный JavaScript
- [Gin](https://gin-gonic.com/) — веб-фреймворк для Go

## 📞 Support

- 📧 Email: support@devtoolbox.dev
- 💬 Discord: [DevToolBox Community](https://discord.gg/devtoolbox)
- 🐛 Issues: [GitHub Issues](https://github.com/JIIL07/devtoolbox/issues)
- 📖 Docs: [Documentation](https://docs.devtoolbox.dev)

---

<div align="center">
  <strong>Сделано с ❤️ для разработчиков</strong>
</div>

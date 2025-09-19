# DevToolBox Project Structure

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
│   ├── src/
│   ├── public/
│   └── package.json
├── plugins/               # Плагины генерации
│   ├── official/          # Официальные плагины
│   └── community/         # Пользовательские плагины
├── tests/                 # Тесты
│   ├── go/                # Go тесты
│   ├── frontend/          # Frontend тесты
│   └── python/            # Python тесты
├── docs/                  # Документация
├── scripts/               # Скрипты сборки и тестирования
├── configs/               # Конфигурационные файлы
├── .github/               # GitHub Actions
│   └── workflows/
├── go.mod                 # Go модуль
├── go.sum                 # Go зависимости
├── Makefile               # Команды сборки
├── .gitignore             # Git исключения
└── README.md              # Основная документация
```
# Docker Structure

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │    Backend      │    │   PostgreSQL    │
│   (Nginx)       │◄──►│   (Go + Python) │◄──►│   (Database)    │
│   Port: 80/3000 │    │   Port: 8080    │    │   Port: 5432    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                    ┌─────────────────┐
                    │     Redis       │
                    │   (Cache)       │
                    │   Port: 6379    │
                    └─────────────────┘
```
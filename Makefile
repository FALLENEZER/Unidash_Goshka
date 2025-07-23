include .env
export

APP_NAME := unidashka_goshka
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
MAIN := cmd/server/main.go

## По умолчанию - запуск сервера
.DEFAULT_GOAL := run

## --- { Go } ---

run: watch tailwind ## Запуск go с air и слежкой за templ
	go run $(MAIN)

build: ## Сборка бинарника
	go build -o bin/$(APP_NAME) $(MAIN)

test: ## Запустить все тесты
	go test ./... -v -cover

lint: ## Прогнать линтер (требуется golangci-lint)
	golangci-lint run ./...

sqlc: ## Сгенерировать код через sqlc
	sqlc generate

tidy: ## марафетик в зависимостях
	go mod tidy

## --- { Docker } ---

up: ## Запуск docker-compose (dev)
	docker-compose up -d

down: ## Остановить docker-compose
	docker-compose down

restart: down up ## Перезапуск docker-compose


docker-build: ## Собрать Docker-образ
	docker build -t $(APP_NAME):latest .

docker-build-nocache: ## Собрать Docker-образ (без кэша)
	docker build --no-cache -t $(APP_NAME):latest .

docker-run: ## Запустить через Docker
	docker run --rm -p $(API_PORT):$(API_PORT) --env-file .env $(APP_NAME):latest

## --- { Dev } ---

watch: # Запуск air для автоперезапуска при изменениях с templ
	air

tailwind: # Запуск tailwind
	npx @tailwindcss/cli -i ./web/src/styles/input.css -o ./web/dist/output.css --watch

templ-generate: ## Генерация templ шаблонов
	templ generate internal/template

## --- { Helpers } ---

help: ## Показать список всех доступных команд
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
# Cервис для вычисления арифметических выражений через HTTP-запрос

## Структура проекта

- `cmd/` - директория с main.go
- `internal/` - директория где храниться конфигурация сервера и он сам
- `pkg/calculator/` - директория с реализацией калькулятора

## Запуск

1. Склонируйте проект с GitHub
    ```bash
    git clone https://github.com/Ne-pr1dumal/Calculation-Service-Yandex
    ```
2. Перейдите в головную папку с проектом и запустите проект
    ```bash
    go run ./cmd/main.go
    ```

## Примеры взаимодействия с сервером

1. Запустить сервер
2. Открыть терминал
3. Ввести POST запрос через curl

Вот пример запроса, который разспознается сервером (вместо ... введите свое выражение)

```
 curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"...\"}" http://localhost:8080
```

-**Для запроса**

    curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"(2+2)*2\"}" http://localhost:8080
    
-**Сервер ответит:**
 
    {"result":8}

-**При делении на 0:**

    curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"(6-5)/0 \"}" http://localhost:8080
    
-**Сервер ответит:**

    {"error":"Expression is not valid"}

В случае других ошибок в выражении сервер ответит

    {"error":"Internal server error"}

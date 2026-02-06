# --- Стадия сборки ---
# Используем официальный образ Go как базу для сборки
FROM golang:1.24-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum для загрузки зависимостей
COPY go.mod ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение. Флаги отключают CGO и создают статичный билд
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# --- Финальная стадия ---
# Используем минимальный образ Alpine Linux
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированный бинарный файл из стадии сборки
COPY --from=builder /app/main .

# Открываем порт 8080
EXPOSE 8080

# Команда для запуска приложения
CMD ["./main"]

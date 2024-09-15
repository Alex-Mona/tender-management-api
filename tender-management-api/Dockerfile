# Используем официальный образ Go как базовый
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum файлы для установки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код в рабочую директорию
COPY . .

# Устанавливаем архитектуру и ОС
ENV GOARCH=amd64
ENV GOOS=linux

# Собираем приложение
RUN go build -o main main.go

# Используем Ubuntu для минимального образа
FROM ubuntu:22.04

# Устанавливаем рабочую директорию
WORKDIR /root/

# Устанавливаем необходимые зависимости
RUN apt-get update && apt-get install -y \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

# Копируем собранное приложение из предыдущего этапа
COPY --from=builder /app/main .

# Устанавливаем права на выполнение для файла main
RUN chmod +x main

# Открываем порт 8080
EXPOSE 8080

# Копируем .env файл в рабочую директорию
COPY .env .env

# Запускаем приложение
CMD ["./main"]

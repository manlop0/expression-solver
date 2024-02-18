
# Этап 1: Сборка бинарного файла
FROM golang:1.21.6-alpine3.18 AS builder

WORKDIR /app

COPY . .

# Загрузка зависимостей
RUN go mod download

# Сборка приложения
WORKDIR /app/cmd
RUN go build -o orchestrator .

# Этап 2: Итоговый образ
FROM alpine:3.14

WORKDIR /app

# Копирование бинарного файла из предыдущего этапа
COPY --from=builder /app/cmd/orchestrator .

# Открываем порт, на котором работает приложение
EXPOSE 8000

# Запуск приложения
CMD [ "./orchestrator" ]
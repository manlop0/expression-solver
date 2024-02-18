
# Этап 1: Сборка бинарного файла
FROM golang:1.21.6-alpine3.18 AS builder

WORKDIR /app

COPY . .

# Загрузка зависимостей
RUN go mod download

# Сборка приложения
WORKDIR /app/cmd
RUN go build -o agent .

# Этап 2: Итоговый образ
FROM alpine:3.14

WORKDIR /app

# Копирование бинарного файла из предыдущего этапа
COPY --from=builder /app/cmd/agent .

# Открываем порт, на котором работает приложение
EXPOSE 8010

# Запуск приложения
CMD [ "./agent" ]
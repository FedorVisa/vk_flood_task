FROM golang:latest

LABEL authors="kotozavr"

WORKDIR /app

# Скопируйте файлы проекта в текущую рабочую директорию в контейнере
COPY . .

# Установите зависимости вашего проекта
RUN go mod download

# Соберите вашу программу
RUN go build -o main .

# Экспортируйте переменную окружения, чтобы ваше приложение могло найти Redis
ENV REDIS_ADDR=redis:6379

# Команда для запуска вашей программы
CMD ["./main"]

# Используем базовый образ с Python 3.11
FROM python:3.11-slim

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем исходный код приложения в контейнер
COPY server.py .

# Открываем порт, который будет использоваться приложением
EXPOSE 11111

# Определяем команду для запуска сервера
CMD ["python", "server.py"]

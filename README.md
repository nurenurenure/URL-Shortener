# URL Shortener Service

Сервис сокращения ссылок, реализованный на Go с использованием REST API и PostgreSQL. Позволяет сохранять длинные URL и получать короткие ссылки с последующим редиректом.

---

## Стек технологий

- Go (Golang)
- PostgreSQL
- Docker / Docker Compose
- REST API (HTTP, JSON)

---

## Функциональность

- Создание короткой ссылки
- Перенаправление по короткому URL
- Хранение данных в PostgreSQL
- Обработка ошибок и корректные HTTP-статусы
- Логирование запросов
- Конфигурация через YAML

---

## Структура проекта

- cmd/url-shortener/        # точка входа
- internal/config/          # конфигурация
- internal/http-server/     # HTTP handlers
- internal/storage/         # работа с БД
- internal/lib/             # утилиты (логгер, генерация, ответы)
- storage/                  # docker-compose

## Запуск проекта
1. Клонировать репозиторий
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
2. Запуск через Docker
docker-compose up --build
## Примеры API
Создание короткой ссылки
POST /url
Content-Type: application/json

{
  "url": "https://example.com"
}
Ответ:
{
  "short_url": "http://localhost:8080/abc123"
}
##Конфигурация

Настройки задаются в config/config.yaml:
- порт сервера
- параметры подключения к БД
- прочие настройки
## Архитектура
- Разделение на слои: handlers / storage / config / lib
- Использование goroutines и context для обработки запросов
- Чистая структура проекта с разделением ответственности

##Дальнейшее развитие
- Добавление авторизации
- Ограничение количества запросов (rate limiting)
- Кэширование (Redis)
- Метрики и мониторинг

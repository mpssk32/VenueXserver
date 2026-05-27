# VenueX

Микросервисная платформа для организации концертов, продажи билетов и общения пользователей.

## Возможности

### Auth Service

* Регистрация пользователей
* Авторизация через JWT
* Роли пользователей:

  * user
  * artist
  * venue_admin
  * admin
  * super_admin
* Middleware авторизации
* Middleware ролей
* Аудит действий
* RabbitMQ notifications

### Concert Service

* Создание площадок
* Создание концертов
* Создание событий
* Покупка билетов
* Управление заявками артистов
* Поиск событий

### Chat Service

* WebSocket чат
* JWT авторизация для WebSocket
* История сообщений
* Личные сообщения между пользователями

---

# Архитектура

```text
VenueX
│
├── auth-service
├── concert-service
├── chat-service
└── shared
    ├── jwt
    └── logger
    
```

---

# Технологии

* Go
* PostgreSQL
* RabbitMQ
* Docker
* Docker Compose
* JWT
* Gorilla Mux
* Gorilla WebSocket
* Swagger
* Zap Logger
* Testify

---

# Микросервисы

## Auth Service

Порт: `8080`

### Основные endpoints

| Method | Endpoint          | Description                  |
| ------ | ----------------- | ---------------------------- |
| POST   | /register         | Регистрация                  |
| POST   | /login            | Авторизация                  |
| GET    | /me               | Информация о себе            |
| GET    | /admin            | Admin panel                  |
| GET    | /admin/users      | Получение всех пользователей |
| DELETE | /admin/users/{id} | Удаление пользователя        |

Swagger:

```text
http://localhost:8080/swagger/index.html
```

---

## Concert Service

Порт: `8082`

### Основные endpoints

| Method | Endpoint       | Description       |
| ------ | -------------- | ----------------- |
| GET    | /events        | Все события       |
| GET    | /events/{id}   | Событие по ID     |
| GET    | /events/search | Поиск событий     |
| POST   | /events        | Создание события  |
| POST   | /venues        | Создание площадки |
| POST   | /concerts      | Создание концерта |
| POST   | /tickets       | Покупка билета    |
| GET    | /my-tickets    | Мои билеты        |

Swagger:

```text
http://localhost:8082/swagger/index.html
```

---

## Chat Service

Порт: `8081`

### Основные endpoints

| Method | Endpoint  | Description       |
| ------ | --------- | ----------------- |
| GET    | /ws       | WebSocket чат     |
| GET    | /messages | История сообщений |

Swagger:

```text
http://localhost:8081/swagger/index.html
```

---

# Shared Module

## shared/logger

Общий logger на базе zap.

## shared/jwt

Общий JWT middleware и JWT validation.

## shared/rabbitmq

Общее подключение к RabbitMQ.

---

# Запуск проекта

## 1. Клонирование

```bash
git clone <repo_url>
cd VenueX
```

---

## 2. Переменные окружения

### Auth Service

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/concert_platform?sslmode=disable
JWT_SECRET=super_secret_key
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
```

### Concert Service

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/concert_platform?sslmode=disable
JWT_SECRET=super_secret_key
```

### Chat Service

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/concert_platform?sslmode=disable
JWT_SECRET=super_secret_key
```

---

# Docker запуск

```bash
docker compose up --build
```

---

# Локальный запуск

## Auth Service

```bash
cd auth-service
go run cmd/main.go
```

## Concert Service

```bash
cd concert-service
go run cmd/main.go
```

## Chat Service

```bash
cd chat-service
go run cmd/main.go
```

---

# Тестирование

## Auth Service

```bash
go test ./... -cover
```

## Concert Service

```bash
go test ./... -cover
```

## Chat Service

```bash
go test ./... -cover
```

---

# WebSocket подключение

Пример подключения:

```text
ws://localhost:8081/ws?token=JWT_TOKEN
```

Пример сообщения:

```json
{
  "sender_id": "uuid",
  "receiver_id": "uuid",
  "content": "hello"
}
```

---

# RabbitMQ Notifications

Auth-service публикует события:

* регистрация пользователя
* audit logs
* notifications

Consumer:

* concert-service

---

# Логирование

Используется общий zap logger:

```go
logger.Log.Info("service started")
```

---

# Безопасность

* JWT Authentication
* Role-based access
* Middleware validation
* Password hashing через bcrypt
* Protected admin routes


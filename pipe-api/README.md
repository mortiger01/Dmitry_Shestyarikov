# API для магазина сантехнических труб (Лабораторные работы 2–4)

REST API на Go, реализующий CRUD для труб, JWT-аутентификацию, OAuth2 Яндекс ID и автоматическую OpenAPI-документацию.

## Быстрый старт

1. Скопируйте `.env.example` в `.env` и заполните реквизиты (особенно секреты JWT и Yandex OAuth).
2. Сгенерируйте документацию: `swag init -g cmd/main.go -o docs` (требуется `go install github.com/swaggo/swag/cmd/swag@latest`).
3. Запустите всё: `docker-compose up --build`.
4. Swagger UI будет доступен по адресу `http://localhost:4200/api/docs/index.html` (только при `APP_ENV=development`). В production эндпоинт `/api/docs` отключён.

## Основные эндпоинты

- **Auth**
  - `POST /auth/register` – регистрация
  - `POST /auth/login` – вход (устанавливает HttpOnly cookie)
  - `POST /auth/refresh` – обновление токенов
  - `GET /auth/whoami` – профиль текущего пользователя
  - `POST /auth/logout` / `POST /auth/logout-all`
  - OAuth: `GET /auth/oauth/yandex`, `GET /auth/oauth/yandex/callback`
- **Трубы** (требуется авторизация)
  - `GET /pipes?page=1&limit=10` – список труб пользователя с пагинацией
  - `GET /pipes/:id` – одна труба
  - `POST /pipes` – создать
  - `PUT /pipes/:id` – полное обновление
  - `PATCH /pipes/:id` – частичное обновление
  - `DELETE /pipes/:id` – мягкое удаление (soft delete)

## Безопасность

- Пароли хешируются bcrypt с уникальной солью для каждого пользователя.
- Токены передаются только в HttpOnly cookies.
- Refresh-токены хранятся в БД в хешированном виде и могут быть отозваны.
- OAuth2 state проверяется для защиты от CSRF.
- Чувствительные поля (пароль, соль, хеши) исключены из Swagger-схем.

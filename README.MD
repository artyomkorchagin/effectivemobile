# Задание для Effectivemobile

RESTful API для управления подписками пользователей, реализованный на Go, с использованием Gin, PostgreSQL и документации через Swagger.

Этот сервис поддерживает:
- Операции CRUDL для подписок
- Фильтрацию и подсчёт суммы подписок
- Документацию API через Swagger
- Запуск через Docker
- Инициализация базы данных через миграции

---

## Требования

- Docker и Docker Compose
- Go 1.23
- Установленный `make`
- `swag`

Установите `swag`:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

---

## Начало работы

### 1. Клонируйте репозиторий

```bash
git clone https://github.com/artyomkorchagin/effectivemobile.git 
cd effectivemobile
```

### 2. Соберите и запустите приложение

```bash
make build
make up
```

Swagger API будет доступен по адресу:

http://localhost:3000/swagger/index.html

---

## Сервисы в Docker Compose

| Сервис     | Описание                          | Порт      |
|------------|-----------------------------------|-----------|
| `db`       | PostgreSQL база данных            | 5432      |
| `app`      | Go Gin API сервер                 | 3000      |

---

## Доступные команды Make

| Команда       | Описание                           |
|----------------|------------------------------------|
| `make build`   | Собрать Docker-образы              |
| `make up`      | Запустить всю инфраструктуру       |
| `make down`    | Остановить и удалить контейнеры    |
| `make restart` | Пересобрать и перезапустить всё    |
| `make migrate` | Применить миграции БД              |
| `make run`     | Запустить приложение без Docker    |
| `make test`    | Запустить тесты                    |
| `make clean`   | Очистить Docker артефакты          |

---

## API Эндпоинты

| Метод  | Путь                      | Описание                         |
|--------|---------------------------|----------------------------------|
| POST   | `/subscriptions`          | Создать новую подписку           |
| GET    | `/subscriptions/{id}`     | Получить подписку по ID          |
| PATCH  | `/subscriptions/{id}`     | Частично обновить подписку       |
| DELETE | `/subscriptions/{id}`     | Удалить подписку                 |
| GET    | `/subscriptions`          | Получить список всех подписок    |
| GET    | `/subscriptions/sum`      | Получить сумму подписок          |

Вся документация доступна через Swagger UI.

---

## Пример запроса (PATCH)

```json
{
  "service_name": "Netflix",
  "price": 1000,
  "start_date": "01-2025"
}
```

---

## Пример ответа

```json
{
  "id": 1,
  "service_name": "Netflix",
  "price": 1000,
  "user_id": "123e4567-e89b-42d3-a456-556642440000",
  "start_date": "01-2025",
  "end_date": ""
}
```

---

## Формат ошибок

```json
{
  "code": 400,
  "error": "invalid ID"
}
```

---

## Запуск тестов

```bash
make test
```

---

## Контакт

Артём Корчагин  
artyomkorchagin333@gmail.com

--- 
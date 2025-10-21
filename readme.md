# 🦦 RealWorld API — Go + Gin Implementation

Реализация **[RealWorld API Specification](https://realworld-docs.netlify.app/specifications/backend/introduction/)** на **Go (Golang)** с использованием **Gin Framework**.

---

## 🚀 Особенности

- RESTful API, совместимый со [спецификацией RealWorld](https://github.com/gothinkster/realworld)
- Аутентификация с использованием JWT
- CRUD для статей, комментариев, пользователей
- Middleware для авторизации
- Поддержка конфигурации через `.env` или `config.yaml`
- Подключение к PostgreSQL через GORM
- Поддержка hot-reload режима разработки

---

## 📦 Технологии

- [Go](https://go.dev/) — 1.21+
- [Gin](https://gin-gonic.com/) — HTTP web framework 
- [GORM](https://gorm.io/) — ORM для PostgreSQL
- [godotenv](https://github.com/joho/godotenv) — для загрузки переменных окружения
- [air](https://github.com/cosmtrek/air) — live reload для dev режима

---

## ⚙️ Установка и запуск

### 1. Клонируй репозиторий
```bash
git clone https://github.com/g28xyz/realworld-go-gin.git
cd realworld-go-gin
make dev
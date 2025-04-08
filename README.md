# Feedback API

A simple CRUD REST API for managing feedback using Go, Fiber framework, GORM ORM, and PostgreSQL.

## 🚀 Tech Stack

- Go (Golang)
- Fiber (Express-like web framework for Go)
- GORM (Go ORM)
- PostgreSQL
- UUID (for unique feedback IDs)
- Docker

## 📦 Features

- Create feedback
- Get all feedback
- Get feedback by ID
- Update feedback
- Delete feedback

## 🛠️ Getting Started

## Install the project dependencies

```bash
go get github.com/gofiber/fiber/v2
go get github.com/google/uuid
go get github.com/go-playground/validator/v10
go get -u gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/spf13/viper
go install github.com/cosmtrek/air@latest

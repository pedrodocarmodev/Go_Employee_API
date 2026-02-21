

---

# Employee API

Simple REST API for employee management built with **Go**, **Gin**, and **PostgreSQL**, following a **Layered Architecture** approach.

---
## 🏗 Architecture

This project follows a **Layered Architecture** pattern to separate responsibilities and reduce coupling.

Flow:

```
HTTP → Handler → Service → Repository → PostgreSQL
```

The service layer depends on an interface, allowing easy replacement of the persistence layer.

---


## Requirements

* Go
* PostgreSQL

---

## Setup

### 1. Create database

```sql
CREATE DATABASE database;
```

---

### 2. Configure environment variables

Copy the example file:

```bash
cp .env.example .env
```

Fill in your database credentials inside `.env`.

---

### 3. Install dependencies

From project root:

```bash
go mod tidy
```

---

### 4. Run the application

```bash
go run ./cmd/api
```

Server runs at:

```
http://localhost:8080
```

---

## Endpoints

* `GET /employees`
* `GET /employees/:id`
* `POST /employees`
* `PATCH /employees/:id/fire`
* `PATCH /employees/:id/employ`

---



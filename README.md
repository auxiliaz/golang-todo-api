# Golang Todo API

Simple REST API for Todo List built with **Golang** and **MySQL**.

---

## Tech Stack
- Golang (net/http)
- MySQL
- go-sql-driver/mysql

---

## Setup & Run

### 1. Clone Repository
```bash
git clone https://github.com/auxiliaz/golang-todo-api.git
cd golang-todo-api
```

---

### 2. Database Setup (MySQL)
```sql
CREATE DATABASE todo_db;
USE todo_db;

CREATE TABLE todos (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  status ENUM('pending','done') DEFAULT 'pending',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

### 3. Run Application
```bash
go run main.go
```

Server running on:
```
http://localhost:8080
```

---

## Endpoints

### Get All Todos
**GET** `/todos`

**Response 200 OK**
```json
[
  {
    "id": 1,
    "title": "Belajar Golang",
    "description": "CRUD Todo API",
    "status": "pending",
    "created_at": "2026-01-06 18:40:12"
  }
]
```

---

### Get Todo By ID
**GET** `/todo?id=1`

**Response 200 OK**
```json
{
  "id": 1,
  "title": "Belajar Golang",
  "description": "CRUD Todo API",
  "status": "pending",
  "created_at": "2026-01-06 18:40:12"
}
```

**Response 404 Not Found**
```json
{
  "error": "Todo not found"
}
```

---

### Create Todo
**POST** `/todo/create`

**Request Body**
```json
{
  "title": "Belajar Golang",
  "description": "Membuat REST API"
}
```

**Response 201 Created**
```json
{
  "id": 2,
  "title": "Belajar Golang",
  "description": "Membuat REST API",
  "status": "pending",
  "created_at": "2026-01-06 18:45:00"
}
```

---

### Update Todo
**PUT** `/todo/update?id=1`

**Request Body**
```json
{
  "title": "Belajar Golang Updated",
  "description": "Sudah selesai",
  "status": "done"
}
```

**Response 200 OK**
```json
{
  "message": "Todo updated successfully"
}
```

---

### Delete Todo
**DELETE** `/todo/delete?id=1`

**Response 204 No Content**

---

## Notes
- API tested using Postman
- All request and response use JSON format
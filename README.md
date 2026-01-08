# Golang Todo API – RESTful Backend Todo List

Golang Todo API adalah aplikasi **REST API sederhana** untuk manajemen todo list yang dibangun menggunakan **Golang** dan **MySQL**.  
Aplikasi ini mendukung operasi CRUD (Create, Read, Update, Delete).

## Tech Stack

- Bahasa: Golang
- HTTP Server: net/http
- Database: MySQL
- Driver Database: go-sql-driver/mysql

## Menjalankan proyek

Pastikan **Golang** dan **MySQL** sudah terpasang.

1. **Clone repository**
```bash
git clone https://github.com/auxiliaz/golang-todo-api.git
cd golang-todo-api
```
2. **Setup Database (MySQL)**
Buat database dan tabel menggunakan perintah berikut:
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
Pastikan konfigurasi koneksi database di file aplikasi (`main.go`) sudah sesuai dengan:
- host
- username
- password
- nama database (`todo_db`)

3. **Jalankan development server**
```bash
go run main.go
```
Server akan berjalan di:
```bash
[go run main.go](http://localhost:8080
```

## Fitur utama

- Menampilkan seluruh data todo.
- Menampilkan detail todo berdasarkan ID.
- Menambahkan todo baru.
- Memperbarui data todo.
- Menghapus todo.
- Penyimpanan data menggunakan database MySQL.

## Alur kerja aplikasi

1. **Koneksi Database**
   - Aplikasi melakukan koneksi ke database MySQL saat server dijalankan.
   - Semua data todo disimpan dan diambil dari tabel `todos`.

2. **Request dari Client**
   - Client (misalnya Postman atau frontend) mengirim request HTTP ke endpoint API.
   - Data dikirim dan diterima dalam format JSON.

3. **Proses CRUD**
   - Create: Menyimpan todo baru ke database.
   - Read: Mengambil seluruh todo atau todo berdasarkan ID.
   - Update: Memperbarui title, description, atau status todo.
   - Delete: Menghapus todo dari database.

4. **Response API**
   - API mengembalikan response JSON sesuai dengan hasil proses (success / error).

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

### Delete Todo 
**DELETE** `/todo/delete?id=1`
**Response 204 No Content** 

## Struktur penting

- `main.go` – Entry point aplikasi dan routing API.
- `database/` – Konfigurasi dan koneksi MySQL.
- `handlers/` – Logika handler untuk setiap endpoint.
- `models/` – Struktur data todo.

## Pengujian 
- API diuji menggunakan Postman.
- Seluruh request dan response menggunakan format JSON.

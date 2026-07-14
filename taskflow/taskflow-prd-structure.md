# Product Requirement Document (PRD) & Architecture Layout
## TaskFlow API & Application

---

## 1. Ringkasan Produk

TaskFlow adalah aplikasi manajemen tugas (*todo list*) berbasis web. Aplikasi ini mewajibkan pengguna untuk melakukan autentikasi (*register/login*) sebelum dapat mengakses, membuat, serta mengelola daftar tugas harian mereka. Sistem ini menjamin isolasi data pengguna, sehingga setiap akun hanya memiliki akses penuh terhadap data miliknya sendiri.

---

## 2. Tujuan & Arsitektur Sistem

### Tujuan Utama
* Menyediakan platform manajemen tugas yang aman, cepat, dan responsif.
* Menerapkan autentikasi berbasis standar industri untuk mengamankan data pengguna.

### Arsitektur Teknis
* **Backend:** Golang (Fiber / Echo / Gin) untuk performa *high-throughput* dan *low-latency*.
* **Frontend:** JavaScript / TypeScript (Next.js / React / Vue.js) dengan Tailwind CSS untuk *UI/UX* yang modern dan responsif.
* **Database:** PostgreSQL / MySQL (GORM / SQLx) untuk penyimpanan data relasional.
* **Autentikasi:** JSON Web Token (JWT) dengan *Stateless Authentication*.

---

## 3. Fitur Utama & Spesifikasi

### A. Fitur Autentikasi Pengguna
* **Registrasi Akun:** Pengguna baru mendaftar menggunakan nama lengkap, email unik, dan kata sandi.
* **Autentikasi (Login):** Validasi kredensial pengguna dan penerbitan *Access Token* (JWT).
* **Manajemen Sesi:** *Token-based auth* yang disimpan secara aman di sisi klien (*HTTP-Only Cookie* atau *Local Storage*).
* **Proteksi Rute (Middleware):** Mencegah akses ke *dashboard* utama tanpa JWT yang valid.

### B. Fitur Manajemen Tugas (CRUD Core)
* **Buat Tugas Baru:** Menambahkan entitas tugas baru dengan informasi judul, deskripsi, tanggal tenggat (*due date*), serta prioritas.
* **Lihat Daftar Tugas:** Menampilkan daftar tugas milik pengguna dengan opsi *filtering* (Selesai/Belum Selesai) dan *sorting* berdasarkan tanggal.
* **Pembaruan Tugas:** Mengubah status tugas (*completed/pending*) atau memperbarui konten tugas.
* **Hapus Tugas:** Menghapus tugas dari sistem (*soft delete* atau *hard delete*).

---

## 4. Struktur Data (Database Schema)

### Tabel `users`
| Field | Tipe Data | Keterangan |
| :--- | :--- | :--- |
| `id` | UUID / BIGINT | Primary Key, Auto-generated |
| `name` | VARCHAR(100) | Nama lengkap pengguna |
| `email` | VARCHAR(150) | Unique, Not Null, Indexed |
| `password_hash` | VARCHAR(255) | Hash menggunakan bcrypt |
| `created_at` | TIMESTAMP | Waktu pembuatan akun |
| `updated_at` | TIMESTAMP | Waktu perbaruan akun |

### Tabel `todos`
| Field | Tipe Data | Keterangan |
| :--- | :--- | :--- |
| `id` | UUID / BIGINT | Primary Key, Auto-generated |
| `user_id` | UUID / BIGINT | Foreign Key -> `users.id` (ON DELETE CASCADE) |
| `title` | VARCHAR(255) | Judul tugas |
| `description` | TEXT | Detail penjelasan tugas |
| `is_completed` | BOOLEAN | Status penyelesaian (Default: `false`) |
| `priority` | ENUM | Opsi: `LOW`, `MEDIUM`, `HIGH` |
| `due_date` | TIMESTAMP | Tanggal tenggat waktu tugas |
| `created_at` | TIMESTAMP | Waktu tugas dibuat |
| `updated_at` | TIMESTAMP | Waktu tugas diperbarui |

---

## 5. Spesifikasi API Endpoint (Golang Service)

### Autentikasi Endpoint

#### 1. Register User
* **Endpoint:** `POST /api/v1/auth/register`
* **Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "SecurePassword123"
}
```
* **Response (201 Created):**
```json
{
  "code": 201,
  "status": "success",
  "message": "User registered successfully",
  "data": {
    "id": "usr-9102-x812",
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

#### 2. Login User
* **Endpoint:** `POST /api/v1/auth/login`
* **Request Body:**
```json
{
  "email": "john@example.com",
  "password": "SecurePassword123"
}
```
* **Response (200 OK):**
```json
{
  "code": 200,
  "status": "success",
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

### Todo Endpoint (Memerlukan Header `Authorization: Bearer <token>`)

#### 3. Get All Todos
* **Endpoint:** `GET /api/v1/todos`
* **Query Parameters:** `is_completed` (optional), `sort_by` (optional)
* **Response (200 OK):**
```json
{
  "code": 200,
  "status": "success",
  "data": [
    {
      "id": "todo-001",
      "title": "Setup Golang Project Structure",
      "description": "Implement clean architecture with fiber framework",
      "is_completed": false,
      "priority": "HIGH",
      "due_date": "2026-07-20T17:00:00Z"
    }
  ]
}
```

#### 4. Create Todo
* **Endpoint:** `POST /api/v1/todos`
* **Request Body:**
```json
{
  "title": "Fix Auth Middleware Bug",
  "description": "Ensure JWT payload parsing handles expired tokens gracefully",
  "priority": "HIGH",
  "due_date": "2026-07-16T10:00:00Z"
}
```
* **Response (201 Created):**
```json
{
  "code": 201,
  "status": "success",
  "message": "Task created successfully",
  "data": {
    "id": "todo-002",
    "title": "Fix Auth Middleware Bug",
    "is_completed": false
  }
}
```

#### 5. Update Status / Detail Todo
* **Endpoint:** `PUT /api/v1/todos/:id`
* **Request Body:**
```json
{
  "title": "Fix Auth Middleware Bug",
  "is_completed": true,
  "priority": "MEDIUM"
}
```
* **Response (200 OK):**
```json
{
  "code": 200,
  "status": "success",
  "message": "Task updated successfully"
}
```

#### 6. Delete Todo
* **Endpoint:** `DELETE /api/v1/todos/:id`
* **Response (200 OK):**
```json
{
  "code": 200,
  "status": "success",
  "message": "Task deleted successfully"
}
```

---

## 6. Alur Pengguna (User Flow)

1. **Unauthenticated State:**
   * Pengguna mengakses URL utama (`/`).
   * Middleware Frontend mendeteksi ketiadaan token, kemudian melakukan *redirect* otomatis ke halaman `/login`.
2. **Authentication Flow:**
   * Pengguna memilih *Register* jika belum memiliki akun, mengisi form, lalu data dikirim ke backend.
   * Pengguna melakukan *Login*. Sistem backend memverifikasi password via hashing function (Bcrypt), lalu mengembalikan JWT.
   * Klien menyimpan JWT secara aman dan mengarahkan pengguna ke `/dashboard`.
3. **Task Management Flow:**
   * Di `/dashboard`, aplikasi memanggil API `GET /api/v1/todos` menggunakan header `Authorization: Bearer <token>`.
   * Pengguna menambah, mengubah, atau menghapus item *todo*. Setiap perubahan langsung disinkronkan ke database melalui REST API backend Golang.

---

## 7. Kriteria Non-Fungsional (Keamanan & Performa)

* **Keamanan Sandi:** Hashing password menggunakan `Bcrypt` dengan *cost factor* minimal 10.
* **Isolasi Data (Data Privacy):** Setiap *query* SQL wajib menyertakan filter `WHERE user_id = :authenticated_user_id` untuk mencegah kerentanan IDOR (*Insecure Direct Object Reference*).
* **Performa API:** Latensi respons API di bawah 100ms untuk operasi CRUD standar.
* **Validasi Input:** Validasi ganda dilakukan di layer JS (Frontend) dan Golang (Backend validator struct).

---

## 8. Rekomendasi Struktur Folder Project

### Separated Repository (Polyrepo)

#### Backend Repository (`taskflow-backend`)
```text
taskflow-backend/
├── cmd/
│   └── api/
│       └── main.go                 # Entry point aplikasi Go
├── internal/
│   ├── config/
│   │   └── config.go               # Load environment variables (.env)
│   ├── delivery/
│   │   └── http/
│   │       ├── handler/            # Controller / Request Handler (Gin/Fiber/Echo)
│   │       │   ├── auth_handler.go
│   │       │   └── todo_handler.go
│   │       ├── middleware/         # Auth JWT middleware, logger, CORS
│   │       │   └── auth_middleware.go
│   │       └── router/             # Routing registrasi endpoint API
│   │           └── router.go
│   ├── domain/                     # Model structs & interfaces
│   │   ├── auth.go
│   │   └── todo.go
│   ├── repository/                 # Layer akses database (GORM/SQLx)
│   │   ├── user_repository.go
│   │   └── todo_repository.go
│   └── usecase/                    # Business logic layer
│       ├── auth_usecase.go
│       └── todo_usecase.go
├── pkg/                            # Helper / Utilities publik
│   ├── jwt/
│   │   └── jwt.go
│   └── password/
│       └── bcrypt.go
├── db/
│   └── migrations/                 # Script SQL migration (golang-migrate)
│       ├── 000001_create_users_table.up.sql
│       └── 000001_create_users_table.down.sql
├── .env.example
├── Dockerfile
├── go.mod
└── go.sum
```

#### Frontend Repository (`taskflow-frontend`)
```text
taskflow-frontend/
├── public/                         # Static assets (favicon, images, logo)
├── src/
│   ├── app/                        # Route pages (Next.js App Router)
│   │   ├── (auth)/
│   │   │   ├── login/
│   │   │   │   └── page.tsx
│   │   │   └── register/
│   │   │       └── page.tsx
│   │   ├── (dashboard)/
│   │   │   └── dashboard/
│   │   │       └── page.tsx
│   │   ├── layout.tsx
│   │   └── page.tsx
│   ├── components/                 # Reusable UI components
│   │   ├── ui/                     # Basic primitives (Button, Input, Modal)
│   │   └── features/               # Feature-specific components
│   │       ├── auth/
│   │       │   └── login-form.tsx
│   │       └── todo/
│   │           ├── todo-list.tsx
│   │           └── todo-item.tsx
│   ├── lib/                        # Konfigurasi library (Axios/Fetch instance)
│   │   └── api.ts
│   ├── services/                   # Function API call ke backend Go
│   │   ├── auth.service.ts
│   │   └── todo.service.ts
│   ├── types/                      # TypeScript interfaces / type definitions
│   │   ├── auth.ts
│   │   └── todo.ts
│   └── utils/                      # Helper functions (date formatter, token store)
├── .env.local.example
├── Dockerfile
├── package.json
└── tailwind.config.js
```
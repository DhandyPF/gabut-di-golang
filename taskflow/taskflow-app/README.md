# TaskFlow

Aplikasi manajemen tugas dengan autentikasi JWT. Backend Go, frontend Next.js.

## Struktur

- `taskflow-backend.zip`: REST API di Go (net/http, JWT, bcrypt, SQLite).
- `taskflow-frontend.zip`: Aplikasi Next.js (App Router, TypeScript, Tailwind).

Kedua bagian sudah diuji dan berhasil build serta jalan bersamaan di lingkungan sandbox ini.

## Menjalankan backend

```
cd taskflow-backend
go mod download
go run ./cmd/api
```

Server jalan di `http://localhost:8080`. Konfigurasi lewat environment variable, lihat `.env.example`:

- `PORT` (default 8080)
- `DATABASE_DSN` (default ./taskflow.db, file SQLite dibuat otomatis)
- `JWT_SECRET` (WAJIB diganti untuk produksi)

## Menjalankan frontend

```
cd taskflow-frontend
npm install
cp .env.local.example .env.local
npm run dev
```

Aplikasi jalan di `http://localhost:3000`. Variabel `NEXT_PUBLIC_API_URL` di `.env.local` menunjuk ke backend.

## Alur pakai

1. Buka `http://localhost:3000`, klik "Create account".
2. Daftar dengan nama, email, password (minimal 8 karakter).
3. Otomatis masuk ke `/dashboard`.
4. Tambah tugas lewat tombol "+ New task", atur prioritas dan tenggat.
5. Klik lingkaran di kiri tugas untuk menandai selesai, filter dengan tab All/Pending/Completed.

## Endpoint API

| Method | Path | Auth |
|---|---|---|
| POST | /api/v1/auth/register | tidak |
| POST | /api/v1/auth/login | tidak |
| GET | /api/v1/todos | ya |
| POST | /api/v1/todos | ya |
| PUT | /api/v1/todos/:id | ya |
| DELETE | /api/v1/todos/:id | ya |

Setiap endpoint todo membaca `Authorization: Bearer <token>` dan hanya mengembalikan data milik user yang login. Semua query database di-filter dengan `user_id` untuk mencegah IDOR.

## Catatan implementasi

- Database: SQLite file lokal untuk kemudahan setup. Skema dan query sudah kompatibel pola dengan PostgreSQL/MySQL kalau nanti mau migrasi, tinggal ganti driver dan DSN di `internal/repository/db.go`.
- Password di-hash dengan bcrypt, cost factor 10.
- Token JWT berlaku 24 jam.
- Middleware proteksi rute ada di dua sisi: backend (`internal/delivery/http/middleware/auth_middleware.go`) dan frontend (`src/app/(dashboard)/layout.tsx` redirect ke `/login` kalau tidak ada token).

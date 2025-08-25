# CRUD Go + Next.js

Project ini adalah contoh aplikasi **CRUD** sederhana menggunakan:
- **Backend** → Go (Gin + GORM) + PostgreSQL  
- **Frontend** → Next.js (App Router) + TailwindCSS  

## 🚀 Cara Menjalankan

## Setup Backend (Go)

### 1. Clone Repo
```bash
git clone https://github.com/mfthrq/crud-go-nextjs.git
cd crud-go-nextjs
```

### 2. Setup Backend (Go + PostgreSQL)
```bash
cd backend
go mod tidy
```

### 3. Buat file backend/.env
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=users
```
### 4. Jalankan server go
```bash
go run main.go
```

## Setup Frontend (Next.js)

### 1. Masuk ke folder frontend
```bash
cd frontend
```

### 2. Masuk ke folder frontend
```bash
npm install
```

### 3. Buat file .env.local
```bash
NEXT_PUBLIC_API_URL=http://localhost:9090 #sesuaikan dengan localhost kamu
```

### 4. Jalankan frontend
```bash
npm run dev
```



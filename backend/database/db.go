package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Gagal load .env, pastikan variabel lingkungan sudah di-set")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal ping ke database:", err)
	}

	fmt.Println("✅ Berhasil konek ke PostgreSQL!")
}

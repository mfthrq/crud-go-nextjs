package main

import (
	"backend/database"
	"backend/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	database.Connect()

	r := mux.NewRouter()

	// Definisikan route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello dari Backend Go + PostgreSQL!")
	}).Methods("GET")

	r.HandleFunc("/users", routes.GetUsers).Methods("GET")
	r.HandleFunc("/users", routes.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", routes.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")

	// Tambahkan middleware CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Metode yang diperbolehkan
		AllowedHeaders: []string{"Content-Type"},                 // Header yang diperbolehkan
	})

	// Gunakan CORS handler
	handler := c.Handler(r)

	// Jalankan server
	fmt.Println("🚀 Server jalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

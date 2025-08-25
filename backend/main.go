package main

import (
	"backend/database"
	"backend/models"
	"backend/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Koneksi DB
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	// Router
	r := mux.NewRouter()

	// Routes user
	r.HandleFunc("/users", routes.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserByID).Methods("GET")
	r.HandleFunc("/users", routes.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", routes.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")

	// Middleware CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"}) // frontend Next.js
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	fmt.Println("🚀 Server jalan di http://localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}

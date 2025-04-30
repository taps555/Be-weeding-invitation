package main

import (
	"log"
	"net/http"
	"os"
	"wedding/config"
	"wedding/controller"
	"wedding/repository"
	"wedding/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	config.ConnectDatabase()

	// Inisialisasi repo dengan koneksi DB
	repo := repository.NewRepository(config.DB)

	// Inisialisasi service dengan repo
	allService := service.NewAllService(repo)

	// Inisialisasi controller dengan service
	allController := controller.NewAllController(allService)

	// Routes
	r.HandleFunc("/admin", allController.AddData).Methods("POST")
	r.HandleFunc("/admin/view", allController.GetAll).Methods("GET")
	r.HandleFunc("/undangan/{name}", allController.GetInvitationLink).Methods("POST") // PERBAIKI DI SINI
	r.HandleFunc("/undangan/{name}/myWedding", allController.GetInvitationLink).Methods("GET") // PERBAIKI DI SINI
	r.HandleFunc("/users/{id}", allController.EditData).Methods("PUT")
	r.HandleFunc("/users/{id}", allController.DeleteData).Methods("DELETE")

	// Setup CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://wedding-two-opal.vercel.app", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}

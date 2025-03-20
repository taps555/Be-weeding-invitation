package routes

import (
	"github.com/gorilla/mux"
)

func SetUpRouter() *mux.Router {
	// r := mux.NewRouter()

	// 	// Buat instance service terlebih dahulu
	// 	allService := service.NewAllService() // Pastikan ada constructor ini

	// 	// Baru kemudian buat controller dengan service sebagai parameter
	// 	allController := controller.NewAllController(allService)


	// r.HandleFunc("/admin", allController.AddData).Methods("POST")
	// r.HandleFunc("/admin/view", allController.GetAll).Methods("GET")
	// r.HandleFunc("/users/{id}", allController.EditData).Methods("PUT")
	// r.HandleFunc("/users/{id}", allController.DeleteData).Methods("DELETE")
	// return r
	
	return nil
}
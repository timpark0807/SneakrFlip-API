package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/timpark0807/restapi/handler"
)

func main() {
	router := mux.NewRouter()

	// people
	router.HandleFunc("/api/tenant", handler.ListTenants).Methods("GET")
	router.HandleFunc("/api/tenant/{ss}", handler.GetTenant).Methods("GET")
	router.HandleFunc("/api/tenant", handler.CreateTenant).Methods("POST")
	router.HandleFunc("/api/tenant/{id}", handler.UpdateTenant).Methods("PUT")
	router.HandleFunc("/api/tenant/{ss}", handler.DeleteTenant).Methods("DELETE")

	// property
	router.HandleFunc("/api/property", handler.ListProperties).Methods("GET")
	router.HandleFunc("/api/property", handler.CreateProperty).Methods("POST")
	router.HandleFunc("/api/property/{_id}", handler.GetProperty).Methods("GET")
	router.HandleFunc("/api/property/{_id}", handler.UpdateProperty).Methods("PUT")
	router.HandleFunc("/api/property/{_id}", handler.DeleteProperty).Methods("DELETE")

	router.HandleFunc("/api/item", handler.ListItems).Methods("GET")
	router.HandleFunc("/api/item", handler.CreateItem).Methods("POST")
	router.HandleFunc("/api/item/{_id}", handler.GetItem).Methods("GET")
	// router.HandleFunc("/api/item/{_id}", handler.UpdateItem).Methods("PUT")
	// router.HandleFunc("/api/item/{_id}", handler.DeleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}

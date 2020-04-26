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

	router.HandleFunc("/api/item", handler.ListItems).Methods("GET")
	router.HandleFunc("/api/item", handler.CreateItem).Methods("POST")
	router.HandleFunc("/api/item/{_id}", handler.GetItem).Methods("GET")
	router.HandleFunc("/api/item/updatestatus", handler.UpdateItemStatus).Methods("POST")
	router.HandleFunc("/api/item", handler.UpdateItem).Methods("PUT")
	router.HandleFunc("/api/item/{_id}", handler.DeleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}

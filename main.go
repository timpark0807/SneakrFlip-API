package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/timpark0807/restapi/handler"
	"github.com/timpark0807/restapi/helper"
	"github.com/timpark0807/restapi/model"
	"go.mongodb.org/mongo-driver/bson"
)

func getAllItems(w http.ResponseWriter, req *http.Request) {

	var items []model.Person

	collection := helper.ConnectDB()

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var person model.Person

		err := cur.Decode(&person) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		items = append(items, person)
	}

	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, req *http.Request) {

	var person model.Person
	var params = mux.Vars(req)
	collection := helper.ConnectDB()

	filter := bson.M{"ss": params["ss"]}
	err := collection.FindOne(context.TODO(), filter).Decode(&person)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(person)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var person model.Person

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&person)

	// connect db
	collection := helper.ConnectDB()

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), person)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	// connect db
	collection := helper.ConnectDB()

	filter := bson.M{"ss": params["ss"]}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)
}

func main() {
	router := mux.NewRouter()

	// people
	router.HandleFunc("/api/people", getAllItems).Methods("GET")
	router.HandleFunc("/api/people/{ss}", getItem).Methods("GET")
	// router.HandleFunc("/api/people/{id}", getPerson).Methods("GET")
	router.HandleFunc("/api/people", createPerson).Methods("POST")
	// router.HandleFunc("/api/people/{id}", updatePerson).Methods("PUT")
	router.HandleFunc("/api/people/{ss}", deleteItem).Methods("DELETE")

	// property
	router.HandleFunc("/api/property", handler.ListProperties).Methods("GET")
	router.HandleFunc("/api/property", handler.CreateProperty).Methods("POST")
	router.HandleFunc("/api/property/{id}", handler.GetProperty).Methods("GET")
	router.HandleFunc("/api/property/{id}", handler.UpdateProperty).Methods("PUT")
	router.HandleFunc("/api/property/{id}", handler.DeleteProperty).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

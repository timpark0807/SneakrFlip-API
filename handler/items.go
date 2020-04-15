package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timpark0807/restapi/helper"
	"github.com/timpark0807/restapi/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func ListItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bearerToken, err := helper.CheckToken(r.Header.Get("Authorization"))

	collection := helper.ConnectDB()
	filter := bson.M{"createdby": bearerToken.Email}
	findOptions := options.Find()

	cur, err := collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	var results []*model.Item

	for cur.Next(context.TODO()) {

		var item model.Item
		err := cur.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &item)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	json.NewEncoder(w).Encode(results)

}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bearerToken, err := helper.CheckToken(r.Header.Get("Authorization"))
	var item model.Item
	item.CreatedBy = bearerToken.Email

	if err != nil {
		return
	}

	// decode the post body request
	_ = json.NewDecoder(r.Body).Decode(&item)

	// connect to mongodb
	collection := helper.ConnectDB()

	// insert the new shoe
	result, err := collection.InsertOne(context.TODO(), item)

	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(result)
}

// GetProperty Comment
func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bearerToken, err := helper.CheckToken(r.Header.Get("Authorization"))
	if err != nil {
		return
	}

	var item model.Item
	var params = mux.Vars(r)
	collection := helper.ConnectDB()
	objID, _ := primitive.ObjectIDFromHex(params["_id"])
	filter := bson.M{"_id": objID}
	err = collection.FindOne(context.TODO(), filter).Decode(&item)

	if err != nil {
		return
	}

	if bearerToken.Email != item.CreatedBy {
		return
	}

	json.NewEncoder(w).Encode(item)
}

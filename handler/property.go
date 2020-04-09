package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timpark0807/restapi/helper"
	"github.com/timpark0807/restapi/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// ListProperties Comment
func ListProperties(w http.ResponseWriter, r *http.Request) {
}

// CreateProperty Comment
func CreateProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bearerToken, err := helper.CheckToken(r.Header.Get("Authorization"))

	var property model.Property
	property.CreatedBy = bearerToken.Email

	if err != nil {
		return
	}

	// decode the post body request
	_ = json.NewDecoder(r.Body).Decode(&property)

	// connect to mongodb
	collection := helper.ConnectDB()

	// insert the new property
	result, err := collection.InsertOne(context.TODO(), property)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// GetProperty Comment
func GetProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bearerToken, err := helper.CheckToken(r.Header.Get("Authorization"))
	if err != nil {
		return
	}

	var property model.Property
	var params = mux.Vars(r)
	collection := helper.ConnectDB()
	objID, _ := primitive.ObjectIDFromHex(params["_id"])
	filter := bson.M{"_id": objID}
	err = collection.FindOne(context.TODO(), filter).Decode(&property)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	if bearerToken.Email != property.CreatedBy {
		return
	}

	json.NewEncoder(w).Encode(property)
}

// UpdateProperty Comment
func UpdateProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var property model.Property
	var params = mux.Vars(r)

	collection := helper.ConnectDB()

	filter := bson.M{"_id": params["_id"]}

	_ = json.NewDecoder(r.Body).Decode(&property)

	update := bson.M{
		"$set": bson.M{
			"id":       property.ID,
			"address":  property.Address,
			"zipcode":  property.Zipcode,
			"price":    property.Price,
			"category": property.Category},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&property)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(property)

}

// DeleteProperty Comment
func DeleteProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	// connect db
	collection := helper.ConnectDB()

	filter := bson.M{"_id": params["_id"]}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)
}

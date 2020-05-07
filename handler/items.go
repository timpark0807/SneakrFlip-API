package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/timpark0807/restapi/helper"
	"github.com/timpark0807/restapi/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// ListItems comment
func ListItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bearerToken, err := helper.CheckToken(r.Header.Get("Authorization"))

	if err != nil {
		return
	}

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

// CreateItem comment
func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bearerToken, err := helper.CheckToken(r.Header.Get("Authorization"))
	var item model.Item
	item.CreatedBy = bearerToken.Email
	item.CreatedOn = time.Now().Format("2006-01-02 15:04:05")

	if err != nil {
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&item)
	collection := helper.ConnectDB()
	result, err := collection.InsertOne(context.TODO(), item)

	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(result)
}

// GetItem Comment
func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bearerToken, err := helper.CheckToken(r.Header.Get("Authorization"))
	if err != nil {
		return
	}

	var params = mux.Vars(r)
	item := getItemHelper(params["_id"])

	if bearerToken.Email != item.CreatedBy {
		return
	}

	json.NewEncoder(w).Encode(item)
}

// DeleteItem Comment
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bearerToken, err := helper.CheckToken(r.Header.Get("Authorization"))

	if err != nil {
		return
	}

	var params = mux.Vars(r)

	item := getItemHelper(params["_id"])
	if item.CreatedBy != bearerToken.Email {
		return
	}

	collection := helper.ConnectDB()

	objID, _ := primitive.ObjectIDFromHex(params["_id"])
	filter := bson.M{"_id": objID}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

// UpdateItemStatus Comment
func UpdateItemStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bearerToken, err := helper.CheckToken(r.Header.Get("Authorization"))
	if err != nil {
		return
	}

	var tempItem model.Item
	_ = json.NewDecoder(r.Body).Decode(&tempItem)
	item := getItemHelper(tempItem.ID.Hex())

	if item.CreatedBy != bearerToken.Email {
		return
	}

	filter := bson.M{"_id": item.ID}
	collection := helper.ConnectDB()
	update := bson.M{
		"$set": bson.M{
			"sold":      !item.Sold,
			"updatedon": time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	item.UpdateSoldStatus()
	json.NewEncoder(w).Encode(result)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item model.Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	bearerToken, err := helper.CheckToken(r.Header.Get("Authorization"))

	if item.CreatedBy != bearerToken.Email || err != nil {
		return
	}

	collection := helper.ConnectDB()
	filter := bson.M{"_id": item.ID}
	update := bson.M{
		"$set": bson.M{
			"_id":         item.ID,
			"category":    item.Category,
			"brand":       item.Brand,
			"description": item.Description,
			"size":        item.Size,
			"condition":   item.Condition,
			"sold":        item.Sold,
			"createdby":   item.CreatedBy,
			"updatedon":   time.Now().Format("2006-01-02 15:04:05")},
	}

	err = collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&item)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(item)

}

func getItemHelper(_id string) model.Item {
	objID, _ := primitive.ObjectIDFromHex(_id)
	filter := bson.M{"_id": objID}
	var item model.Item
	collection := helper.ConnectDB()
	collection.FindOne(context.TODO(), filter).Decode(&item)
	return item
}

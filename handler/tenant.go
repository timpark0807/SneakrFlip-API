package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timpark0807/restapi/helper"
	"github.com/timpark0807/restapi/model"
	"gopkg.in/mgo.v2/bson"
)

// ListTenants comment
func ListTenants(w http.ResponseWriter, req *http.Request) {

	var tenants []model.Tenant

	collection := helper.ConnectDB()

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var tenant model.Tenant

		err := cur.Decode(&tenant) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		tenants = append(tenants, tenant)
	}

	json.NewEncoder(w).Encode(tenants)
}

// GetTenant comment
func GetTenant(w http.ResponseWriter, req *http.Request) {

	var tenant model.Tenant
	var params = mux.Vars(req)
	collection := helper.ConnectDB()

	filter := bson.M{"ss": params["ss"]}
	err := collection.FindOne(context.TODO(), filter).Decode(&tenant)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(tenant)
}

// CreateTenant comment
func CreateTenant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tenant model.Tenant

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&tenant)

	// connect db
	collection := helper.ConnectDB()

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), tenant)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// UpdateTenant comment
func UpdateTenant(w http.ResponseWriter, r *http.Request) {
}

// DeleteTenant comment
func DeleteTenant(w http.ResponseWriter, r *http.Request) {
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

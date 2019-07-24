package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/YAWAL/tcbe/src/routers"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/YAWAL/tcbe/src/db"
	"github.com/YAWAL/tcbe/src/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	timeout = 30 * time.Second
	nameKey = "name"
	idKey   = "id"
)

func setHeader(resp http.ResponseWriter) {
	resp.Header().Set("content-type", "application/json")
}

func errorMessage(resp http.ResponseWriter, err error) {
	resp.WriteHeader(http.StatusInternalServerError)
	_, err = resp.Write([]byte(`{ "message": "` + err.Error() + `" }`))
}

func Init(conf *db.MongoConfig) map[string]func(resp http.ResponseWriter, req *http.Request) {
	var handlers = map[string]func(resp http.ResponseWriter, req *http.Request){
		// One service to return patient information
		routers.PatientRoute: func(resp http.ResponseWriter, req *http.Request) {
			setHeader(resp)
			params := mux.Vars(req)
			id, err := primitive.ObjectIDFromHex(params[idKey])
			if err != nil {
				errorMessage(resp, err)
				return
			}
			var patient models.Patient
			collection := db.Client.Database(conf.DatabaseName).Collection(conf.PatientCollectionName)
			ctx, _ := context.WithTimeout(context.Background(), timeout)
			err = collection.FindOne(ctx, models.Patient{ID: id}).Decode(&patient)
			if err != nil {
				errorMessage(resp, err)
				return
			}
			err = json.NewEncoder(resp).Encode(patient)
			if err != nil {
				errorMessage(resp, err)
				return
			}
		},
		//One service to return illness information
		routers.IllnessRoute: func(resp http.ResponseWriter, req *http.Request) {
			setHeader(resp)
			params := mux.Vars(req)
			name, ok := params[nameKey]
			if !ok {
				errorMessage(resp, errors.New("cannot find proper parameter"))
				return
			}
			var illness models.Illness
			collection := db.Client.Database(conf.DatabaseName).Collection(conf.IllnessCollectionName)
			ctx, _ := context.WithTimeout(context.Background(), timeout)
			err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&illness)
			if err != nil {
				errorMessage(resp, err)
				return
			}
			err = json.NewEncoder(resp).Encode(illness)
			if err != nil {
				errorMessage(resp, err)
				return
			}
		},

	}

	return handlers
}

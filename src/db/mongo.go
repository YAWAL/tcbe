package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

type MongoConfig struct {
	ApplyURI              string `json:"apply_uri"`
	DatabaseName          string `json:"database_name"`
	PatientCollectionName string `json:"patient_collection_name"`
	IllnessCollectionName string `json:"illness_collection_name"`
}

func Mongo(conf *MongoConfig) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(conf.ApplyURI)
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	return nil
}

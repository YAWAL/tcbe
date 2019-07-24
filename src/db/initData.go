package db

import (
	"context"
	"time"

	"github.com/YAWAL/tcbe/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FillInitData(conf *MongoConfig) (err error) {
	patients := []models.Patient{
		{
			ID:        primitive.ObjectID{},
			Firstname: "john",
			Lastname:  "doe",
		},
		{
			ID:        primitive.ObjectID{},
			Firstname: "ivan",
			Lastname:  "franko",
		},
		{
			ID:        primitive.ObjectID{},
			Firstname: "taras",
			Lastname:  "kost",
		},
		{
			ID:        primitive.ObjectID{},
			Firstname: "hyj",
			Lastname:  "huj",
		},
		{
			ID:        primitive.ObjectID{},
			Firstname: "arin",
			Lastname:  "roberts",
		},
	}
	illnesses := []models.Illness{
		{
			ID:                          primitive.ObjectID{},
			Name:                        "firstilln",
			TemperatureInDegreesCelsius: "38",
			CoughPresence:               false,
			CoughType:                   "none",
			HeartRate:                   "80",
		},
		{
			ID:                          primitive.ObjectID{},
			Name:                        "secondilln",
			TemperatureInDegreesCelsius: "40",
			CoughPresence:               true,
			CoughType:                   "CoughType2",
			HeartRate:                   "90",
		},
		{
			ID:                          primitive.ObjectID{},
			Name:                        "fifthilln",
			TemperatureInDegreesCelsius: "39",
			CoughPresence:               false,
			CoughType:                   "none",
			HeartRate:                   "100",
		},
	}
	collection := Client.Database(conf.DatabaseName).Collection(conf.PatientCollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	for _, p := range patients {
		_, err := collection.InsertOne(ctx, p)
		if err != nil {
			return err
		}
	}
	collection = Client.Database(conf.DatabaseName).Collection(conf.IllnessCollectionName)
	for _, i := range illnesses {
		_, err := collection.InsertOne(ctx, i)
		if err != nil {
			return err
		}
	}
	return nil
}

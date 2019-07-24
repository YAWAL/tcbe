package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Patient contains information about individual patients
type Patient struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

// Illness contains information about illnesses
type Illness struct {
	ID                          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                        string             `json:"name,omitempty" bson:"name,omitempty"`
	TemperatureInDegreesCelsius string             `json:"temperature,omitempty" bson:"temperature,omitempty"`
	CoughPresence               bool               `json:"cough_presence,omitempty" bson:"cough_presence,omitempty"`
	CoughType                   string             `json:"cough_type,omitempty" bson:"cough_type,omitempty"`
	HeartRate                   string             `json:"heart_rate,omitempty" bson:"heart_rate,omitempty"`
}

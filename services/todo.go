package services

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// type struct that holds data for todo - specifying json n bson {for mongodb data} in correct format
type Todo struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool `json:"completed,omitempty" bson:"completed,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

var client *mongo.Client

//  returns instance data of type Todo -> any method that belongs to type Todo or instance will be accessible through this
func NewMongoClient(mongo *mongo.Client) Todo {
	client = mongo
	return Todo{}
}



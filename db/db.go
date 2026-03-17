package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type data that stores mongo db collections
var collection *mongo.Collection

func ConnectToMongo() (*mongo.Client,error) {

	// loading and using env var from env file
	godotenv.Load()
	username := os.Getenv("MONGO_DB_USERNAME") 
	password := os.Getenv("MONGO_DB_PASSWORD") 
	
	// MongoDB connection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Configuring credentials for mongoDB connection
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	// finally connecting to mongo
	client,err := mongo.Connect(context.Background(),clientOptions)
	
	// if caught erro while connecting to the mongo db 
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	log.Println("successfully connected to MongoDB🚀")
	return client,err
}

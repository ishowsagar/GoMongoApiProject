package main

import (
	"context"
	_ "fmt"
	"log"
	"net/http"
	"server/db"
	"server/handlers"
	"server/services"
	"time"
)

// types
type Application struct {
	// JInCase := if there are methods that belongs to Todo, so it means as they belongs to Todo, we can access through Todo through direct type or even instance
	Models services.Models //@ type that stores Todo
}

func main() {
	// todo ~ connection to mongoDB using db package func
	mongoClient,err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}

	//& func only proceeds if satisfies this --> if like connection takes more than set timeout it won't wait more & exit fnc early
	context,cancel := context.WithTimeout(context.Background(), 7 * time.Second)
	defer cancel()
	// if at the end there is some disconnection error --> stop func
	defer func ()  {
		if err = mongoClient.Disconnect(context); err != nil {
			panic(err)
		}	
	}()

	//* if successfully connected to db
	services.NewMongoClient(mongoClient)
	log.Println("Server is running on port",8080)
	// routing request to this Listen N serve http method which route to this handler
	log.Fatal(http.ListenAndServe(":8080",handlers.CreateRouter())) // ! listen & serve on this port addr and handled {pass handler} by this chi router func

} 

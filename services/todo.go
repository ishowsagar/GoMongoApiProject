package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// type struct that holds data for todo - specifying json n bson {for mongodb data} in correct format
type Todo struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool `json:"completed" bson:"completed"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

var client *mongo.Client

//  returns instance data of type Todo -> any method that belongs to type Todo or instance will be accessible through this
func NewMongoClient(mongo *mongo.Client) Todo {
	client = mongo
	return Todo{}
}

// @ These all methods using recievers belong to type Todo --> accessible by native type or its instance
//! method that insert todo into the db
func (t *Todo) InsertTodo(entry Todo) error {
	//  as we would not be returning anything but making a db call and insertion
	
	// # step 1 --> retrieve the access to the collection
	collection := returnCollectionPointer("todos") //&returns collection (think like table "todos" from database "that")

	// # step 2 --> insert into the collection
	_,err := collection.InsertOne(context.TODO(),Todo{
		Task: entry.Task,
		Completed: entry.Completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Println("Error :",err)
	}
	return nil
}

// ! method that belongs to type Todo &  Get todo by its id -- /api/todos/{id}
func (t *Todo) GetTodoByID(id string) (Todo,error) {
	collection := returnCollectionPointer("todos")
	var todo Todo

	// since mongo id is formatted in struct of object --> have to convert into that
	mongoID,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error :",err)
		return Todo{},err
	}

	// & finding that one entry with this method passing bg,bson.M{key:SetToThisVarValue} --> what are we looking for
	err = collection.FindOne(context.Background(),bson.M{"_id" : mongoID}).Decode(&todo)
	if err != nil {
		log.Println("Error :",err)
		return Todo{},err
	}

	return todo,nil 
}

// ! method that belongs to type Todo &  update a todo -- /api/todos/
func (t *Todo) UpdateTodo(id string,entry Todo) (*mongo.UpdateResult ,error) {
	collection := returnCollectionPointer("todos")
	mongoID,err := primitive.ObjectIDFromHex(id) //* parsed from mongoObject id into valid id
	if err != nil {
		return nil,err
	}

	update := bson.D{
		{"$set", bson.D{
			{"task", entry.Task},
			{"completed", entry.Completed},
			{"updated_at", time.Now()},
		}},
	}

	res,err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id":mongoID},
		update,
	)

	if err != nil {
		return nil,err
	}

	return res,nil
}
// ! method that belongs to type Todo &  Delete todos -- /api/todos/
func(t *Todo) DeleteTodo(id string) error {
	// get collection from client(mongo db client connection)
	collection := returnCollectionPointer("todos") 
	// parse id from mongoStructId using prim.ObjectIdFromHec method passing id string
	mongoID,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error :",err)
		return err
	}

	// delete entry from collection by using delete method
	_,err = collection.DeleteOne(context.Background(),bson.M{"_id": mongoID})
	if err != nil {
		log.Println("Error :",err)
		return err
	}
	return nil
}


// ! method that belongs to type Todo &  Get all todos -- /api/todos/all
func (t *Todo) GetAllTodos() ([]Todo,error) {
	// as this would return array of Todo type data and ofc error handeling
	var todos_holder_slice []Todo
	
	collection := returnCollectionPointer("todos")
	
	cursor,err := collection.Find(context.TODO(),bson.D{})
	if err != nil {
		log.Println("Error :",err)
		return nil,err
	}
	defer cursor.Close(context.Background())

	// Scan each field and values into type that holds all todos in slice
	for cursor.Next(context.Background()) {
		// # for each collection's entry , decode data n inject into todo and append todo to ealier Slice
		var todo Todo
		cursor.Decode(&todo) //* decoding into where --> memory addr of this todo type
		todos_holder_slice = append(todos_holder_slice, todo) 
	}
	return todos_holder_slice,nil
}


// returns collection for passed string - Database is like database model and collection set of data stored in them
func returnCollectionPointer(collection string) *mongo.Collection {
	// & client is like mongo library connection, have databases and databases have collection like tables
	return client.Database("todos_db").Collection(collection)
}



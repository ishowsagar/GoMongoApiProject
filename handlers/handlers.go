package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/services"

	"github.com/go-chi/chi"
)

var todo_TypeStruct services.Todo

// health check func for checking if server is running with no data, just void server run check❕❕
func  HealthCheck(w http.ResponseWriter,r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_,err := w.Write([]byte("Go server is running 🧑‍💻⚡..."))
	
	if err != nil {
		http.Error(w,"failed to boot Go Server🛑🛑",http.StatusBadRequest)
		return
	}
}

// route - todos/all
func GetAllTodos(w http.ResponseWriter,r *http.Request) {
	todos,err := todo_TypeStruct.GetAllTodos()
	if err != nil {
		http.Error(w,"failed to fetch todos.",http.StatusBadRequest)
		return
	}

	// * Send back response to client with todos n headers
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
	// ! since this is struct data, w.Write does not know about this --> have to use jsonEncoder
	json.NewEncoder(w).Encode(todos)
}

// route - todos/{id}
func GetTodoByID(w http.ResponseWriter,r *http.Request) {
	
	id := chi.URLParam(r,"id")
	todo,err := todo_TypeStruct.GetTodoByID(id)
	if err != nil {
		http.Error(w,"Failed to fetch todo due to incorrect ID",http.StatusBadRequest)
		return
	}
	
	// * Send back response to client with todos n headers
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
	
}

// route - todos/create
func CreateTodo(w http.ResponseWriter,r *http.Request) {
	// todo --> Decode incoming r.Body and inject data into type that stores todo type Data
	var recieved_todo_payload services.Todo
	err := json.NewDecoder(r.Body).Decode(&recieved_todo_payload) //* since we need to interpret actual value so inject into addr of type storing incoming r.Body data

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	//if  successfully retreived data in recieved_todo_payload, insert into db by making call using func
	err = todo_TypeStruct.InsertTodo(recieved_todo_payload) //! since all method belongs to type Todo, we are accessing methods from it
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Failed to create todo due to unrecognized data")
		return
	}

	jsonStr,err := json.Marshal("successfully created todo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}

	// * sending response back to client with msg n headers
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStr)
}

// route - todos/update
func UpdateTodo(w http.ResponseWriter,r *http.Request) {
	id := chi.URLParam(r,"id")
	var todo_payload services.Todo
	err :=json.NewDecoder(r.Body).Decode(&todo_payload)
	if err != nil {
		http.Error(w,"Failed to fetch todo due to incorrect data passed to update",http.StatusBadRequest)
		return
	}
	updated_todo,err := todo_TypeStruct.UpdateTodo(id,todo_payload)
	// * sending response back to client with msg n headers
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated_todo) //! when we use json.NewEncoder it encodes data to reponse Writer ==> sends data

}

func DeleteTodo(w http.ResponseWriter,r *http.Request) {
	id_param := chi.URLParam(r,"id")
	err := todo_TypeStruct.DeleteTodo(id_param)
	if err != nil {
		http.Error(w,"Failed to delete todo due to incorrect ID",http.StatusBadRequest)
		return
	}
	jsonStr,err := json.Marshal("successfully deleted todo.")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}
	// * sending response back to client with msg n headers
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(jsonStr)	
}

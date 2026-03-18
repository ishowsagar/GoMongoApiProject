package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// you could either use chi return type or native http.Handler --> both are fine as unde the hood chi uses that
func CreateRouter() http.Handler {
	// ! initializing router via chi method 
	router := chi.NewRouter()
	// setting cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// @ Routes & their respective functions to serve response back to the client 
	
// health check func for checking if server is running with no data, just void server run check❕❕ 
router.Get("/health",HealthCheck)

//@ nested routing with parent route paths
router.Route("/api",func(r chi.Router) {
	// nested routes under "/api" route path and use "r" as nested router
	r.Get("/health",HealthCheck)
	r.Get("/todos/all",GetAllTodos)
	r.Get("/todos/{id}",GetTodoByID)
	r.Post("/todos/create",CreateTodo)
	r.Put("/todos/update/{id}",UpdateTodo)
	r.Delete("/todos/delete/{id}",DeleteTodo)

	// r.Get("/todos/all",)
})

	return router
}
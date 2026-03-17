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
	router.Get("/health",func(w http.ResponseWriter,r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_,err := w.Write([]byte("Go server is running🧑‍💻⚡..."))
		
		if err != nil {
			http.Error(w,"failed to boot Go Server🛑🛑",http.StatusBadRequest)
		}
	})

	return router
}
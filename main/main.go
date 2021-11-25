package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo-app/router"
)

func main() {
	r := mux.NewRouter()
	router.RegisterTodoRoutes(r)
	http.Handle("/", r)
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE","HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	log.Fatal(http.ListenAndServe("localhost:3000", handlers.CORS(headers, credentials, methods, origins)(r)))
}

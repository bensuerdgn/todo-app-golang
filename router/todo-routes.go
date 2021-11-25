package router

import (
	"github.com/gorilla/mux"
	"todo-app/controller"
)

var RegisterTodoRoutes = func(router *mux.Router) {
	router.HandleFunc("/createTodo/", controller.CreateTodo).Methods("POST")
	router.HandleFunc("/getTodo/", controller.GetTodo).Methods("GET")
	router.HandleFunc("/getTodoById/{todoId}", controller.GetTodoById).Methods("GET")
	router.HandleFunc("/updateTodo/{todoId}", controller.UpdateTodo).Methods("PUT")
	router.HandleFunc("/deleteTodo/{todoId}", controller.DeleteTodo).Methods("Delete")
}


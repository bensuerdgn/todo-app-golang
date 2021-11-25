package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"todo-app/model"
	"todo-app/utils"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	CreateTodo := &model.Todo{}
	utils.ParseBody(r, CreateTodo)
	b := CreateTodo.CreateTodo()
	res, _ := json.Marshal(b)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetTodo(w http.ResponseWriter, r *http.Request) {
	newTodos := model.GetAllTodos()
	res, _ := json.Marshal(newTodos)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("error parsing")
	}
	todoDetails, _ := model.GetTodoById(ID)
	res, _ := json.Marshal(todoDetails)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var updateTodo = &model.Todo{}
	utils.ParseBody(r, updateTodo)

	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("parsing error")
	}
	todoDetails, db := model.GetTodoById(ID)
	if updateTodo.Name != "" {
		todoDetails.Name = updateTodo.Name
	}
	if updateTodo.Date != "" {
		todoDetails.Date = updateTodo.Date
	}
	if updateTodo.Priority != "" {
		todoDetails.Priority = updateTodo.Priority
	}
	if updateTodo.Status != "" {
		todoDetails.Status = updateTodo.Status
	}
	db.Save(&todoDetails)
	res, _ := json.Marshal(todoDetails)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("parsing error")
	}
	todo := model.DeleteTodo(ID)
	res, _ := json.Marshal(todo)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

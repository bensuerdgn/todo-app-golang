package model

import (
	"gorm.io/gorm"
	"todo-app/config"
)

var db *gorm.DB

type Todo struct{
	gorm.Model
	Name string `gorm:"" json:"Name"`
	Date string `json:"Date"`
	Priority string `json:"Priority"`
	Status string `json:"Status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
}

func (t *Todo) CreateTodo() *Todo {
	db.Create(&t)
	return t
}

func GetAllTodos() []Todo {
	var Todos []Todo
	db.Find(&Todos)
	return Todos
}

func GetTodoById(Id int64) (*Todo, *gorm.DB)  {
	var getTodo Todo
	db := db.Where("ID=?", Id).Find(&getTodo)
	return &getTodo, db
}

func DeleteTodo(ID int64) Todo {
	var todo Todo
	db.Where("ID=?", ID).Delete(&todo)
	return todo
}
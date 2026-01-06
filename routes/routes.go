package routes

import (
	"net/http"

	"golang-todo-api/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/todos", handlers.GetTodos)
	http.HandleFunc("/todo/create", handlers.CreateTodo)
	http.HandleFunc("/todo/delete", handlers.DeleteTodo)
	http.HandleFunc("/todo", handlers.GetTodoByID)
	http.HandleFunc("/todo/update", handlers.UpdateTodo)
}
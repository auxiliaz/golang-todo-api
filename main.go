package main

import (
	"log"
	"net/http"

	"golang-todo-api/config"
	"golang-todo-api/handlers"
	"golang-todo-api/routes"
)

func main() {
	db := config.ConnectDB()
	handlers.DB = db

	routes.RegisterRoutes()

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
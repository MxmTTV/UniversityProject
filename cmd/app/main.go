package main

import (
	"github.com/gorilla/mux"
	"go.mod/internal/database"
	"go.mod/internal/handlers"
	"go.mod/internal/taskService"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetHandlerGetTask).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", handler.PutTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

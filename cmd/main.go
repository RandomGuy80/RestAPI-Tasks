package main

import (
	"RestAPITasks/internal/delivery"
	"RestAPITasks/internal/repository"
	"RestAPITasks/internal/storage"
	"RestAPITasks/internal/usecase"
	"net/http"
)

func main() {
	db := storage.NewStorage()

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := delivery.NewHandler(userUsecase)

	r := http.NewServeMux()
	r.HandleFunc("POST /tasks", userHandler.SaveTask)
	r.HandleFunc("GET /tasks/{id}", userHandler.GetTaskById)

	http.ListenAndServe(":8080", r)
}

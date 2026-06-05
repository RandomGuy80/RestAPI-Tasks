package main

import (
	"RestAPITasks/internal/delivery"
	"RestAPITasks/internal/repository"
	"RestAPITasks/internal/storage"
	"RestAPITasks/internal/usecase"
	"log"
	"net/http"
)

func main() {
	db, err := storage.NewStorage("postgres://postgres:password@localhost:5432/tasks_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := delivery.NewHandler(userUsecase)

	r := http.NewServeMux()
	r.HandleFunc("POST /tasks", userHandler.SaveTask)
	r.HandleFunc("GET /tasks/{id}", userHandler.GetTaskById)
	r.HandleFunc("DELETE /tasks/{id}", userHandler.DeleteTaskById)
	r.HandleFunc("PUT /tasks/{id}", userHandler.UpdateById)

	http.ListenAndServe(":8080", r)
}

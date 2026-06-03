package delivery

import (
	"RestAPITasks/internal/domain"
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	usecase domain.UserUsecase
}

func NewHandler(usecase domain.UserUsecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) SaveTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if err := h.usecase.Save(task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	task, err := h.usecase.GetTaskById(id)
	if err != nil {
		http.Error(w, "task undefined", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

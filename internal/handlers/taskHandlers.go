package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mod/internal/taskService" // Импортируем наш сервис
	"net/http"
	"strconv"
)

type Handler struct {
	Service *taskService.TaskService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetHandlerGetTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

func (h *Handler) PutTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	taskID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Неверный идентификатор задачи", http.StatusBadRequest)
		return
	}

	var updateTask taskService.Task
	err = json.NewDecoder(r.Body).Decode(&updateTask)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	// Обновляем задачу в сервисе
	updatedTask, err := h.Service.PatchTask(uint(taskID), updateTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Устанавливаем заголовки ответа
	w.Header().Set("Content-Type", "application/json")
	// Кодируем обновленную задачу в JSON и отправляем ответ
	json.NewEncoder(w).Encode(updatedTask)
}

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	taskID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.Service.DeleteTask(uint(taskID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Отправляем пустой ответ
	w.WriteHeader(http.StatusNoContent)
}

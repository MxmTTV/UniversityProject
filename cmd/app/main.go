package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func PostHandlerPostTask(w http.ResponseWriter, r *http.Request) {
	// читаю тело запроса
	var task Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}
	DB.Create(&task)
	// Ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetHandlerGetTask(w http.ResponseWriter, r *http.Request) {
	var tasks []Message
	DB.Find(&tasks)

	// Ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func PatchHandlerPutTask(w http.ResponseWriter, r *http.Request) {
	// получить айди из урла, так как строка, сделать конвертацию
	params := mux.Vars(r)
	taskID := params["id"]

	// конвертация строки в число
	id, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}
	// нахождение задачи по айдишнику
	var task Message
	result := DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, "Айди не найден", http.StatusBadRequest)
		return
	}

	// меняю задачи task и is_done
	var updateTask Message
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updateTask)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	task.Task = updateTask.Task
	task.IsDone = updateTask.IsDone
	// Сохраняю в БД новые значения
	DB.Save(&task)
	// Отправляю ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteHandlerDeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID := params["id"]

	// конвертация
	id, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "Ошибка кон ID", http.StatusBadRequest)
		return
	}

	var task Message
	result := DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, "Запись не найдена", http.StatusBadRequest)
		return
	}
	if err := DB.Delete(&task).Error; err != nil {
		http.Error(w, "Запись не удалена", http.StatusBadRequest)
		return
	}
	// ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()
	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/post", PostHandlerPostTask).Methods("POST")
	router.HandleFunc("/get", GetHandlerGetTask).Methods("GET")
	router.HandleFunc("/patch/{id}", PatchHandlerPutTask).Methods("PATCH")
	router.HandleFunc("/delete/{id}", DeleteHandlerDeleteTask).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

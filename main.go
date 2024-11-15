package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Глобальная переменная для хранения сообщения

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/post", handlerPost).Methods("POST")
	router.HandleFunc("/get", handlerGet).Methods("GET")

	// посылаю ПОСТ-запрос
	//http.Post("http://localhost:8080/post", "application/json", strings.NewReader(`{
	//"task": "Тестовая задача",
	//"is_done": false}`))

	http.ListenAndServe(":8080", router)

}

func handlerPost(w http.ResponseWriter, r *http.Request) {
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

func handlerGet(w http.ResponseWriter, r *http.Request) {
	var tasks []Message
	DB.Find(&tasks)

	// Ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

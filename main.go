package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/post", PostHandlerPostTask).Methods("POST")
	router.HandleFunc("/get", GetHandlerGetTask).Methods("GET")

	// посылаю ПОСТ-запрос
	//http.Post("http://localhost:8080/post", "application/json", strings.NewReader(`{
	//"task": "Тестовая задача",
	//"is_done": false}`))

	http.ListenAndServe(":8080", router)

}

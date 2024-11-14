package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Структура для декодирования JSON
type requestBody struct {
	Message string `json:"message"`
}

// Глобальная переменная для хранения сообщения
var task string

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/get", handlerGet).Methods("GET")
	router.HandleFunc("/api/hello", handlerPost).Methods("POST")

	// посылаю ПОСТ-запрос
	http.Post("http://localhost:8080/api/hello", "application/json", strings.NewReader(`{"message": "World"}`))

	http.ListenAndServe(":8080", router)

}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	// читаю тело запроса
	var body requestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	// присвоил значение глобальной переменной
	task = body.Message

}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	// возвращаю текущее значение глобальной переменной
	fmt.Fprintf(w, "Hello, %s\n", task)
}

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mod/internal/database"
	"go.mod/internal/handlers"
	"go.mod/internal/models"
	"go.mod/internal/taskService"
	"go.mod/internal/userService"
	"go.mod/internal/web/tasks"
	"go.mod/internal/web/users"
	"log"
	"net/http"
)

func main() {
	database.InitDB()

	if err := database.DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	// Создание репозитория для задач
	taskRepo := taskService.NewTaskRepository(database.DB)
	// Создание сервиса для задач, передаем taskRepo
	taskService := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Создание репозитория для пользователей
	userRepo := userService.NewUserRepository(database.DB)
	// Создание сервиса для пользователей, теперь передаем taskRepo, чтобы сервис мог работать с задачами
	userService := userService.NewUserService(userRepo, taskRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Инициализируем echo
	e := echo.New()

	// Используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
	}))

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTaskHandler := tasks.NewStrictHandler(taskHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictTaskHandler)

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictUserHandler := users.NewStrictHandler(userHandler, nil) // тут будет ошибка
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}

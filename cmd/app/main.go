package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mod/internal/database"
	"go.mod/internal/handlers"
	"go.mod/internal/taskService"
	"go.mod/internal/userService"
	"go.mod/internal/web/tasks"
	"go.mod/internal/web/users"
	"log"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	if err := database.DB.AutoMigrate(&userService.User{}); err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

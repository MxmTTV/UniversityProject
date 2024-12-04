package handlers

import (
	"fmt"
	"go.mod/internal/models"
	"go.mod/internal/taskService" // Импортируем наш сервис
	"go.mod/internal/web/tasks"
	"golang.org/x/net/context"
	"strconv"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) GetTasksUserId(ctx context.Context, request tasks.GetTasksUserIdRequestObject) (tasks.GetTasksUserIdResponseObject, error) {
	// Здесь извлекаем ID пользователя из запроса
	userID := request.Id

	// Получаем задачи для указанного пользователя
	_, err := h.Service.GetTasksByUserID(uint(userID))
	if err != nil {
		return nil, err
	}

	// Создаем ответ
	response := tasks.GetTasksUserId200JSONResponse{}

	return response, nil
}

// Метод для получения всех задач
func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}
	return response, nil
}

// Метод для получения задач по user_id
func (h *Handler) GetTasksForUser(ctx context.Context, request tasks.GetTasksUserIdRequestObject) (tasks.GetTasksUserIdResponseObject, error) {
	userID, err := strconv.Atoi(strconv.Itoa(int(request.Id))) // Получаем user_id из запроса
	if err != nil {
		return nil, err
	}

	// Получаем задачи для конкретного пользователя
	tasksForUser, err := h.Service.GetTasksByUserID(uint(userID))
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasksUserId200JSONResponse{}
	for _, task := range tasksForUser {
		taskResponse := tasks.Task{
			Id:     &task.ID,
			Task:   &task.Task,
			IsDone: &task.IsDone,
		}
		response = append(response, taskResponse)
	}
	return response, nil
}

// Метод для создания задачи
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	// Проверяем, что UserId не nil
	if taskRequest.UserId == nil {
		return nil, fmt.Errorf("UserId is required")
	}

	// Создаем задачу с привязкой к пользователю
	taskToCreate := models.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId, // Привязываем задачу к пользователю
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		UserId: &createdTask.UserID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

// Метод для обновления задачи
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id
	updateRequest := request.Body

	toUpdate := models.Task{
		Task:   *updateRequest.Task,
		IsDone: *updateRequest.IsDone,
	}

	updateTask, err := h.Service.PatchTask(id, toUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updateTask.ID,
		Task:   &updateTask.Task,
		IsDone: &updateTask.IsDone,
	}

	return response, nil
}

// Метод для удаления задачи
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id

	err := h.Service.DeleteTask(taskID)
	if err != nil {
		return nil, err
	}

	response := tasks.DeleteTasksId200Response{}
	return response, nil
}

// Конструктор для создания Handler
func NewTaskHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

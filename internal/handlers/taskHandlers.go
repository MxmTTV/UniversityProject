package handlers

import (
	"go.mod/internal/taskService" // Импортируем наш сервис
	"go.mod/internal/web/tasks"
	"golang.org/x/net/context"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// получаем все задачи из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	// создаём переменную, которая будет давать ответ(респонс) типа 200джейсонРеспонс
	// которая будет передана дальше в качества ответа
	response := tasks.GetTasks200JSONResponse{}
	// заполнение слайса всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// читаем тело напрямую!!!
	taskRequest := request.Body
	// обращаемся к сервису и создаём задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}
	// создаём структуру респонс, чтобы показать созданную задачу
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Получаем ID из запроса
	id := request.Id

	// Получаем тело запроса для обновления
	updateRequest := request.Body

	// структура для обновления
	toUpdate := taskService.Task{
		Task:   *updateRequest.Task,
		IsDone: *updateRequest.IsDone,
	}

	// метод для обновления задачи
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

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Получаем ID задачи из запроса
	taskID := request.Id

	// Вызываем сервис для удаления задачи по ID
	err := h.Service.DeleteTask(taskID)
	if err != nil {
		return nil, err
	}
	response := tasks.DeleteTasksId200Response{}
	return response, nil
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewTaskHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

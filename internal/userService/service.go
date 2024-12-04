package userService

import (
	"go.mod/internal/models"
	"go.mod/internal/taskService"
)

type UserService struct {
	repo     UserRepository
	taskRepo taskService.TaskRepository
}

func NewUserService(repo UserRepository, taskRepo taskService.TaskRepository) *UserService {
	return &UserService{repo: repo, taskRepo: taskRepo}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateTask(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllTasks()
}

func (s *UserService) PatchUser(id uint, updates models.User) (models.User, error) {
	return s.repo.UpdateTaskByID(id, updates)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

func (s *UserService) GetTasksForUser(id uint) ([]models.Task, error) {
	return s.taskRepo.GetTasksByUserID(id) // Вызов метода из taskRepo
}

package taskService

import (
	"go.mod/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task models.Task) (models.Task, error)
	// GetAllTasks - Возвращаем массив из models.Task задач в БД и ошибку
	GetAllTasks() ([]models.Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task models.Task) (models.Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
	// GetTasksByUserID GetTasksByID Получение всех задач определенного пользователя
	GetTasksByUserID(userID uint) ([]models.Task, error)
}
type taskRepository struct {
	db *gorm.DB
}

func (r *taskRepository) UpdateTaskByID(id uint, task models.Task) (models.Task, error) {
	err := r.db.Model(&models.Task{}).Where("id = ?", id).Updates(task).Error
	return task, err
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию

func (r *taskRepository) CreateTask(task models.Task) (models.Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	err := r.db.Where("id = ?", id).Delete(&models.Task{}).Error
	return err
}
func (r *taskRepository) GetTasksByUserID(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

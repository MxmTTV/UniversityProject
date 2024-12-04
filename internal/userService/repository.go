package userService

import (
	"go.mod/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(user models.User) (models.User, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]models.User, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, user models.User) (models.User, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
}
type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) UpdateTaskByID(id uint, user models.User) (models.User, error) {
	err := r.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
	return user, err
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию

func (r *userRepository) CreateTask(user models.User) (models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllTasks() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) DeleteTaskByID(id uint) error {
	err := r.db.Where("id = ?", id).Delete(&models.User{}).Error
	return err
}

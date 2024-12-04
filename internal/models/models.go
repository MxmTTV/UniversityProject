package models

type Task struct {
	ID     uint   `gorm:"primaryKey"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"` // Foreign key
	User   User   `json:"user"`    // Определение связи с моделью User
}

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`    // Добавляем поле email
	Password string `json:"password"` // Добавляем поле password
}

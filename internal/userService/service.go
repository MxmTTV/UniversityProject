package userService

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateTask(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllTasks()
}

func (s *UserService) PatchUser(id uint, updates User) (User, error) {
	return s.repo.UpdateTaskByID(id, updates)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

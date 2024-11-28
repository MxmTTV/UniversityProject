package handlers

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
	"go.mod/internal/userService"
	"go.mod/internal/web/users"
	"golang.org/x/net/context"
)

type UserHandler struct {
	Service *userService.UserService
}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    (*openapi_types.Email)(&usr.Email),
			Password: &usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, request users.CreateUserRequestObject) (users.CreateUserResponseObject, error) {
	// читаем тело напрямую!!!
	userRequest := request.Body
	// обращаемся к сервису и создаём задачу
	userCreate := userService.User{
		Email:    string(*userRequest.Email),
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userCreate)
	if err != nil {
		return nil, err
	}
	// создаём структуру респонс, чтобы показать созданную задачу
	response := users.CreateUser201JSONResponse{
		Id:       &createdUser.ID,
		Email:    (*openapi_types.Email)(&createdUser.Email),
		Password: &createdUser.Password,
	}
	return response, nil
}

func (h *UserHandler) DeleteUserByID(ctx context.Context, request users.DeleteUserByIDRequestObject) (users.DeleteUserByIDResponseObject, error) {
	userId := request.Id

	err := h.Service.DeleteUser(uint(userId))
	if err != nil {
		return nil, err
	}

	response := users.DeleteUserByID204Response{}
	return response, nil
}

func (h *UserHandler) UpdateUserByID(ctx context.Context, request users.UpdateUserByIDRequestObject) (users.UpdateUserByIDResponseObject, error) {
	userId := request.Id

	userUpdate := request.Body

	toUpdateUser := userService.User{
		Email:    string(*userUpdate.Email),
		Password: *userUpdate.Password,
	}

	updatedUser, err := h.Service.PatchUser(uint(userId), toUpdateUser)
	if err != nil {
		return nil, err
	}
	response := users.UpdateUserByID200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    (*openapi_types.Email)(&updatedUser.Email),
		Password: &updatedUser.Password,
	}
	return response, nil
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

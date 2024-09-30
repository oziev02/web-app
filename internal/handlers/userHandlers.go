package handlers

import (
	"context"
	"web-app/internal/userService"
	"web-app/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {

	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:    &usr.ID,
			Email: &usr.Email,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {

	userRequest := request.Body

	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:    &createdUser.ID,
		Email: &createdUser.Email,
	}
	return response, nil
}

func (h *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {

	userRequest := request.Body

	userToUpdate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	updatedUser, err := h.Service.UpdateUserByID(request.Id, userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:    &updatedUser.ID,
		Email: &updatedUser.Email,
	}
	return response, nil
}

func (h *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id

	err := h.Service.DeleteUserByID(id)
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

package service

import (
	"gitlab.com/kiplagatcollins/flowershop/data/request"
	"gitlab.com/kiplagatcollins/flowershop/data/response"
	"gitlab.com/kiplagatcollins/flowershop/model"
	"gitlab.com/kiplagatcollins/flowershop/repository"

	"github.com/go-playground/validator/v10"
	"gitlab.com/kiplagatcollins/flowershop/helper"
)

type UserService interface {
	Create(User request.CreateUserRequest)
	Update(User request.UpdateUserRequest)
	Delete(UserId int)
	FindById(UserId int) response.UserResponse
	FindAll() []response.UserResponse
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Create implements UserService
func (t *UserServiceImpl) Create(User request.CreateUserRequest) {
	err := t.Validate.Struct(User)
	helper.ErrorPanic(err)
	tagModel := model.User{
		Name: User.Name,
	}
	t.UserRepository.Save(tagModel)
}

// Delete implements UserService
func (t *UserServiceImpl) Delete(UserId int) {
	t.UserRepository.Delete(UserId)
}

// FindAll implements UserService
func (t *UserServiceImpl) FindAll() []response.UserResponse {
	result := t.UserRepository.FindAll()

	var User []response.UserResponse
	for _, value := range result {
		tag := response.UserResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		User = append(User, tag)
	}

	return User
}

// FindById implements UserService
func (t *UserServiceImpl) FindById(UserId int) response.UserResponse {
	tagData, err := t.UserRepository.FindById(UserId)
	helper.ErrorPanic(err)

	tagResponse := response.UserResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

// Update implements UserService
func (t *UserServiceImpl) Update(User request.UpdateUserRequest) {
	tagData, err := t.UserRepository.FindById(User.Id)
	helper.ErrorPanic(err)
	tagData.Name = User.Name
	t.UserRepository.Update(tagData)
}

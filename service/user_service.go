package service

import (
	"errors"

	"gitlab.com/kiplagatcollins/flowershop/data/request"
	"gitlab.com/kiplagatcollins/flowershop/data/response"
	"gitlab.com/kiplagatcollins/flowershop/model"
	"gitlab.com/kiplagatcollins/flowershop/repository"

	"github.com/go-playground/validator/v10"
)

var ErrUserNotFound = errors.New("user not found")

type UserService interface {
	Create(User request.CreateUserRequest) error
	Update(User request.UpdateUserRequest) error
	Delete(UserId int) error
	FindById(UserId int) (*response.UserResponse, error)
	FindAll() ([]*response.UserResponse, error)
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
func (t *UserServiceImpl) Create(User request.CreateUserRequest) error {
	err := t.Validate.Struct(User)
	if err != nil {
		return err // Return validation error
	}

	userModel := model.User{
		Username:    User.Username,
		FirstName:   User.FirstName,
		LastName:    User.LastName,
		Address:     User.Address,
		Email:       User.Email,
		DOB:         User.DOB,
		PhoneNumber: User.PhoneNumber,
		Role:        User.Role,
	}

	return t.UserRepository.Save(userModel)
}

// Delete implements UserService
func (t *UserServiceImpl) Delete(UserId int) error {
	return t.UserRepository.Delete(UserId)
}

// FindAll implements UserService
func (t *UserServiceImpl) FindAll() ([]*response.UserResponse, error) {
	result, err := t.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var User []*response.UserResponse
	for _, value := range result {
		user := &response.UserResponse{
			ID:          value.ID,
			Username:    value.Username,
			FirstName:   value.FirstName,
			LastName:    value.LastName,
			Address:     value.Address,
			Email:       value.Email,
			DOB:         value.DOB,
			PhoneNumber: value.PhoneNumber,
			Role:        value.Role,
			CreateAt:    value.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt:    value.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		User = append(User, user)
	}

	return User, nil
}

// FindById implements UserService
func (t *UserServiceImpl) FindById(UserId int) (*response.UserResponse, error) {
	userData, err := t.UserRepository.FindById(UserId)
	if err != nil {
		return nil, err
	}

	userResponse := &response.UserResponse{
		ID:          userData.ID,
		Username:    userData.Username,
		FirstName:   userData.FirstName,
		LastName:    userData.LastName,
		Address:     userData.Address,
		Email:       userData.Email,
		DOB:         userData.DOB,
		PhoneNumber: userData.PhoneNumber,
		Role:        userData.Role,
		CreateAt:    userData.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:    userData.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return userResponse, nil
}

// Update implements UserService
func (t *UserServiceImpl) Update(User request.UpdateUserRequest) error {
	userData, err := t.UserRepository.FindById(User.Id)
	if err != nil {
		return err
	}

	userData.Username = User.Username
	userData.FirstName = User.FirstName
	userData.LastName = User.LastName
	userData.Address = User.Address
	userData.Email = User.Email
	userData.DOB = User.DOB
	userData.PhoneNumber = User.PhoneNumber
	userData.Role = User.Role

	return t.UserRepository.Update(userData)
}

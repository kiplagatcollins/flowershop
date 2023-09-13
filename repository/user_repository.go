package repository

import (
	"errors"

	"gitlab.com/kiplagatcollins/flowershop/data/request"
	"gitlab.com/kiplagatcollins/flowershop/helper"
	"gitlab.com/kiplagatcollins/flowershop/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(User model.User)
	Update(User model.User)
	Delete(UserId int)
	FindById(UserId int) (User model.User, err error)
	FindAll() []model.User
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserREpositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete implements UserRepository
func (t *UserRepositoryImpl) Delete(UserId int) {
	var User model.User
	result := t.Db.Where("id = ?", UserId).Delete(&User)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository
func (t *UserRepositoryImpl) FindAll() []model.User {
	var User []model.User
	result := t.Db.Find(&User)
	helper.ErrorPanic(result.Error)
	return User
}

// FindById implements UserRepository
func (t *UserRepositoryImpl) FindById(UserId int) (User model.User, err error) {
	var tag model.User
	result := t.Db.Find(&tag, UserId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// Save implements UserRepository
func (t *UserRepositoryImpl) Save(User model.User) {
	result := t.Db.Create(&User)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepository
func (t *UserRepositoryImpl) Update(User model.User) {
	var updateTag = request.UpdateUserRequest{
		Id:   User.Id,
		Name: User.Name,
	}
	result := t.Db.Model(&User).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

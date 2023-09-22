package repository

import (
	"errors"

	"gitlab.com/kiplagatcollins/flowershop/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user model.User) error
	Update(user model.User) error
	Delete(userID int) error
	FindById(userID int) (model.User, error)
	FindAll() ([]model.User, error)
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

// Save implements UserRepository
func (r *UserRepositoryImpl) Save(user model.User) error {
	result := r.Db.Create(&user)
	return result.Error
}

// Update implements UserRepository
func (r *UserRepositoryImpl) Update(user model.User) error {
	result := r.Db.Save(&user)
	return result.Error
}

// Delete implements UserRepository
func (r *UserRepositoryImpl) Delete(userID int) error {
	result := r.Db.Delete(&model.User{}, userID)
	return result.Error
}

// FindAll implements UserRepository
func (r *UserRepositoryImpl) FindAll() ([]model.User, error) {
	var users []model.User
	result := r.Db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// FindById implements UserRepository
func (r *UserRepositoryImpl) FindById(userID int) (model.User, error) {
	var user model.User
	result := r.Db.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, result.Error
	}
	return user, nil
}

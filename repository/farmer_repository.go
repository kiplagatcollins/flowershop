package repository

import (
	"errors"

	"gitlab.com/kiplagatcollins/flowershop/model"
	"gorm.io/gorm"
)

type FarmerRepository interface {
	Save(farmer model.Farmer) error
	Update(farmer model.Farmer) error
	Delete(farmerID int) error
	FindById(farmerID int) (model.Farmer, error)
	FindAll() ([]model.Farmer, error)
}

type FarmerRepositoryImpl struct {
	Db *gorm.DB
}

func NewFarmerRepositoryImpl(db *gorm.DB) FarmerRepository {
	return &FarmerRepositoryImpl{Db: db}
}

// Save implements FarmerRepository
func (r *FarmerRepositoryImpl) Save(farmer model.Farmer) error {
	result := r.Db.Create(&farmer)
	return result.Error
}

// Update implements FarmerRepository
func (r *FarmerRepositoryImpl) Update(farmer model.Farmer) error {
	result := r.Db.Save(&farmer)
	return result.Error
}

// Delete implements FarmerRepository
func (r *FarmerRepositoryImpl) Delete(farmerID int) error {
	result := r.Db.Delete(&model.Farmer{}, farmerID)
	return result.Error
}

// FindAll implements FarmerRepository
func (r *FarmerRepositoryImpl) FindAll() ([]model.Farmer, error) {
	var farmers []model.Farmer
	result := r.Db.Preload("User").Find(&farmers)
	r.Db.Preload("User").First(&farmers)
	if result.Error != nil {
		return nil, result.Error
	}
	return farmers, nil
}

// FindById implements FarmerRepository
func (r *FarmerRepositoryImpl) FindById(farmerID int) (model.Farmer, error) {
	var farmer model.Farmer
	result := r.Db.Preload("User").First(&farmer, farmerID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Farmer{}, errors.New("farmer not found")
		}
		return model.Farmer{}, result.Error
	}
	return farmer, nil
}

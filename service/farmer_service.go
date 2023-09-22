package service

import (
	"errors"

	"gitlab.com/kiplagatcollins/flowershop/data/request"
	"gitlab.com/kiplagatcollins/flowershop/data/response"
	"gitlab.com/kiplagatcollins/flowershop/model"
	"gitlab.com/kiplagatcollins/flowershop/repository"

	"github.com/go-playground/validator/v10"
)

var ErrFarmerNotFound = errors.New("farmer not found")

type FarmerService interface {
	Create(Farmer request.CreateFarmerRequest) error
	Update(Farmer request.UpdateFarmerRequest) error
	Delete(FarmerId int) error
	FindById(FarmerId int) (*response.FarmerResponse, error)
	FindAll() ([]*response.FarmerResponse, error)
}

type FarmerServiceImpl struct {
	FarmerRepository repository.FarmerRepository
	Validate         *validator.Validate
}

func NewFarmerServiceImpl(farmerRepository repository.FarmerRepository, validate *validator.Validate) FarmerService {
	return &FarmerServiceImpl{
		FarmerRepository: farmerRepository,
		Validate:         validate,
	}
}

// Create implements FarmerService
func (t *FarmerServiceImpl) Create(Farmer request.CreateFarmerRequest) error {
	err := t.Validate.Struct(Farmer)
	if err != nil {
		return err // Return validation error
	}

	farmerModel := model.Farmer{
		UserId:             Farmer.UserId,
		CompanyName:        Farmer.CompanyName,
		CompanyTradingName: Farmer.CompanyTradingName,
		CompanyEmail:       Farmer.CompanyEmail,
	}

	return t.FarmerRepository.Save(farmerModel)
}

// Delete implements FarmerService
func (t *FarmerServiceImpl) Delete(FarmerId int) error {
	return t.FarmerRepository.Delete(FarmerId)
}

// FindAll implements FarmerService
func (t *FarmerServiceImpl) FindAll() ([]*response.FarmerResponse, error) {
	result, err := t.FarmerRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var Farmer []*response.FarmerResponse
	for _, value := range result {
		farmer := &response.FarmerResponse{
			ID: value.ID,
			User: response.UserResponse{
				ID:          value.User.ID,
				Username:    value.User.Username,
				FirstName:   value.User.FirstName,
				LastName:    value.User.LastName,
				Address:     value.User.Address,
				Email:       value.User.Email,
				DOB:         value.User.DOB,
				PhoneNumber: value.User.PhoneNumber,
				Role:        value.User.Role,
				CreateAt:    value.User.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdateAt:    value.User.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
			CompanyName:        value.CompanyName,
			CompanyTradingName: value.CompanyTradingName,
			CompanyEmail:       value.CompanyEmail,
			CreateAt:           value.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt:           value.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		Farmer = append(Farmer, farmer)
	}

	return Farmer, nil
}

// FindById implements FarmerService
func (t *FarmerServiceImpl) FindById(FarmerId int) (*response.FarmerResponse, error) {
	farmerData, err := t.FarmerRepository.FindById(FarmerId)
	if err != nil {
		return nil, err
	}

	farmerResponse := &response.FarmerResponse{
		ID: farmerData.ID,
		User: response.UserResponse{
			ID:          farmerData.User.ID,
			Username:    farmerData.User.Username,
			FirstName:   farmerData.User.FirstName,
			LastName:    farmerData.User.LastName,
			Address:     farmerData.User.Address,
			Email:       farmerData.User.Email,
			DOB:         farmerData.User.DOB,
			PhoneNumber: farmerData.User.PhoneNumber,
			Role:        farmerData.User.Role,
			CreateAt:    farmerData.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt:    farmerData.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
		CompanyName:        farmerData.CompanyName,
		CompanyTradingName: farmerData.CompanyTradingName,
		CompanyEmail:       farmerData.CompanyEmail,
		CreateAt:           farmerData.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:           farmerData.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return farmerResponse, nil
}

// Update implements FarmerService
func (t *FarmerServiceImpl) Update(farmer request.UpdateFarmerRequest) error {
	farmerData, err := t.FarmerRepository.FindById(farmer.ID)
	if err != nil {
		return err
	}

	farmerData.UserId = farmer.UserId
	farmerData.CompanyName = farmer.CompanyName
	farmerData.CompanyTradingName = farmer.CompanyTradingName
	farmerData.CompanyEmail = farmer.CompanyEmail
	return t.FarmerRepository.Update(farmerData)
}

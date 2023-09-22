package controller

import (
	"net/http"
	"strconv"

	"gitlab.com/kiplagatcollins/flowershop/data/request"
	"gitlab.com/kiplagatcollins/flowershop/data/response"
	"gitlab.com/kiplagatcollins/flowershop/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type FarmerController struct {
	FarmerService service.FarmerService
}

func NewFarmerController(service service.FarmerService) *FarmerController {
	return &FarmerController{
		FarmerService: service,
	}
}

// Create handles the creation of a new farmer.
func (controller *FarmerController) Create(ctx *gin.Context) {
	log.Info().Msg("create Farmer")
	createFarmerRequest := request.CreateFarmerRequest{}
	if err := ctx.ShouldBindJSON(&createFarmerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "Bad request", Data: err.Error()})
		return
	}

	if err := controller.FarmerService.Create(createFarmerRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: http.StatusInternalServerError, Status: "Internal server error", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "Ok", Data: nil})
}

// Update handles updating farmer details.
func (controller *FarmerController) Update(ctx *gin.Context) {
	log.Info().Msg("update Farmer")
	updateFarmerRequest := request.UpdateFarmerRequest{}
	if err := ctx.ShouldBindJSON(&updateFarmerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "Bad request", Data: err.Error()})
		return
	}

	farmerId := ctx.Param("farmerId")
	id, err := strconv.Atoi(farmerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "Bad request", Data: err.Error()})
		return
	}
	updateFarmerRequest.ID = id

	if err := controller.FarmerService.Update(updateFarmerRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: http.StatusInternalServerError, Status: "Internal server error", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "Ok", Data: nil})
}

// Delete handles deleting a farmer by ID.
func (controller *FarmerController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete Farmer")
	farmerId := ctx.Param("farmerId")
	id, err := strconv.Atoi(farmerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "Bad request", Data: err.Error()})
		return
	}

	if err := controller.FarmerService.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: http.StatusInternalServerError, Status: "Internal server error", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "Ok", Data: nil})
}

// FindById returns a single farmer by ID.
func (controller *FarmerController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid Farmer")
	farmerId := ctx.Param("farmerId")
	id, err := strconv.Atoi(farmerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "Bad request", Data: err.Error()})
		return
	}

	farmerResponse, err := controller.FarmerService.FindById(id)
	if err != nil {
		if err.Error() == service.ErrFarmerNotFound.Error() {
			ctx.JSON(http.StatusNotFound, response.Response{Code: http.StatusNotFound, Status: "Farmer not found", Data: err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, response.Response{Code: http.StatusInternalServerError, Status: "Internal server error", Data: err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "Ok", Data: farmerResponse})
}

// FindAll returns a list of all farmers.
func (controller *FarmerController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll Farmer")
	farmersResponse, err := controller.FarmerService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: http.StatusInternalServerError, Status: "Internal server error", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "Ok", Data: farmersResponse})
}

package controller

import (
	"net/http"
	"strconv"

	"gitlab.com/kiplagatcollins/flowershop/data/request"
	"gitlab.com/kiplagatcollins/flowershop/data/response"
	"gitlab.com/kiplagatcollins/flowershop/helper"
	"gitlab.com/kiplagatcollins/flowershop/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

// @Summary 		Create User User
// @Description 	Create User User details
// @Tags 			User
// @Accept 			json
// @Produce 		json
// @Param 			input body request.CreateUserRequest true "request.CreateUserRequest details"
// @Success 		200 {object} response.Response{data=model.User} "Successfully retrieved user"
// @Failure 		400 {object} response.Response "Bad request"
// @Failure 		401 {object} response.Response "Unauthorized"
// @Failure 		404 {object} response.Response "User not found"
// @Failure 		409 {object} response.Response "Conflict"
// @Failure 		500 {object} response.Response "Internal server error"
// @Router 			/users [post]
func (controller *UserController) Create(ctx *gin.Context) {
	log.Info().Msg("create User")
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	controller.UserService.Create(createUserRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateUser		godoc
// @Summary 		Update User
// @Description 	Get products based on agent and customer details
// @Tags 			User
// @Accept 			application/json
// @Produce 		application/json
// @Param 			input body request.UpdateUserRequest true "request.UpdateUserRequest details"
// @Success 		200 {object} response.Response{data=model.User} "Successfully retrieved user"
// @Failure 		400 {object} response.Response "Bad request"
// @Failure 		401 {object} response.Response "Unauthorized"
// @Failure 		404 {object} response.Response "User not found"
// @Failure 		409 {object} response.Response "Conflict"
// @Failure 		500 {object} response.Response "Internal server error"
// @Router			/users/userId} [patch]
func (controller *UserController) Update(ctx *gin.Context) {
	log.Info().Msg("update User")
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateUserRequest.Id = id

	controller.UserService.Update(updateUserRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteUser 		godoc
// @Summary			Get Single User by id.
// @Description 	Return the user whoes userId value mathes id.
// @Tags 			User
// @Accept 			application/json
// @Produce     	application/json
// @Success 		200 {object} response.Response{data=model.User} "Successfully retrieved user"
// @Failure 		400 {object} response.Response "Bad request"
// @Failure 		401 {object} response.Response "Unauthorized"
// @Failure 		404 {object} response.Response "User not found"
// @Failure 		409 {object} response.Response "Conflict"
// @Failure 		500 {object} response.Response "Internal server error"
// @Router			/users/{userId} [delete]
func (controller *UserController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete User")
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.UserService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdUser 	godoc
// @Summary 		Get Single User by ID
// @Description 	Returns the user whose userID matches the provided ID.
// @Tags 			User
// @Accept 			json
// @Produce 		json
// @Param 			userId path int true "User ID to fetch"
// @Success 		200 {object} response.Response{data=model.User} "Successfully retrieved user"
// @Failure 		400 {object} response.Response "Bad request"
// @Failure 		401 {object} response.Response "Unauthorized"
// @Failure 		404 {object} response.Response "User not found"
// @Failure 		409 {object} response.Response "Conflict"
// @Failure 		500 {object} response.Response "Internal server error"
// @Router 			/users/{userId} [get]
func (controller *UserController) FindById(ctx *gin.Context) {

	log.Info().Msg("findbyid User")
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	userResponse := controller.UserService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllUser 		godoc
// @Summary Get Occupations
// @Description Get Occupations based on agent and customer details
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]model.User} "Successfully retrieved user"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 404 {object} response.Response "User not found"
// @Failure 409 {object} response.Response "Conflict"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users [get]
func (controller *UserController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll User")
	usersResponse := controller.UserService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   usersResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

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

type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

// Create handles the creation of a new user.
func (controller *UserController) Create(ctx *gin.Context) {
	log.Info().Msg("create User")
	createUserRequest := request.CreateUserRequest{}
	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "Bad request", Data: err.Error()})
		return
	}

	if err := controller.UserService.Create(createUserRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: http.StatusInternalServerError, Status: "Internal server error", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "Ok", Data: nil})
}

// Update handles updating user details.
func (controller *UserController) Update(ctx *gin.Context) {
	log.Info().Msg("update User")
	updateUserRequest := request.UpdateUserRequest{}
	if err := ctx.ShouldBindJSON(&updateUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "Bad request", Data: err.Error()})
		return
	}

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "Bad request", Data: err.Error()})
		return
	}
	updateUserRequest.Id = id

	if err := controller.UserService.Update(updateUserRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: http.StatusInternalServerError, Status: "Internal server error", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "Ok", Data: nil})
}

// Delete handles deleting a user by ID.
func (controller *UserController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete User")
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "Bad request", Data: err.Error()})
		return
	}

	if err := controller.UserService.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: http.StatusInternalServerError, Status: "Internal server error", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "Ok", Data: nil})
}

// FindById returns a single user by ID.
func (controller *UserController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid User")
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "Bad request", Data: err.Error()})
		return
	}

	userResponse, err := controller.UserService.FindById(id)
	if err != nil {
		if err.Error() == service.ErrUserNotFound.Error() {
			ctx.JSON(http.StatusNotFound, response.Response{Code: http.StatusNotFound, Status: "User not found", Data: err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, response.Response{Code: http.StatusInternalServerError, Status: "Internal server error", Data: err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "Ok", Data: userResponse})
}

// FindAll returns a list of all users.
func (controller *UserController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll User")
	usersResponse, err := controller.UserService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: http.StatusInternalServerError, Status: "Internal server error", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "Ok", Data: usersResponse})
}

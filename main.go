package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"gitlab.com/kiplagatcollins/flowershop/config"
	"gitlab.com/kiplagatcollins/flowershop/controller"
	"gitlab.com/kiplagatcollins/flowershop/helper"
	"gitlab.com/kiplagatcollins/flowershop/model"
	"gitlab.com/kiplagatcollins/flowershop/repository"
	"gitlab.com/kiplagatcollins/flowershop/router"
	"gitlab.com/kiplagatcollins/flowershop/service"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("user").AutoMigrate(&model.User{})

	// Repository
	userRepository := repository.NewUserREpositoryImpl(db)

	// Service
	userService := service.NewUserServiceImpl(userRepository, validate)

	// Controller
	userController := controller.NewUserController(userService)

	// Router
	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}

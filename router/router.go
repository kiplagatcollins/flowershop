package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/kiplagatcollins/flowershop/controller"
)

func NewRouter(userController *controller.UserController, farmerController *controller.FarmerController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	// Users routes
	UserRouter := baseRouter.Group("/users")
	UserRouter.GET("", userController.FindAll)
	UserRouter.GET("/:userId", userController.FindById)
	UserRouter.POST("", userController.Create)
	UserRouter.PATCH("/:userId", userController.Update)
	UserRouter.DELETE("/:userId", userController.Delete)

	// Users routes
	FarmerRouter := baseRouter.Group("/farmers")
	FarmerRouter.GET("", farmerController.FindAll)
	FarmerRouter.GET("/:farmerId", farmerController.FindById)
	FarmerRouter.POST("", farmerController.Create)
	FarmerRouter.PATCH("/:farmerId", farmerController.Update)
	FarmerRouter.DELETE("/:farmerId", farmerController.Delete)

	return router
}

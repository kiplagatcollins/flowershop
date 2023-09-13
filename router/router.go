package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/kiplagatcollins/flowershop/controller"
)

func NewRouter(UserController *controller.UserController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	UserRouter := baseRouter.Group("/User")
	UserRouter.GET("", UserController.FindAll)
	UserRouter.GET("/:tagId", UserController.FindById)
	UserRouter.POST("", UserController.Create)
	UserRouter.PATCH("/:tagId", UserController.Update)
	UserRouter.DELETE("/:tagId", UserController.Delete)

	return router
}

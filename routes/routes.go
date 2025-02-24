package routes

import (
	"users-backend/controllers"
	"users-backend/services"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) {
	userService := services.NewUserService(services.DB)
	userController := controllers.NewUserController(userService)

	// กำหนด Routes
	router.DELETE("/users/:id", userController.DeleteUserHandler)
	router.POST("/users", userController.CreateUserHandler)
	router.GET("/users", userController.GetAllUsersHandler)
	router.GET("/users/:id", userController.GetUserByIDHandler)
	router.PUT("/users/:id", userController.UpdateUserHandler)

}

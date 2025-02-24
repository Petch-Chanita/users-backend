package controllers

import (
	"net/http"
	"users-backend/models"
	"users-backend/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// CreateUserHandler - Handler สำหรับการสร้างผู้ใช้
func (c *UserController) CreateUserHandler(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := c.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success full",
		"result":  createdUser,
		"status":  "ok",
	})
}

// GetUserHandler - Handler สำหรับการดึงข้อมูลผู้ใช้
func (c *UserController) GetUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.UserService.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// UpdateUserHandler - Handler สำหรับการอัปเดตข้อมูลผู้ใช้
func (c *UserController) UpdateUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := c.UserService.UpdateUser(id, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success full",
		"result":  updatedUser,
		"status":  "ok",
	})
}

// DeleteUserHandler - Handler สำหรับการลบผู้ใช้
func (c *UserController) DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.UserService.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *UserController) GetUserByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.UserService.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetAllUsersHandler(ctx *gin.Context) {
	users, err := c.UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success full",
		"result":  users,
		"status":  "ok",
	})
}

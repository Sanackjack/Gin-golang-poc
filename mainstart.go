package main

import (
	"edge/config"
	"edge/controllers"
	"edge/helper"
	"edge/models"
	"edge/repository"
	"edge/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	//set connection database
	db := config.DatabaseConnection()
	db.Table("cat").AutoMigrate(&models.User{})
	validate := validator.New()
	// Initialize services and controllers , Repository

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := services.NewUserService(userRepository, validate)
	userController := controllers.NewUserController(userService)
	r := gin.Default()

	r.Use(gin.CustomRecovery(helper.ErrorHandler))

	// set route
	r.POST("/users", userController.CreateUserHandler)
	r.GET("/users", userController.FindAllUser)
	r.GET("/user/:userId", userController.GetUserByIdHandler)

	// run default 8080
	r.Run()
}

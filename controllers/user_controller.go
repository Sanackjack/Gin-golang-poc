package controllers

import (

	//"edge/data/request"
	//"edge/data/request"
	//"edge/data/request"
	. "edge/data/request"
	"edge/data/response"
	. "edge/helper"
	"edge/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetUserByIdHandler(ctx *gin.Context) {

	idparam := ctx.Param("userId")
	log.Info().Msg("find user id " + idparam)

	id, err := strconv.Atoi(idparam)
	if err != nil {
		// Handle the error if conversion fails
		ErrorPanic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	res := c.userService.GetUserById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   res,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UserController) CreateUserHandler(ctx *gin.Context) {

	createUsersRequest := CreateUsersRequest{}

	if err := ctx.ShouldBindJSON(&createUsersRequest); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//return
		ErrorPanic(err)
	}

	//if err := c.userService.CreateUser(&createUsersRequest); err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
	//	return
	//}

	c.userService.CreateUser(createUsersRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   gin.H{"message": "User created successfully"},
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UserController) FindAllUser(ctx *gin.Context) {
	log.Info().Msg("findAll users")
	userResponse := controller.userService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

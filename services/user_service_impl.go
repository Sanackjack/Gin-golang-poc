package services

import (
	"edge/data/request"
	. "edge/helper"
	"net/http"

	//"edge/helper/CustomError"
	"edge/models"
	"edge/repository"
	"github.com/go-playground/validator/v10"
	//"net/http"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
	// Service-specific dependencies or state can be added here
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) UserService {

	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}

}

func (s *UserServiceImpl) CreateUser(user request.CreateUsersRequest) models.User {
	// Perform business logic (e.g., save user to a database)
	// In this example, just print the user information
	//println("Creating user:", user.Username)

	err := s.Validate.Struct(user)
	ErrorPanic(err)

	ErrorPanic(ValidateException{Message: "Validate Error Message test", Code: http.StatusBadRequest})

	//ErrorPanic(&ValidateException{Message: "Validate Error Message test", Code: http.StatusBadRequest})
	//	helper.ErrorPanic(&helper.CustomException{Message: "Custom Error Message test", Code: http.StatusBadRequest})

	userModel := models.User{
		Name: user.Name,
	}
	err = s.UserRepository.Save(userModel)

	if err != nil {
		ErrorPanic(err)
	}
	return userModel
}

func (s *UserServiceImpl) GetUserById(userId int) models.User {

	user, err := s.UserRepository.FindById(userId)
	//helper.ErrorPanic(&helper.CustomError{Message: "Custom Error Message", Code: http.StatusBadRequest})
	ErrorPanic(err)
	//panic(&CustomError{Message: "Custom Error Message", Code: http.StatusBadRequest})
	return user
}

func (t *UserServiceImpl) FindAll() []models.User {
	result := t.UserRepository.FindAll()

	return result
}

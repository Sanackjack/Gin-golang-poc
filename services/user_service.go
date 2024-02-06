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

type UserService struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
	// Service-specific dependencies or state can be added here
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) *UserService {

	return &UserService{
		UserRepository: userRepository,
		Validate:       validate,
	}

}

func (s *UserService) CreateUser(user *request.CreateUsersRequest) {
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
	s.UserRepository.Save(userModel)

}

func (s *UserService) GetUserById(userId int) models.User {

	user, err := s.UserRepository.FindById(userId)
	//helper.ErrorPanic(&helper.CustomError{Message: "Custom Error Message", Code: http.StatusBadRequest})
	ErrorPanic(err)
	//panic(&CustomError{Message: "Custom Error Message", Code: http.StatusBadRequest})
	return user
}

func (t *UserService) FindAll() []models.User {
	result := t.UserRepository.FindAll()

	return result
}

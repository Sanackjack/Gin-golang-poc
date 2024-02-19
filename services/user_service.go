package services

import (
	"edge/data/request"
	"edge/models"
)

type UserService interface {
	CreateUser(req request.CreateUsersRequest) (user models.User)
	GetUserById(userId int) (user models.User)
	FindAll() (user []models.User)
}

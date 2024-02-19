package repository

import (
	"edge/models"
)

type UserRepository interface {
	Save(user models.User) (err error)

	Delete(tagsId int)
	FindById(tagsId int) (tags models.User, err error)
	FindAll() []models.User
}

package repository

import (
	"edge/models"
)

type UserRepository interface {
	Save(tags models.User)

	Delete(tagsId int)
	FindById(tagsId int) (tags models.User, err error)
	FindAll() []models.User
}

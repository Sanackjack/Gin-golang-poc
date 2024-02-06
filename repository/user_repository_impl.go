package repository

import (
	"errors"
	"net/http"

	//"golang-crud-gin/data/request"
	"edge/helper"
	"edge/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete implements TagsRepository
func (t *UserRepositoryImpl) Delete(tagsId int) {
	var tags models.User
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository
func (t *UserRepositoryImpl) FindAll() []models.User {
	var tags []models.User
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

// FindById implements TagsRepository
func (t *UserRepositoryImpl) FindById(tagsId int) (tags models.User, err error) {

	var usr models.User
	result := t.Db.Find(&usr, tagsId)
	if result != nil {
		return usr, nil
	} else {
		return usr, errors.New("usr is not found")
	}

}

// Save implements TagsRepository
func (t *UserRepositoryImpl) Save(tags models.User) {
	result := t.Db.Create(&tags)
	//result.Error = errors.New("test")
	if result.Error != nil {
		helper.ErrorPanic(&helper.CustomException{Message: "Save data error", Code: http.StatusBadRequest})
	}

}

package user

import (
	"Project/playground/be6/layered/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) Get() ([]entities.User, error) {
	arrUser := []entities.User{}

	if err := ur.database.Find(&arrUser).Error; err != nil {
		return nil, err
	}

	return arrUser, nil
}

func (ur *UserRepository) Insert(newUser entities.User) (entities.User, error) {
	if err := ur.database.Create(&newUser).Error; err != nil {
		return newUser, err
	}

	return newUser, nil
}

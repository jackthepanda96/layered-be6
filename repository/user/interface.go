package user

import "Project/playground/be6/layered/entities"

type User interface {
	Get() ([]entities.User, error)
	Insert(newUser entities.User) (entities.User, error)
}

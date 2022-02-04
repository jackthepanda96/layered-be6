package user

import (
	"Project/playground/be6/layered/delivery/controllers/common"
	"Project/playground/be6/layered/entities"
	"Project/playground/be6/layered/repository/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo user.User
}

func New(repository user.User) *UserController {
	return &UserController{
		repo: repository,
	}
}

func (uc *UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All User", res))
	}
}

func (uc *UserController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		newUser := RegisterRequestFormat{}

		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := uc.repo.Insert(entities.User{Nama: newUser.Nama, Email: newUser.Email, Password: newUser.Password})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create User", res))
	}
}

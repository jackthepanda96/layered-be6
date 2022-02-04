package routes

import (
	"Project/playground/be6/layered/delivery/controllers/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uc *user.UserController) {
	e.GET("users", uc.Get())
	e.POST("users", uc.Insert())
}

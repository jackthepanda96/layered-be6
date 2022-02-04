package main

import (
	"Project/playground/be6/layered/configs"
	"Project/playground/be6/layered/delivery/controllers/user"
	"Project/playground/be6/layered/delivery/routes"
	userRepo "Project/playground/be6/layered/repository/user"
	"Project/playground/be6/layered/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	userRepo := userRepo.New(db)

	userController := user.New(userRepo)

	e := echo.New()

	routes.RegisterPath(e, userController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}

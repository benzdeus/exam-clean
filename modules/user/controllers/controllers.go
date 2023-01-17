package controllers

import (
	"demo-clean/entitiies"

	"github.com/labstack/echo/v4"
)

type userController struct {
	usecase entitiies.UserUsecase
}

func InitUserController(usecase entitiies.UserUsecase) *userController {
	return &userController{
		usecase,
	}
}

func (con userController) AddUser(ctx echo.Context) error {
	user := new(entitiies.UserModel)

	err := ctx.Bind(&user)

	if err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userRes, err := con.usecase.AddUser(user)

	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, userRes)
}

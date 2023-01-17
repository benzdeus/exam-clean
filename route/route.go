package route

import (
	"demo-clean/modules/user/controllers"
	"demo-clean/modules/user/repositories"
	"demo-clean/modules/user/usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, dbPG *gorm.DB) {

	userRepo := repositories.NewUserRepo(dbPG)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userController := controllers.InitUserController(userUsecase)

	e.POST("/user", userController.AddUser)

}

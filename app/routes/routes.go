package routes

import (
	"gamePicker/controllers/user"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController user.UserController
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.POST("user/login", cl.UserController.Login)
	e.POST("user/register", cl.UserController.CreateUserController)
	e.GET("users", cl.UserController.GetUserController)
	e.GET("user/:id", cl.UserController.GetUserByIDController)
}

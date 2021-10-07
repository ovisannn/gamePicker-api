package routes

import (
	"gamePicker/controllers/user"
	walley "gamePicker/controllers/wallet"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController   user.UserController
	WalletController walley.WalletController
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	// user
	e.POST("user/login", cl.UserController.Login)
	e.POST("user/register", cl.UserController.CreateUserController)
	e.GET("users", cl.UserController.GetUserController)
	e.GET("user/:id", cl.UserController.GetUserByIDController)
	// wallet
	e.POST("wallet/update/:id", cl.WalletController.InsertWalletController)
}

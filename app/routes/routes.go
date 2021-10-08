package routes

import (
	"gamePicker/controllers/user"
	walley "gamePicker/controllers/wallet"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	UserController   user.UserController
	WalletController walley.WalletController
	JWTConfig        middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.Use(middleware.Logger())
	// user
	e.POST("user/login", cl.UserController.Login)
	e.POST("user/register", cl.UserController.CreateUserController)
	e.GET("users", cl.UserController.GetUserController)
	e.GET("user/:id", cl.UserController.GetUserByIDController)
	// wallet
	e.PUT("wallet/update/:id", cl.WalletController.InsertWalletController)
	e.GET("wallet/Get/:id", cl.WalletController.GetWalletByIDController)
}

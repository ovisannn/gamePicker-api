package main

import (
	"log"
	"time"

	// _userDb "gamePicker/drivers/database/user"
	"gamePicker/app/routes"
	_mysqlDriver "gamePicker/drivers/mysql"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_userUseCase "gamePicker/business/user"
	_userController "gamePicker/controllers/user"
	_userRepository "gamePicker/drivers/database/user"

	_walleyUseCase "gamePicker/business/wallet"
	_walleyController "gamePicker/controllers/wallet"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	Conn := configDB.InitialDB()

	e := echo.New()

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUseCase.NewUserUseCase(userRepository, timeoutContext)
	userController := _userController.NewUserController(userUseCase)

	walletRepository := _userRepository.NewMysqlWalletRepository(Conn)
	walletUseCase := _walleyUseCase.NewWalletUseCase(walletRepository)
	walletController := _walleyController.NewWalletController(walletUseCase)

	routesInit := routes.ControllerList{
		UserController:   *userController,
		WalletController: *walletController,
	}

	routesInit.RouteRegister(*e)

	log.Fatal(e.Start((viper.GetString("server.address"))))
}

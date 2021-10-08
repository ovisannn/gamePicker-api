package handler

import (
	"gamePicker/business/wallet"
	"gamePicker/controllers"
	walletResponse "gamePicker/controllers/wallet/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type WalletController struct {
	WalletUseCae wallet.UseCase
}

func NewWalletController(walletUseCase wallet.UseCase) *WalletController {
	return &WalletController{
		WalletUseCae: walletUseCase,
	}
}

func (handler WalletController) InsertWalletController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	walletInsert := wallet.Domain{}
	walletInsert.IndieWallet, _ = strconv.Atoi(c.FormValue("indieWallet"))
	walletInsert.MoneyTarget, _ = strconv.Atoi(c.FormValue("moneyTarget"))
	walletInsert.MoneySaved, _ = strconv.Atoi(c.FormValue("moneySaved"))

	ctx := c.Request().Context()
	wallet, err := handler.WalletUseCae.InsertWalletController(ctx, walletInsert, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, walletResponse.FromDomainWallet(wallet))
}

func (handler WalletController) GetWalletByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ctx := c.Request().Context()
	wallet, err := handler.WalletUseCae.GetWalletByIDController(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, wallet)
}

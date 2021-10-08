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
	walletInsert.IndieWallet, _ = strconv.ParseFloat(c.FormValue("indieWallet"), 64)
	walletInsert.MoneyTarget, _ = strconv.ParseFloat(c.FormValue("moneyTarget"), 64)
	walletInsert.MoneySaved, _ = strconv.ParseFloat(c.FormValue("moneySaved"), 64)

	ctx := c.Request().Context()
	wallet, err := handler.WalletUseCae.InsertWalletController(ctx, walletInsert, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, walletResponse.FromDomainWallet(wallet))
}

package response

import "gamePicker/business/wallet"

type WalletResponse struct {
	MoneyTarget float64 `json:"moneyTarget"`
	MoneySaved  float64 `json:"moneySaved"`
}

func FromDomainWallet(domain wallet.Domain) WalletResponse {
	return WalletResponse{
		MoneyTarget: domain.MoneyTarget,
		MoneySaved:  domain.MoneySaved,
	}
}

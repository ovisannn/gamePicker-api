package response

import "gamePicker/business/wallet"

type WalletResponse struct {
	MoneyTarget int `json:"moneyTarget"`
	MoneySaved  int `json:"moneySaved"`
}

func FromDomainWallet(domain wallet.Domain) WalletResponse {
	return WalletResponse{
		MoneyTarget: domain.MoneyTarget,
		MoneySaved:  domain.MoneySaved,
	}
}

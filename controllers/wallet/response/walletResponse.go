package response

import "gamePicker/business/wallet"

type WalletResponse struct {
	Id_wallet   int `json:"id_wallet"`
	MoneySaved  int `json:"moneySaved"`
	MoneyTarget int `json:"moneyTarget"`
	IndieWallet int `json:"indieWallet"`
}

func FromDomainWallet(domain wallet.Domain) WalletResponse {
	return WalletResponse{
		MoneyTarget: domain.MoneyTarget,
		MoneySaved:  domain.MoneySaved,
	}
}

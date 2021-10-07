package request

type WalletRequest struct {
	MoneyTarget int `json:"moneyTarget"`
	MoneySaved  int `json:"moneySaved"`
}

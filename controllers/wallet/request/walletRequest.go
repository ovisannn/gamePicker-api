package request

type WalletRequest struct {
	MoneyTarget float64 `json:"moneyTarget"`
	MoneySaved  float64 `json:"moneySaved"`
}

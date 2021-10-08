package request

type WalletRequest struct {
	MoneyTarget float32 `json:"moneyTarget"`
	MoneySaved  float32 `json:"moneySaved"`
}

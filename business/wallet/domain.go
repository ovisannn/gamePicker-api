package wallet

import "context"

type Domain struct {
	Id_wallet   int `json:"id_wallet"`
	MoneySaved  int `json:"moneySaved"`
	MoneyTarget int `json:"moneyTarget"`
	IndieWallet int `json:"indieWallet"`
}

type UseCase interface {
	InsertWalletController(ctx context.Context, data Domain, id int) (Domain, error)
	GetWalletByIDController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	UpdateMoneySaved(ctx context.Context, data Domain, id int) (Domain, error)
	GetWalletByID(ctx context.Context, id int) (Domain, error)
}

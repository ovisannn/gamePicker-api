package wallet

import "context"

type Domain struct {
	Id_wallet   int
	MoneySaved  float64
	MoneyTarget float64
	IndieWallet float64
}

type UseCase interface {
	InsertWalletController(ctx context.Context, data Domain, id int) (Domain, error)
}

type Repository interface {
	UpdateMoneySaved(ctx context.Context, data Domain, id int) (Domain, error)
	GetWalletByID(ctx context.Context, id int) (Domain, error)
}

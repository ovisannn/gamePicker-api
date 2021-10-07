package wallet

import "context"

type Domain struct {
	Id_wallet   int
	MoneySaved  int
	MoneyTarget int
	IndieWallet int
}

type UseCase interface {
	InsertWalletController(ctx context.Context, data Domain, id int) (Domain, error)
}

type Repository interface {
	UpdateMoneySaved(ctx context.Context, data Domain, id int) (Domain, error)
}

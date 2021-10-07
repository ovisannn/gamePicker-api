package wallet

import "context"

type WalletUseCase struct {
	Repo Repository
}

func NewWalletUseCase(repo Repository) UseCase {
	return &WalletUseCase{
		Repo: repo,
	}
}

func (wl *WalletUseCase) InsertWalletController(ctx context.Context, data Domain, id int) (Domain, error) {
	wallet, err := wl.Repo.UpdateMoneySaved(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}
	return wallet, nil
}

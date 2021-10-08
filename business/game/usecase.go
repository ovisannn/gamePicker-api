package game

type GameUseCase struct {
	Repo Repository
}

func NewGameUseCase(repo Repository) UseCase {
	return &GameUseCase{
		Repo: repo,
	}
}

func (gc *GameUseCase) GetSalesController() ([]Domain_consume, error) {
	result := gc.Repo.GetOnSales()
	return result, nil
}

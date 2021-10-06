package user

import (
	"context"
	"errors"
	"time"
)

type UserUseCase struct {
	Repo           Repository
	contextTimeout context.Context
}

func NewUserUseCase(repo Repository, timeout time.Duration) UseCase {
	return &UserUseCase{
		Repo: repo,
		// contextTimeout: timeout,
	}
}

func (uc *UserUseCase) Login(ctx context.Context, email string, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("email empty")
	}
	if password == "" {
		return Domain{}, errors.New("password empty")
	}

	user, err := uc.Repo.Login(ctx, email, password)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetUserController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetUser(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}

func (uc *UserUseCase) GetUserByIDController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetUser_byID(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUseCase) CreateUserController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.CreateUser(ctx, data)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

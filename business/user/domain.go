package user

import (
	"context"
	"time"
)

type Domain struct {
	Id_user         int
	Username        string
	Password        string
	Email           string
	Name            string
	SteamProfile_id string `gorm:"steamProfile_id"`
	Detail_id       int
	Created_at      time.Time
	Updated_at      time.Time
}

type UseCase interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	GetUserController(ctx context.Context) ([]Domain, error)
	GetUserByIDController(ctx context.Context, id int) (Domain, error)
	CreateUserController(ctx context.Context, data Domain) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	GetUser(ctx context.Context) ([]Domain, error)
	GetUser_byID(ctx context.Context, id int) (Domain, error)
	CreateUser(ctx context.Context, data Domain) (Domain, error)
}

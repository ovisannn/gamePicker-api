package user

import (
	"gamePicker/business/user"
	_user "gamePicker/business/user"
	"time"
)

type User struct {
	Id_user         int `gorm:"primaryKey"`
	Username        string
	Password        string
	Email           string
	Name            string
	SteamProfile_id string
	Detail_id       string
	Created_at      time.Time
	Updated_at      time.Time
}

func (user *User) ToDomain() _user.Domain {
	return _user.Domain{
		Id_user:         user.Id_user,
		Username:        user.Username,
		Password:        user.Password,
		Email:           user.Email,
		Name:            user.Name,
		SteamProfile_id: user.SteamProfile_id,
		Detail_id:       user.Detail_id,
		Created_at:      user.Created_at,
		Updated_at:      user.Updated_at,
	}
}

func FromDomain(domain user.Domain) User {
	return User{
		Id_user:         domain.Id_user,
		Username:        domain.Username,
		Password:        domain.Password,
		Email:           domain.Email,
		Name:            domain.Name,
		SteamProfile_id: domain.SteamProfile_id,
		Detail_id:       domain.Detail_id,
		Created_at:      domain.Created_at,
		Updated_at:      domain.Updated_at,
	}
}

func ToDomains(u []User) []user.Domain {
	var Domains []user.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}

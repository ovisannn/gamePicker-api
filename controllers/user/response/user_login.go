package response

import (
	"gamePicker/business/user"
	"time"
)

type UserResponse struct {
	Id_user         int
	Username        string
	Password        string
	Email           string
	Name            string
	SteamProfile_id string
	Detail_id       int
	Created_at      time.Time
	Updated_at      time.Time
}

func FromDomain(domain user.Domain) UserResponse {
	return UserResponse{
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

package response

import "gamePicker/business/user"

type UserInsert struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	SteamProfile_id string `json:"steamProfile_id"`
}

func CreateFromDomain(domain user.Domain) UserInsert {
	return UserInsert{
		Username:        domain.Username,
		Email:           domain.Email,
		Name:            domain.Name,
		SteamProfile_id: domain.SteamProfile_id,
	}
}

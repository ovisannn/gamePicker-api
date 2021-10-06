package request

type UserInsert struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	SteamProfile_id string `json:"steamProfile_id"`
}

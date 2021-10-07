package user

import (
	_user "gamePicker/business/user"
	_wallet "gamePicker/business/wallet"
	"time"
)

type User struct {
	Id_user         int `gorm:"primaryKey"`
	Username        string
	Password        string
	Email           string
	Name            string
	SteamProfile_id string `gorm:"steamProfile_id" json:"steamProfile_id"`
	Detail_id       int
	Created_at      time.Time
	Updated_at      time.Time
}

type UserDetail struct {
	Id_userDetail        int
	WishList             WishList_struct
	Wallet_id            int
	Preference_list      Preference_struct
	Inven_id             int
	TopUp_history        TopUp_struct
	IndieTransac_history indieGame_struct
}

type User_inventory struct {
	Id_inven        int
	Game_owned      Game_struct
	IndieGame_owned indieGame_struct
}

type Wallet struct {
	Id_wallet   int
	MoneySaved  int
	MoneyTarget int
	IndieWallet int
}

// wishList
type WishList_struct struct {
	Wishlist []int
}

// preference_list
type Preference_struct struct {
	Preference []int
}

// topup topUp_history
type TopUp_struct struct {
	History []int
}

// indieTransac_history
type IndieTransac_struct struct {
	History []int
}

// game game_owned
type Game_struct struct {
	Game_owned []int
}

// indieGame_owned
type indieGame_struct struct {
	Indie_owned []int
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

func FromDomain(domain _user.Domain) User {
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

func ToDomains(u []User) []_user.Domain {
	var Domains []_user.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}

func FromDomainWallet(domain _wallet.Domain) Wallet {
	return Wallet{
		Id_wallet:   domain.Id_wallet,
		MoneySaved:  domain.MoneySaved,
		MoneyTarget: domain.MoneyTarget,
		IndieWallet: domain.IndieWallet,
	}
}

func (wallet *Wallet) ToDomainWallet() _wallet.Domain {
	return _wallet.Domain{
		Id_wallet:   wallet.Id_wallet,
		MoneySaved:  wallet.MoneySaved,
		MoneyTarget: wallet.MoneyTarget,
		IndieWallet: wallet.IndieWallet,
	}
}

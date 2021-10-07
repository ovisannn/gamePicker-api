package user

import (
	"context"
	"gamePicker/business/user"
	_user "gamePicker/business/user"
	"gamePicker/business/wallet"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) _user.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (DB *MysqlUserRepository) Login(ctx context.Context, email string, password string) (user.Domain, error) {
	var CurrentUser User
	result := DB.Conn.Table("user").First(&CurrentUser, "email = ? AND password = ?", email, password)
	if result.Error != nil {
		return user.Domain{}, result.Error
	}

	return CurrentUser.ToDomain(), nil
}

func (DB *MysqlUserRepository) GetUser(ctx context.Context) ([]user.Domain, error) {
	var CurrentUser []User
	result := DB.Conn.Table("user").Find(&CurrentUser)
	if result.Error != nil {
		return []user.Domain{}, result.Error
	}
	return ToDomains(CurrentUser), nil
}

func (DB *MysqlUserRepository) GetUser_byID(ctx context.Context, user_id int) (user.Domain, error) {
	var CurrentUser User
	result := DB.Conn.Table("user").Where("id_user = ?", user_id).First(&CurrentUser)
	if result.Error != nil {
		return user.Domain{}, result.Error
	}
	return CurrentUser.ToDomain(), nil
}

func (DB *MysqlUserRepository) CreateUser(ctx context.Context, data user.Domain) (user.Domain, error) {
	insertUser := FromDomain(data)
	const (
		raw_user   = `INSERT INTO user (username, password, email, name, steamProfile_id, detail_id, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?)`
		raw_detail = `INSERT INTO userdetail (id_userDetail, wishList, wallet_id, preference_list, inven_id, topUp_history, indieTransac_history) VALUES (?,?,?,?,?,?,?)`
		raw_inven  = `INSERT INTO user_inventory (id_inven, game_owned, indieGame_owned) VALUES(?,?,?)`
		raw_wallet = `INSERT INTO wallet (id_wallet, moneySaved, moneyTarget, indieWallet) VALUES (?,?,?,?)`
		raw_update = `UPDATE USER SET detail_id = ? WHERE id_user = ?`
	)

	result := DB.Conn.Exec(raw_user, insertUser.Username, insertUser.Password, insertUser.Email, insertUser.Name, insertUser.SteamProfile_id, insertUser.Detail_id, insertUser.Created_at, insertUser.Updated_at)
	var getMax int
	maxRaw := DB.Conn.Table("user").Select("max(id_user)").Row()
	maxRaw.Scan(&getMax)
	DB.Conn.Exec(raw_detail, getMax, "", getMax, "", getMax, "", "")
	DB.Conn.Exec(raw_inven, getMax, "", "")
	DB.Conn.Exec(raw_wallet, getMax, 0, 0, 0)
	DB.Conn.Exec(raw_update, getMax, getMax)
	// fmt.Println(getBiggest_id)
	if result.Error != nil {
		return user.Domain{}, result.Error
	}
	return insertUser.ToDomain(), nil
}

func (DB *MysqlUserRepository) UpdateMoneySaved(ctx context.Context, data wallet.Domain, id int) (wallet.Domain, error) {
	const raw_UpdateWallet = `UPDATE wallet SET moneySaved = ? WHERE id_wallet = ?`
	NewWallet := FromDomainWallet(data)
	result := DB.Conn.Raw(raw_UpdateWallet, NewWallet.MoneySaved, id).Scan(&NewWallet)
	if result.Error != nil {
		return wallet.Domain{}, result.Error
	}
	return NewWallet.ToDomainWallet(), nil
}
func (DB *MysqlUserRepository) UpdateMoneyTarget() {}
func (DB *MysqlUserRepository) UpdateMoney()       {}

// get wallet

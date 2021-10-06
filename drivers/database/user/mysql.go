package user

import (
	"context"
	"gamePicker/business/user"
	_user "gamePicker/business/user"

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
	result := DB.Conn.Table("user").Where("user_id = ?", user_id).First(&CurrentUser)
	if result.Error != nil {
		return user.Domain{}, result.Error
	}
	return CurrentUser.ToDomain(), nil
}

func (DB *MysqlUserRepository) CreateUser(ctx context.Context, data user.Domain) (user.Domain, error) {
	insertUser := FromDomain(data)
	// fmt.Println(insertUser.SteamProfile_id)
	const raw = `INSERT INTO user (username,password,email,name,steamProfile_id,detail_id,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?)`
	// result := DB.Conn.Table("user").Select("user_id", "username", "password", "email", "name", "steamProfile_id", "created_at", "updated_at").Create(&insertUser)
	result := DB.Conn.Exec(raw, insertUser.Username, insertUser.Password, insertUser.Email, insertUser.Name, insertUser.SteamProfile_id, insertUser.Detail_id, insertUser.Created_at, insertUser.Updated_at)
	if result.Error != nil {
		return user.Domain{}, result.Error
	}
	return insertUser.ToDomain(), nil
}

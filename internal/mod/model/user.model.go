package model

import (
	"asyncfiber/internal/types"
	"asyncfiber/pkg/database"
	"gorm.io/gorm"
)

type UsersRepo struct {
	data *types.Users
	conn *gorm.DB
}

func NewUser() types.IUser {
	conn, err := database.Connection()
	if err != nil {
		panic(err.Error())
	}
	return &UsersRepo{
		data: &types.Users{},
		conn: conn,
	}
}

func (user *UsersRepo) GetByEmail(email string) (*types.Users, error) {
	if err := user.conn.Raw("SELECT * FROM users WHERE email = ?", email).Scan(user.data).Error; err != nil {
		return nil, err
	}
	return user.data, nil
}

func (user *UsersRepo) GetByID(Id string) (*types.Users, error) {
	conn, err := database.Connection()
	if err != nil {
		return nil, err
	}
	// if err := conn.Raw("SELECT * FROM users WHERE id = ?", Id).Scan(_user).Error; err != nil {
	// 	return nil, err
	// }
	if err := conn.Table("users").Where("id=?", Id).First(user.data).Error; err != nil {
		return nil, err
	}
	return user.data, nil
}

func (user *UsersRepo) Insert(_user *types.Users) error {
	if _user == nil {
		return nil
	}
	if err := user.conn.Exec(
		"INSERT INTO users (id, first_name, last_name, email, phone_number, role, title, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		_user.Id, _user.FirstName, _user.LastName, _user.Email, _user.PhoneNumber, _user.Role, _user.Title, _user.Password).Error; err != nil {
		return err
	}
	return nil
}

func (user *UsersRepo) Empty() bool {
	return user.data.Email == ""
}

func (user *UsersRepo) PromoteAdmin(id, role, password, email, phoneNumber string) error {
	if err := user.conn.Exec(
		"UPDATE users SET password = ?, role=?, email = ?, phone_number = ? WHERE id = ?;",
		password, role, email, phoneNumber, id).Error; err != nil {
		return err
	}
	return nil
}

func (user *UsersRepo) GetPassword() string {
	return user.data.Password
}

func (user *UsersRepo) GetEmail() string {
	return user.data.Email
}

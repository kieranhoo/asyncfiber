package model

import "qrcheckin/pkg/database"

type Users struct {
	Id          string `json:"id" gorm:"column:id"`
	FirstName   string `json:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" gorm:"column:last_name"`
	Email       string `json:"email" gorm:"column:email"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Role        string `json:"role" gorm:"column:role"`
	Title       string `json:"title" gorm:"column:title"`
	Password    string `json:"password" gorm:"column:password"`
}

func (user *Users) GetByEmail(email string) (*Users, error) {
	_user := new(Users)
	conn, err := database.Connection()
	if err != nil {
		return nil, err
	}
	if err := conn.Raw("SELECT * FROM users WHERE email = ?", email).Scan(_user).Error; err != nil {
		return nil, err
	}
	return _user, nil
}

func (user *Users) GetByID(Id string) (*Users, error) {
	_user := new(Users)
	conn, err := database.Connection()
	if err != nil {
		return nil, err
	}
	// if err := conn.Raw("SELECT * FROM users WHERE id = ?", Id).Scan(_user).Error; err != nil {
	// 	return nil, err
	// }
	if err := conn.Table("users").Where("id=?", Id).First(_user).Error; err != nil {
		return nil, err
	}
	return _user, nil
}

func (user *Users) Insert(_user *Users) error {
	conn, err := database.Connection()
	if err != nil {
		return err
	}
	if err := conn.Exec(
		"INSERT INTO users (id, first_name, last_name, email, phone_number, role, title, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		_user.Id, _user.FirstName, _user.LastName, _user.Email, _user.PhoneNumber, _user.Role, _user.Title, _user.Password).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) Empty() bool {
	return user.Email == ""
}

func (user *Users) PromoteAdmin(id, role, password, email, phoneNumber string) error {
	conn, err := database.Connection()
	if err != nil {
		return err
	}
	if err := conn.Exec(
		"UPDATE users SET password = ?, role=?, email = ?, phone_number = ? WHERE id = ?;",
		password, role, email, phoneNumber, id).Error; err != nil {
		return err
	}
	return nil
}

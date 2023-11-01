package model

type IUser interface {
	GetByEmail(email string) (*Users, error)
	PromoteAdmin(id, role, password, email, phoneNumber string) error
	GetByID(Id string) (*Users, error)
	Insert(u *Users) error
	Empty() bool
	GetPassword() string
	GetEmail() string
}

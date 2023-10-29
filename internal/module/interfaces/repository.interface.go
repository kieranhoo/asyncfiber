package interfaces

import "asyncfiber/internal/module/entity"

type IUser interface {
	GetByEmail(email string) (*entity.Users, error)
	PromoteAdmin(id, role, password, email, phoneNumber string) error
	GetByID(Id string) (*entity.Users, error)
	Insert(u *entity.Users) error
	Empty() bool
	GetPassword() string
	GetEmail() string
}

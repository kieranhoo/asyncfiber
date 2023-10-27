package interfaces

type IUser interface {
	GetByEmail(email string) (IUser, error)
	PromoteAdmin(id, role, password, email, phoneNumber string) error
	GetByID(Id string) (IUser, error)
	Insert(u IUser) error
	Empty() bool
	GetPassword() string
	GetEmail() string
}

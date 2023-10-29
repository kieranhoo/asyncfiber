package entity

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

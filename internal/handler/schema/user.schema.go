package schema

type SignUpRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Id          string `json:"id" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type SignUpResponse struct {
	Success bool   `json:"success" validate:"required"`
	Msg     string `json:"msg" validate:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Success bool   `json:"success" validate:"required"`
	Token   string `json:"token" validate:"required"`
	Email   string `json:"email" validate:"required"`
}

type UserResponse struct {
	Id          string `json:"id" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

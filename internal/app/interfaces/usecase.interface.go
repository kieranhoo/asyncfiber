package interfaces

import "asyncfiber/internal/app/schema"

type IAuth interface {
	SignUp(req schema.SignUpRequest) error
	SignIn(req schema.SignInRequest) (string, error)
}

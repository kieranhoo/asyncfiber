package types

import "asyncfiber/internal/mod/schema"

type IAuth interface {
	SignUp(req schema.SignUpRequest) error
	SignIn(req schema.SignInRequest) (string, error)
}

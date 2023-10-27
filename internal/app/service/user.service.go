package service

import (
	"asyncfiber/internal/app/interfaces"
	"asyncfiber/internal/app/model"
	"asyncfiber/internal/app/schema"
	"asyncfiber/internal/app/tasks"
	"asyncfiber/pkg/utils"
)

type Auth struct {
	repo interfaces.IUser
}

func NewAuth() interfaces.IAuth {
	return &Auth{
		repo: model.NewUser(),
	}
}

func (auth *Auth) SignUp(req schema.SignUpRequest) error {
	_pass, err := utils.GenPassword(req.Password)
	if err != nil {
		return err
	}
	return tasks.SignUp(
		req.Id,
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
		req.Email,
		string(_pass),
	)
}

func (auth *Auth) SignIn(req schema.SignInRequest) (string, error) {
	_user, err := auth.repo.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := utils.ComparePassword(_user.GetPassword(), req.Password); err != nil {
		return "", err
	}
	token, err := utils.GenerateToken(_user.GetEmail())
	if err != nil {
		return "", err
	}
	return token, nil
}

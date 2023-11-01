package service

import (
	"asyncfiber/internal/config"
	"asyncfiber/internal/module/model"
	"asyncfiber/internal/module/schema"
	"asyncfiber/internal/module/tasks"
	"asyncfiber/pkg/utils"
	"asyncfiber/pkg/x/worker"
	"errors"
)

type Auth struct {
	repo model.IUser
}

func NewAuth() IAuth {
	return &Auth{
		repo: model.NewUser(),
	}
}

func (auth *Auth) SignUp(req schema.SignUpRequest) error {
	_pass, err := utils.GenPassword(req.Password)
	if err != nil {
		return err
	}
	_user, err := auth.repo.GetByID(req.Id)
	if err != nil || _user == nil {
		return worker.Exec(config.CriticalQueue, worker.NewTask(
			tasks.WorkerSaveUser,
			model.Users{
				Id:          req.Id,
				FirstName:   req.FirstName,
				LastName:    req.LastName,
				PhoneNumber: req.PhoneNumber,
				Email:       req.Email,
				Password:    string(_pass),
			},
		))
	}
	if _user.Password == "" {
		return auth.repo.PromoteAdmin(req.Id, "admin", string(_pass), req.Email, req.PhoneNumber)
	}
	return errors.New("user already exists")
}

func (auth *Auth) SignIn(req schema.SignInRequest) (string, error) {
	_user, err := auth.repo.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := utils.ComparePassword(_user.Password, req.Password); err != nil {
		return "", err
	}
	token, err := utils.GenerateToken(_user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

package service

import (
	"qrcheckin/internal/handler/model"
	"qrcheckin/internal/handler/schema"
	"qrcheckin/internal/handler/tasks"
	"qrcheckin/pkg/utils"
)

func SignUp(req schema.SignUpRequest) error {
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

func SignIn(req schema.SignInRequest) (string, error) {
	user := new(model.Users)
	_user, err := user.GetByEmail(req.Email)
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

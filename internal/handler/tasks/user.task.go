package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"qrcheckin/internal/handler/model"

	"github.com/hibiken/asynq"
)

func SignUp(id, firstName, lastName, phoneNumber, email, password string) error {
	_user, err := new(model.Users).GetByID(id)
	if err != nil {
		return err
	}
	if _user.Empty() {
		return _user.Insert(&model.Users{
			Id:          id,
			FirstName:   firstName,
			LastName:    lastName,
			PhoneNumber: phoneNumber,
			Email:       email,
			Password:    password,
		})
	}
	if _user.Password == "" {
		return _user.PromoteAdmin(id, "admin", password, email, phoneNumber)
	}
	return errors.New("user already exists")
}



func SaveUser(id, firstName, lastName, phoneNumber, email string) error {
	return new(model.Users).Insert(&model.Users{
		Id:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	})
}

func HandleSaveUser(c context.Context, task *asynq.Task) error {
	var user model.Users
	if err := json.Unmarshal(task.Payload(), &user); err != nil {
		return err
	}
	return new(model.Users).Insert(&user)
}

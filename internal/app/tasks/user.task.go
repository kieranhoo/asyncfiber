package tasks

import (
	"asyncfiber/internal/app/model"
	"asyncfiber/internal/config"
	"asyncfiber/internal/worker"
	"context"
	"encoding/json"
	"errors"

	"github.com/hibiken/asynq"
)

var user = model.NewUser()

func SignUp(id, firstName, lastName, phoneNumber, email, password string) error {
	_user, err := user.GetByID(id)
	if err != nil || _user == nil {
		// return user.Insert(&model.Users{
		// 	Id:          id,
		// 	FirstName:   firstName,
		// 	LastName:    lastName,
		// 	PhoneNumber: phoneNumber,
		// 	Email:       email,
		// 	Password:    password,
		// })
		return worker.Exec(config.CriticalQueue, worker.NewTask(
			WorkerSaveUser,
			model.Users{
				Id:          id,
				FirstName:   firstName,
				LastName:    lastName,
				PhoneNumber: phoneNumber,
				Email:       email,
				Password:    password,
			},
		))
	}
	if _user.GetPassword() == "" {
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

func HandleSaveUser(_ context.Context, task *asynq.Task) error {
	var _user model.Users
	if err := json.Unmarshal(task.Payload(), &_user); err != nil {
		return err
	}
	return user.Insert(&_user)
}

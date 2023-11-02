package tasks

import (
	"asyncfiber/internal/config"
	"asyncfiber/internal/mod/model"
	"asyncfiber/internal/types"
	"asyncfiber/pkg/x/worker"
	"context"
	"encoding/json"
	"errors"

	"github.com/hibiken/asynq"
)

var user = model.NewUser()

func SignUp(id, firstName, lastName, phoneNumber, email, password string) error {
	_user, err := user.GetByID(id)
	if err != nil || _user == nil {
		return worker.Exec(config.CriticalQueue, worker.NewTask(
			WorkerSaveUser,
			types.Users{
				Id:          id,
				FirstName:   firstName,
				LastName:    lastName,
				PhoneNumber: phoneNumber,
				Email:       email,
				Password:    password,
			},
		))
	}
	if _user.Password == "" {
		return user.PromoteAdmin(id, "admin", password, email, phoneNumber)
	}
	return errors.New("user already exists")
}

func SaveUser(id, firstName, lastName, phoneNumber, email string) error {
	return new(model.UsersRepo).Insert(&types.Users{
		Id:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	})
}

func HandleSaveUser(_ context.Context, task *asynq.Task) error {
	var _user types.Users
	if err := json.Unmarshal(task.Payload(), &_user); err != nil {
		return err
	}
	return user.Insert(&_user)
}

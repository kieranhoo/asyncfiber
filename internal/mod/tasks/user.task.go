package tasks

import (
	"asyncfiber/internal/mod/model"
	"asyncfiber/internal/types"
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

var user = model.NewUser()

func HandleSaveUser(_ context.Context, task *asynq.Task) error {
	var _user types.Users
	if err := json.Unmarshal(task.Payload(), &_user); err != nil {
		return err
	}
	return user.Insert(&_user)
}

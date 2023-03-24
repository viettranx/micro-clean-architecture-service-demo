package repository

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/pkg/errors"
	"strings"
)

func (repo *repository) GetTaskById(ctx context.Context, id int, extras ...string) (*entity.Task, error) {
	data, err := repo.taskStore.GetTaskById(ctx, id)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	wantedUserInfo := false
	for _, k := range extras {
		if strings.ToLower(k) == "user" {
			wantedUserInfo = true
		}
	}

	// Get extra infos: User
	if wantedUserInfo {
		user, err := repo.userStore.GetUserById(ctx, data.UserId)

		if err != nil {
			return nil, errors.Wrap(err, "cannot get more info: user")
		}

		data.User = user
	}

	return data, nil
}

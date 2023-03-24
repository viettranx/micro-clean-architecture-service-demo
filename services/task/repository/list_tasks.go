package repository

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/pkg/errors"
	"github.com/viettranx/service-context/core"
	"strings"
)

func (repo *repository) ListTasks(ctx context.Context, filter *entity.Filter, paging *core.Paging, extras ...string) ([]entity.Task, error) {
	tasks, err := repo.taskStore.ListTasks(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	wantedUserInfo := false
	for _, k := range extras {
		if strings.ToLower(k) == "user" {
			wantedUserInfo = true
		}
	}

	// Get extra infos: User
	if wantedUserInfo {
		userIds := make([]int, len(tasks))

		for i := range userIds {
			userIds[i] = tasks[i].UserId
		}

		users, err := repo.userStore.GetUsersByIds(ctx, userIds)

		if err != nil {
			return nil, errors.Wrap(err, "cannot get more info: user")
		}

		// For speed up mapping data
		userMap := make(map[int]*core.SimpleUser)

		for i, u := range users {
			userMap[u.Id] = &users[i]
		}

		for i, t := range tasks {
			tasks[i].User = userMap[t.UserId]
		}
	}

	return tasks, nil
}

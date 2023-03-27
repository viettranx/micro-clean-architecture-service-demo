package business

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/viettranx/service-context/core"
)

func (biz *business) ListTasks(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Task, error) {
	tasks, err := biz.taskRepo.ListTasks(ctx, filter, paging)

	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListTask.Error()).
			WithDebug(err.Error())
	}

	// Get extra infos: User
	userIds := make([]int, len(tasks))

	for i := range userIds {
		userIds[i] = tasks[i].UserId
	}

	users, err := biz.userRepo.GetUsersByIds(ctx, userIds)

	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListTask.Error()).
			WithDebug(err.Error())
	}

	// For speed up mapping data
	userMap := make(map[int]*core.SimpleUser)

	for i, u := range users {
		userMap[u.Id] = &users[i]
	}

	for i, t := range tasks {
		tasks[i].User = userMap[t.UserId]
	}

	return tasks, nil
}

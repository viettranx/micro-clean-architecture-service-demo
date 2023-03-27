package business

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/viettranx/service-context/core"
)

func (biz *business) GetTaskById(ctx context.Context, id int) (*entity.Task, error) {
	data, err := biz.taskRepo.GetTaskById(ctx, id)

	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetTask.Error()).
			WithDebug(err.Error())
	}

	if data.Status == entity.StatusDeleted {
		return nil, core.ErrNotFound.WithError(entity.ErrTaskNotFound.Error())
	}

	// Get extra infos: User
	user, err := biz.userRepo.GetUserById(ctx, data.UserId)

	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetTask.Error()).
			WithDebug(err.Error())
	}

	data.User = user

	return data, nil
}

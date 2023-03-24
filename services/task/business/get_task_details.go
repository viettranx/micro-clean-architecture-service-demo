package business

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/viettranx/service-context/core"
)

func (biz *business) GetTaskDetails(ctx context.Context, id int, extras ...string) (*entity.Task, error) {
	// Get full Task data and its associations from repository
	data, err := biz.repository.GetTaskById(ctx, id, extras...)

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

	return data, nil
}

package business

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/viettranx/service-context/core"
)

func (biz *business) UpdateTask(ctx context.Context, id int, data *entity.TaskDataUpdate) error {
	// Get task data, without extra infos
	task, err := biz.taskRepo.GetTaskById(ctx, id)

	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetTask.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetTask.Error()).
			WithDebug(err.Error())
	}

	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	// Only task owner can do this
	if requesterId != task.UserId {
		return core.ErrForbidden.WithError(entity.ErrRequesterIsNotOwner.Error())
	}

	// Only update task with doing status
	if task.Status != entity.StatusDoing {
		return core.ErrForbidden.
			WithError(entity.ErrCannotUpdateTask.Error()).
			WithReason("Only update task with doing status")
	}

	if err := biz.taskRepo.UpdateTask(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotUpdateTask.Error()).
			WithDebug(err.Error())
	}

	return nil
}

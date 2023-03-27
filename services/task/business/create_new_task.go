package business

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/viettranx/service-context/core"
)

func (biz *business) CreateNewTask(ctx context.Context, data *entity.TaskDataCreation) error {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID()) // task owner id, id of who creates this new task

	data.Prepare(requesterId, entity.StatusDoing)

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.taskRepo.AddNewTask(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateTask.Error())
	}

	return nil
}

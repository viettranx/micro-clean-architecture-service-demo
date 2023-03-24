package business

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/viettranx/service-context/core"
)

func (biz *business) ListTasks(ctx context.Context, filter *entity.Filter, paging *core.Paging, extras ...string) ([]entity.Task, error) {
	// Get full Task data and its associations from repository
	data, err := biz.repository.ListTasks(ctx, filter, paging, extras...)

	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListTask.Error()).
			WithDebug(err.Error())
	}

	return data, nil
}

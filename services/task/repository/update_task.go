package repository

import (
	"context"
	"demo-service/services/task/entity"
)

func (repo *repository) UpdateTask(ctx context.Context, id int, data *entity.TaskDataUpdate) error {
	// nothing to do, just delegate to storage
	return repo.taskStore.UpdateTask(ctx, id, data)
}

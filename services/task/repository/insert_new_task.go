package repository

import (
	"context"
	"demo-service/services/task/entity"
)

func (repo *repository) AddNewTask(ctx context.Context, data *entity.TaskDataCreation) error {
	// Nothing to do here, just delegate to storage
	return repo.taskStore.InsertTask(ctx, data)
}

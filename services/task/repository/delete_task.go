package repository

import (
	"context"
)

func (repo *repository) DeleteTask(ctx context.Context, id int) error {
	// nothing to do, just delegate to storage
	return repo.taskStore.DeleteTask(ctx, id)
}

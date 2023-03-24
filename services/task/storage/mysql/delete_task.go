package mysql

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/pkg/errors"
)

func (store *mysqlStore) DeleteTask(ctx context.Context, id int) error {
	// Soft delete
	if err := store.db.Table(entity.Task{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": entity.StatusDeleted,
		}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

package mysql

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/pkg/errors"
)

func (store *mysqlStore) InsertTask(ctx context.Context, data *entity.TaskDataCreation) error {
	if err := store.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

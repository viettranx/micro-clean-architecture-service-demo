package mysql

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/pkg/errors"
)

func (repo *mysqlRepo) AddNewTask(ctx context.Context, data *entity.TaskDataCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

package mysql

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/pkg/errors"
	"github.com/viettranx/service-context/core"
	"gorm.io/gorm"
)

func (repo *mysqlRepo) GetTaskById(ctx context.Context, id int) (*entity.Task, error) {
	var data entity.Task

	if err := repo.db.
		Table(data.TableName()).
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.ErrRecordNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &data, nil
}

package mysql

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/pkg/errors"
	"github.com/viettranx/service-context/core"
)

func (repo *mysqlRepo) ListTasks(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Task, error) {

	var tasks []entity.Task

	db := repo.db.
		Table(entity.Task{}.TableName()).
		Where("status <> ?", entity.StatusDeleted)

	if userId := filter.UserId; userId != nil {
		uid, err := core.FromBase58(*userId)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		db = db.Where("user_id = ?", uid.GetLocalID())
	}

	if status := filter.Status; status != nil {
		db = db.Where("status = ?", *status)
	}

	// Count total records match conditions
	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	// Query data with paging
	if err := db.Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&tasks).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return tasks, nil
}

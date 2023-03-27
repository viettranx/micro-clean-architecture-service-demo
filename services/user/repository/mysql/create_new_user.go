package mysql

import (
	"context"
	"demo-service/services/user/entity"
	"github.com/pkg/errors"
)

func (repo *mysqlRepo) CreateNewUser(ctx context.Context, data *entity.UserDataCreation) error {
	if err := repo.db.Table(data.TableName()).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

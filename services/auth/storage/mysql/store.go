package mysql

import (
	"context"
	"demo-service/services/auth/entity"
	"github.com/pkg/errors"
	"github.com/viettranx/service-context/core"
	"gorm.io/gorm"
)

type mysqlStore struct {
	db *gorm.DB
}

func NewMySQLStore(db *gorm.DB) *mysqlStore {
	return &mysqlStore{db: db}
}

func (store *mysqlStore) AddNewAuth(ctx context.Context, data *entity.Auth) error {
	if err := store.db.Table(data.TableName()).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (store *mysqlStore) GetAuth(ctx context.Context, email string) (*entity.Auth, error) {
	var data entity.Auth

	if err := store.db.
		Table(data.TableName()).
		Where("email = ?", email).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.ErrRecordNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &data, nil
}

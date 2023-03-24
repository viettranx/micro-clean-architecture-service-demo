package mysql

import "gorm.io/gorm"

type mysqlStore struct {
	db *gorm.DB
}

func NewMySQLStore(db *gorm.DB) *mysqlStore {
	return &mysqlStore{db: db}
}

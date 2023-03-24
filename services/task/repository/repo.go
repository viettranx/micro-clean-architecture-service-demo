package repository

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/viettranx/service-context/core"
)

type TaskStorage interface {
	ListTasks(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Task, error)
	GetTaskById(ctx context.Context, id int) (*entity.Task, error)
	InsertTask(ctx context.Context, data *entity.TaskDataCreation) error
	UpdateTask(ctx context.Context, id int, data *entity.TaskDataUpdate) error
	DeleteTask(ctx context.Context, id int) error
}

// UserStorage is UserService in microservices,
// the different is we can implement with REST-RPC or gRPC
// instead of a MySQL storage
type UserStorage interface {
	GetUsersByIds(ctx context.Context, ids []int) ([]core.SimpleUser, error)
	GetUserById(ctx context.Context, id int) (*core.SimpleUser, error)
}

type repository struct {
	taskStore TaskStorage
	userStore UserStorage
}

func NewRepository(taskStore TaskStorage, userStore UserStorage) *repository {
	return &repository{taskStore: taskStore, userStore: userStore}
}

package business

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/viettranx/service-context/core"
)

type Repository interface {
	AddNewTask(ctx context.Context, data *entity.TaskDataCreation) error
	UpdateTask(ctx context.Context, id int, data *entity.TaskDataUpdate) error
	DeleteTask(ctx context.Context, id int) error
	GetTaskById(ctx context.Context, id int, extras ...string) (*entity.Task, error)
	ListTasks(ctx context.Context, filter *entity.Filter, paging *core.Paging, extras ...string) ([]entity.Task, error)
}

type business struct {
	repository Repository
}

func NewBusiness(repository Repository) *business {
	return &business{
		repository: repository,
	}
}

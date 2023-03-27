package api

import (
	"context"
	"demo-service/services/task/entity"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	CreateNewTask(ctx context.Context, data *entity.TaskDataCreation) error
	GetTaskById(ctx context.Context, id int) (*entity.Task, error)
	ListTasks(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Task, error)
	UpdateTask(ctx context.Context, id int, data *entity.TaskDataUpdate) error
	DeleteTask(ctx context.Context, id int) error
	//SetTaskToDone(ctx context.Context, id int) error
	//SetTaskToDoing(ctx context.Context, id int) error
}

type api struct {
	serviceCtx sctx.ServiceContext
	business   Business
}

func NewAPI(serviceCtx sctx.ServiceContext, business Business) *api {
	return &api{
		serviceCtx: serviceCtx,
		business:   business,
	}
}

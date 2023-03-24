package api

import (
	"demo-service/common"
	"demo-service/services/task/entity"
	"github.com/gin-gonic/gin"
	"github.com/viettranx/service-context/core"
	"net/http"
)

func (api *api) CreateTaskHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.TaskDataCreation

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		// Set requester to context
		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		// we can set user_id directly to data creation, but I don't recommend it
		// uid, _ := core.FromBase58(requester.GetSubject())
		// data.UserId = int(uid.GetLocalID())

		if err := api.business.CreateNewTask(ctx, &data); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		data.Mask()
		c.JSON(http.StatusOK, core.ResponseData(data.FakeId))
	}
}

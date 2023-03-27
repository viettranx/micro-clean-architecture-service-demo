package api

import (
	"demo-service/common"
	"demo-service/services/task/entity"
	"github.com/gin-gonic/gin"
	"github.com/viettranx/service-context/core"
	"net/http"
)

func (api *api) ListTaskHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		type reqParam struct {
			entity.Filter
			core.Paging
		}

		var rp reqParam

		if err := c.ShouldBind(&rp); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		rp.Paging.Process()

		tasks, err := api.business.ListTasks(c.Request.Context(), &rp.Filter, &rp.Paging)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		for i := range tasks {
			tasks[i].Mask()
		}

		c.JSON(http.StatusOK, core.SuccessResponse(tasks, rp.Paging, rp.Filter))
	}
}

package api

import (
	"demo-service/common"
	"github.com/gin-gonic/gin"
	"github.com/viettranx/service-context/core"
	"net/http"
)

func (api *api) DeleteTaskHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("task-id"))

		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		// Set requester to context
		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.DeleteTask(ctx, int(uid.GetLocalID())); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

package api

import (
	"context"
	"demo-service/common"
	"demo-service/services/user/entity"
	"github.com/gin-gonic/gin"
	"github.com/viettranx/service-context/core"
	"net/http"
)

type Business interface {
	GetUserProfile(ctx context.Context) (*entity.User, error)
}

type api struct {
	business Business
}

func NewAPI(business Business) *api {
	return &api{business: business}
}

func (api *api) GetUserProfileHdl() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Set requester to context
		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		user, err := api.business.GetUserProfile(ctx)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		user.Mask()

		c.JSON(http.StatusOK, core.ResponseData(user))
	}
}

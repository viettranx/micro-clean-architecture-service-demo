package api

import (
	"context"
	"demo-service/common"
	"demo-service/services/auth/entity"
	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
	"net/http"
)

type Business interface {
	Login(ctx context.Context, data *entity.AuthEmailPassword) (*entity.TokenResponse, error)
	Register(ctx context.Context, data *entity.AuthRegister) error
}

type api struct {
	serviceCtx sctx.ServiceContext
	business   Business
}

func NewAPI(serviceCtx sctx.ServiceContext, business Business) *api {
	return &api{serviceCtx: serviceCtx, business: business}
}

func (api *api) LoginHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.AuthEmailPassword

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		response, err := api.business.Login(c.Request.Context(), &data)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(response))
	}
}

func (api *api) RegisterHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.AuthRegister

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		err := api.business.Register(c.Request.Context(), &data)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

package common

import (
	"github.com/gin-gonic/gin"
	"github.com/viettranx/service-context/core"
	"net/http"
)

func WriteErrorResponse(c *gin.Context, err error) {
	if errSt, ok := err.(core.StatusCodeCarrier); ok {
		c.JSON(errSt.StatusCode(), errSt)
		return
	}

	c.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithError(err.Error()))
}

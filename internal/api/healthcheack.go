package api

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	HealthCheckService interfaces.IHealthCheckService
}

func (api *HealthCheck) HealthCheckHandleHttp(ctx *gin.Context) {
	msg, err := api.HealthCheckService.HealthCheckServices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	helpers.SendResponse(ctx, 200, msg, nil)
}

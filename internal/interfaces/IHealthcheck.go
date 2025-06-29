package interfaces

import "github.com/gin-gonic/gin"

type IHealthCheckService interface {
	HealthCheckServices() (string, error)
}

type IHealthCheckHandler interface {
	HealthCheckHandleHttp(ctx *gin.Context)
}

type IHealthCheckRepository interface {
}

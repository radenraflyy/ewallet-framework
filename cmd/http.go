package cmd

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/api"
	"ewallet-framework/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healthCheckSvc := &services.Healthcheck{}
	healthCheckApi := api.HealthCheck{
		HealthCheckService: healthCheckSvc,
	}

	r := gin.Default()
	r.GET("/health", healthCheckApi.HealthCheckHandleHttp)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080")) // listen and serve on
	if err != nil {
		log.Fatal(err)
	}
}

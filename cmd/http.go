package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healthCheckSvc := &services.Healthcheck{}
	healthCheckApi := api.HealthCheck{
		HealthCheckService: healthCheckSvc,
	}

	registerRepo := &repository.RegisterRepository{
		DB: helpers.DB,
	}
	registerSvc := &services.RegisterService{
		RegisterRepository: registerRepo,
	}
	registerApi := api.RegisterUserHandler{
		RegisterService: registerSvc,
	}

	r := gin.Default()
	r.GET("/health", healthCheckApi.HealthCheckHandleHttp)

	userV1 := r.Group("/v1/users")
	userV1.POST("/register", registerApi.RegisterUserHandlerHttp)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080")) // listen and serve on
	if err != nil {
		log.Fatal(err)
	}
}

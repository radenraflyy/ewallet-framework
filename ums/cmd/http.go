package cmd

import (
	"log"
	"ums/helpers"
	"ums/internal/api"
	"ums/internal/interfaces"
	"ums/internal/repository"
	"ums/internal/services"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	depedency := depedencyInject()

	r := gin.Default()
	r.GET("/health", depedency.HealthCheckApi.HealthCheckHandleHttp)

	userV1 := r.Group("/v1/users")
	userV1.POST("/register", depedency.RegisterApi.RegisterUserHandlerHttp)
	userV1.POST("/login", depedency.LoginApi.LoginUserHandlerHttp)
	userV1.DELETE("/logout", depedency.AuthMiddleware, depedency.LogoutApi.LogoutHandler)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080")) // listen and serve on
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	UserRepository interfaces.IUserRepository
	HealthCheckApi interfaces.IHealthCheckHandler
	RegisterApi    interfaces.IRegisterUserHandler
	LoginApi       interfaces.ILoginHandler
	LogoutApi      interfaces.ILogoutHandler

	TokenValidationApi *api.TokenValidationHandler
}

func depedencyInject() Dependency {
	healthCheckSvc := &services.Healthcheck{}
	healthCheckApi := &api.HealthCheck{
		HealthCheckService: healthCheckSvc,
	}

	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}
	registerSvc := &services.RegisterService{
		UserRepository: userRepo,
	}
	registerApi := &api.RegisterUserHandler{
		RegisterService: registerSvc,
	}

	loginSvc := &services.LoginService{
		UserRepo: userRepo,
	}
	loginApi := &api.LoginHandler{
		LoginService: loginSvc,
	}

	logoutSvc := &services.LogoutService{
		UserRepo: userRepo,
	}
	logoutApi := &api.LogoutHandler{
		LogoutService: logoutSvc,
	}

	// tokenValidationAPI := api.TokenValidationHandler{
	// 	TokenValidationService: &services.TokenValidationService{
	// 		UserRepository: userRepo,
	// 	},
	// }

	tokenValidationSvc := &services.TokenValidationService{
		UserRepository: userRepo,
	}
	tokenValidationAPI := &api.TokenValidationHandler{
		TokenValidationService: tokenValidationSvc,
	}

	return Dependency{
		UserRepository:     userRepo,
		HealthCheckApi:     healthCheckApi,
		RegisterApi:        registerApi,
		LoginApi:           loginApi,
		LogoutApi:          logoutApi,
		TokenValidationApi: tokenValidationAPI,
	}
}

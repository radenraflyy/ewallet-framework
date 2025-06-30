package interfaces

import (
	"context"
	"ums/internal/models"

	"github.com/gin-gonic/gin"
)

type IRegisterRepository interface {
	RegisterUser(ctx context.Context, req *models.User) (interface{}, error)
}

type IRegisterUserHandler interface {
	RegisterUserHandlerHttp(c *gin.Context)
}

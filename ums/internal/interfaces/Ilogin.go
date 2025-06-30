package interfaces

import (
	"context"
	"ums/internal/models"

	"github.com/gin-gonic/gin"
)

type ILoginService interface {
	Login(c context.Context, req models.UserLogin) (models.UserLoginResponse, error)
}

type ILoginHandler interface {
	LoginUserHandlerHttp(c *gin.Context)
}

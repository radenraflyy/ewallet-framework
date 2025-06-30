package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ILogoutService interface {
	Logout(c context.Context, token string, userId uint) error
}

type ILogoutHandler interface {
	LogoutHandler(c *gin.Context)
}

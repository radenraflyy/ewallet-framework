package api

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (l *LogoutHandler) LogoutHandler(c *gin.Context) {
	var log = helpers.Logger

	token := c.GetHeader("Authorization")
	userId := c.MustGet("userId").(uint)
	err := l.LogoutService.Logout(c.Request.Context(), token, userId)
	if err != nil {
		log.Error("Failed to logout user: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	helpers.SendResponse(c, http.StatusOK, "Logout successful", nil)
}

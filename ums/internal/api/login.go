package api

import (
	"net/http"
	"ums/constant"
	"ums/helpers"
	"ums/internal/interfaces"
	"ums/internal/models"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (api *LoginHandler) LoginUserHandlerHttp(c *gin.Context) {
	var (
		log  = helpers.Logger
		req  models.UserLogin
		resp models.UserLoginResponse
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Info("Failed to parse request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constant.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Info("Validation failed: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constant.ErrFailedBadRequest, nil)
		return
	}

	resp, err := api.LoginService.Login(c.Request.Context(), req)
	if err != nil {
		log.Error("Failed to login user: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constant.ErrInternalServerError, nil)
		return
	}
	helpers.SendResponse(c, http.StatusOK, constant.SuccessMessage, resp)
}

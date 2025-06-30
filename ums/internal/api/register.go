package api

import (
	"ums/constant"
	"ums/helpers"
	"ums/internal/interfaces"
	"ums/internal/models"

	"github.com/gin-gonic/gin"
)

type RegisterUserHandler struct {
	RegisterService interfaces.IRegisterRepository
}

func (r *RegisterUserHandler) RegisterUserHandlerHttp(c *gin.Context) {
	var log = helpers.Logger
	req := models.User{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Info("Failed to parse request body: ", err)
		helpers.SendResponse(c, 400, constant.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Info("Validation failed: ", err)
		helpers.SendResponse(c, 400, constant.ErrFailedBadRequest, nil)
		return
	}

	resp, err := r.RegisterService.RegisterUser(c.Request.Context(), &req)
	if err != nil {
		log.Error("Failed to register user: ", err)
		helpers.SendResponse(c, 500, constant.ErrInternalServerError, nil)
		return
	}

	helpers.SendResponse(c, 200, constant.SuccessMessage, resp)
}

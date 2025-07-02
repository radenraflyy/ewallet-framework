package cmd

import (
	"net/http"
	"wallet/external"
	"wallet/helpers"

	"github.com/gin-gonic/gin"
)

func MiddlewareValidateToken(c *gin.Context) {
	var log = helpers.Logger

	tokenStr := c.Request.Header.Get("Authorization")

	if tokenStr == "" {
		log.Println("auth empty")
		helpers.SendResponse(c, http.StatusUnauthorized, "unathorized", nil)
		c.Abort()
		return
	}

	tokenData, err := external.ValidateToken(c, tokenStr)
	if err != nil {
		log.Error(err)
		helpers.SendResponse(c, http.StatusUnauthorized, "unauthorized", nil)
		c.Abort()
		return
	}

	c.Set("token", tokenData)
	c.Next()
}

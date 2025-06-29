package cmd

import (
	"ewallet-ums/helpers"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) AuthMiddleware(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		log.Println("authorization empty")
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unathorized", nil)
		ctx.Abort()
		return
	}

	parts := strings.Split(token, " ")
	token = parts[1]

	_, err := d.UserRepository.GetUserSessionByToken(ctx, token)
	if err != nil {
		log.Println("failed to get userSession in DB")
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unathorized", nil)
		ctx.Abort()
		return
	}

	claim, err := helpers.ValidateToken(ctx, token)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unathorized", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired")
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unathorized", nil)
		ctx.Abort()
		return
	}
	ctx.Set("userId", claim.UserID)
	ctx.Set("token", claim)
	ctx.Next()
}

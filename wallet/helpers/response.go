package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(ctx *gin.Context, code int, message string, data interface{}) {
	resp := &Response{
		Message: message,
		Data:    data,
	}
	ctx.JSON(code, resp)
}

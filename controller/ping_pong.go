package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingPongController struct {
}

func (p *PingPongController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func NewPingPongController() PingPongController {
	return PingPongController{}
}

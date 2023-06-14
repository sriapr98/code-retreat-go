package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingPongController struct {
}

func (p *PingPongController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func NewPingPongController() PingPongController {
	return PingPongController{}
}

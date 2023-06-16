package controller

import (
	testdata "code-retreat-go/testData"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingPongController struct {
}

func (p *PingPongController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (p *PingPongController) TestEmployeeResignations(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, testdata.GetTerminations())
}

func NewPingPongController() PingPongController {
	return PingPongController{}
}

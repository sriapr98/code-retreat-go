package controller

import (
	"code-retreat-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TerminationController struct {
	terminationService service.TerminationService
}

func (p *TerminationController) GetAllTerminations(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, p.terminationService.GetAllTerminations())
}

func NewTerminationController(terminationService service.TerminationService) TerminationController {
	return TerminationController{terminationService}
}

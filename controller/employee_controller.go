package controller

import (
	"code-retreat-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeService service.EmployeeService
}

func (p *EmployeeController) GetAllEmployees(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, p.employeeService.GetAllEmployees())
}

func NewEmployeeController(employeeService service.EmployeeService) EmployeeController {
	return EmployeeController{employeeService}
}

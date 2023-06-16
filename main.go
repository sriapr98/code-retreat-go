package main

import (
	"code-retreat-go/controller"
	"code-retreat-go/repository"
	"code-retreat-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	initializeRoutes(router)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func initializeRoutes(router *gin.Engine) {
	employeeRepository := repository.NewEmployeeRepository()
	err := employeeRepository.LoadSeedData()
	if err != nil {
		panic(err)
	}

	employeeService := service.NewEmployeeService(employeeRepository)

	pingPongController := controller.NewPingPongController()
	employeeController := controller.NewEmployeeController(employeeService)

	router.GET("/ping", pingPongController.Ping)
	router.GET("/test", pingPongController.TestEmployeeResignations)
	router.GET("/employees", employeeController.GetAllEmployees)
}

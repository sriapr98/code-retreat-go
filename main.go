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
	terminationRepository := repository.NewTerminationRepository()
	err = terminationRepository.LoadTerminationSeedData(employeeRepository)
	if err != nil {
		panic(err)
	}

	employeeService := service.NewEmployeeService(employeeRepository)
	terminationService := service.NewTerminationService(terminationRepository)

	pingPongController := controller.NewPingPongController()
	employeeController := controller.NewEmployeeController(employeeService)
	terminationController := controller.NewTerminationController(terminationService)

	router.GET("/ping", pingPongController.Ping)
	router.GET("/employees", employeeController.GetAllEmployees)
	router.GET("/terminations", terminationController.GetAllTerminations)
}

package main

import (
	"code-retreat-go/controller"
	"github.com/gin-gonic/gin"
	"net/http"
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
	pingPongController := controller.NewPingPongController()

	router.GET("/ping", pingPongController.Ping)
}

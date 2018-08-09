package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/L-oris/herokuDeploy/controller"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	controller := controller.NewController()

	router.GET("/", controller.PrintMessage)
	router.POST("/", controller.AddToMessage)
	router.DELETE("/", controller.ResetMessage)

	port := ":80"
	log.Println("listen on port", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Println("error listening to port", port, ":", err)
		return
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/L-oris/herokuDeploy/controller"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	controller := controller.NewController()

	router.GET("/", controller.PrintMessage)
	router.POST("/", controller.AddToMessage)
	router.DELETE("/", controller.ResetMessage)

	port := determineEnvPort()
	log.Println("listen on port", port)
	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Println("error listening to port", port, ":", err)
		return
	}
}

func determineEnvPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT not found in env variable, using default setting to :80")
		port = ":80"
	}

	return port
}

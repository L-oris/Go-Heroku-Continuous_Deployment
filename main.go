package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/L-oris/herokuDeploy/controller"
	"github.com/julienschmidt/httprouter"
)

const defaultPort = "8080"

func main() {
	router := httprouter.New()
	controller := controller.NewController()

	router.GET("/", controller.PrintMessage)
	router.POST("/", controller.AddToMessage)
	router.DELETE("/", controller.ResetMessage)

	port := determineEnvPort()
	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Println("error listening to port", port, ":", err)
		return
	}
}

func determineEnvPort() string {
	port := os.Getenv("PORT")
	log.Println("'PORT' received from env variables: ", port)

	if port == "" {
		log.Println("error 'PORT' not found in env variables, using default ", defaultPort)
		port = defaultPort
	}

	log.Println("listen on port:", port)
	return ":" + port
}

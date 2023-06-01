package main

import (
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/williamluisan/golang_pdfgenerator/config"
	routes "github.com/williamluisan/golang_pdfgenerator/routes"
)	

func init() {
	var config config.Config

	// initialize rabbitmq
	config.InitRabbitmq()
}

func main() {
	// defer config.RabbitmqChPubl.Close()
	// defer config.RabbitmqChCons.Close()

	// initialize gin
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":4747"))
}
package main

import (
	"commercial-propfloor/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8800"
	}

	router := gin.New()
	router.Use(gin.Logger())
	//	routes.PublicRoutes(router)
	routes.PrivateRoutes(router)

	log.Fatal(router.Run(":" + port))

}

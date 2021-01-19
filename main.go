package main

import (
	"api/src/core/config"
	"github.com/gin-gonic/gin"

	routes "api/src/social/routes"
	//"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	_ = godotenv.Load()
	config.Container().Get("app-provider")
}

func main() {
	r := gin.Default()
	apiGroup := r.Group("/api/v1/")
	routes.Init(apiGroup)

	err := r.Run(":" + serverPort())

	if err != nil {
		log.Println(err.Error())
	}

}

func serverPort() string {
	ginPort := os.Getenv("GIN_PORT")
	if ginPort == "" {
		ginPort = "8080"
	}
	return ginPort
}


package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load("../.env")

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	init_router(router)

	port := "8080"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	err := router.Run(":" + port)

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err.Error())
	}
}

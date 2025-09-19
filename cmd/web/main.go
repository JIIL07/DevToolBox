package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/JIIL07/devtoolbox/internal/api"
	"github.com/JIIL07/devtoolbox/internal/core"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	registry := core.NewGeneratorRegistry()
	handler := api.NewHandler(registry)

	router := gin.Default()
	
	router.Use(api.CORSMiddleware())
	router.Use(api.Logger())

	router.GET("/health", handler.Health)
	router.GET("/generators", handler.ListGenerators)
	router.POST("/generate", handler.Generate)

	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

package main

import (
	"log"
	"os"

	"github.com/JIIL07/devtoolbox/internal/api"
	"github.com/JIIL07/devtoolbox/internal/core"
	"github.com/JIIL07/devtoolbox/internal/plugins"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	registry := core.NewGeneratorRegistry()

	pluginManager := plugins.NewPluginManager()
	customPlugins, err := pluginManager.LoadPlugins()
	if err == nil {
		for _, pluginInfo := range customPlugins {
			if pluginInfo.Type == "python" {
				pythonPlugin := plugins.NewPythonPlugin(pluginInfo.Name, pluginInfo.Description, pluginInfo.Path)
				registry.Register(pythonPlugin)
			}
		}
	}

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

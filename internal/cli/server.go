package cli

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/JIIL07/devtoolbox/internal/api"
	"github.com/JIIL07/devtoolbox/internal/core"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the web server",
	Long: `Start the DevToolBox web server.

The server provides a web interface for code generation and exposes REST API endpoints.

Examples:
  devtoolbox server
  devtoolbox server --port 3000
  devtoolbox server --host 0.0.0.0 --port 8080`,
	Run: runServer,
}

var serverHost string
var serverPort string

func init() {
	serverCmd.Flags().StringVar(&serverHost, "host", "localhost", "Server host")
	serverCmd.Flags().StringVar(&serverPort, "port", "8080", "Server port")
}

func runServer(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)
	
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	
	registry := core.NewGeneratorRegistry()
	handler := api.NewHandler(registry)
	
	router.GET("/health", handler.Health)
	router.GET("/generators", handler.ListGenerators)
	router.POST("/generate", handler.Generate)
	
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		
		c.Next()
	})
	
	addr := fmt.Sprintf("%s:%s", serverHost, serverPort)
	
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	
	go func() {
		fmt.Printf("🚀 DevToolBox server starting on http://%s\n", addr)
		fmt.Printf("📋 Available endpoints:\n")
		fmt.Printf("   GET  /health     - Health check\n")
		fmt.Printf("   GET  /generators - List available generators\n")
		fmt.Printf("   POST /generate   - Generate code\n")
		fmt.Printf("\nPress Ctrl+C to stop the server\n")
		
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	fmt.Println("\n🛑 Shutting down server...")
}

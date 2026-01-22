package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/app/Http/Middleware"
	"github.com/y-as-7/go-askar/routes"
)

const Version = "1.4.0"

type Application struct {
	Router *gin.Engine
	Server *http.Server
}

func NewApplication() *Application {
	gin.SetMode(gin.ReleaseMode)
	
	router := gin.New()
	
	// Add middleware
	router.Use(Middleware.Logger())
	router.Use(gin.Recovery())
	router.Use(Middleware.CORS())
	
	// Load routes (Laravel-style separation)
	routes.RegisterWebRoutes(router)
	routes.RegisterAPIRoutes(router)
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	
	return &Application{
		Router: router,
		Server: server,
	}
}

func (app *Application) Run() {
	displayLogo()
	
	// Start server in goroutine
	go func() {
		fmt.Printf("\n✓ Server listening on http://localhost:8080\n\n")
		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("\n❌ Error: %s\n", err.Error())
		}
	}()
	
	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\n🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		fmt.Printf("❌ Server forced to shutdown: %v\n", err)
	}

	fmt.Println("✓ Server exited gracefully")
}

func displayLogo() {
	logo := `
   █████╗ ███████╗██╗  ██╗ █████╗ ██████╗ 
  ██╔══██╗██╔════╝██║ ██╔╝██╔══██╗██╔══██╗
  ███████║███████╗█████╔╝ ███████║██████╔╝
  ██╔══██║╚════██║██╔═██╗ ██╔══██║██╔══██╗
  ██║  ██║███████║██║  ██╗██║  ██║██║  ██║
  ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝

              A Laravel-inspired Go Framework
              Version ` + Version + ` ✨
`
	fmt.Println(logo)
}

package foundation

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/pkg/ui"
	"github.com/y-as-7/go-askar/routes"
)

// Application version
const Version = "1.3.3"

type Application struct {
	Router *gin.Engine
}

func NewApp() *Application {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)
	
	router := gin.Default()
	
	// Initialize routes
	routes.SetupRoutes(router)
	
	return &Application{
		Router: router,
	}
}

func (app *Application) Run() {
	ui.DisplayLogo(Version)
	time.Sleep(500 * time.Millisecond)

	ui.PrintStep("Starting askar application...")
	
	// Start server in a goroutine so we can handle signals
	go func() {
		if err := app.Router.Run(":8080"); err != nil {
			ui.PrintError("Failed to start server: " + err.Error())
		}
	}()

	fmt.Printf("\n  %s✓%s Server listening on :http://localhost:8080\n\n", ui.Green, ui.Reset)
	
	// Wait for interruption signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ui.PrintStep("Shutting down server...")
}

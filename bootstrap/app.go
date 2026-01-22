package bootstrap

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/routes"
)

const Version = "1.4.0"

type Application struct {
	Router *gin.Engine
}

func NewApplication() *Application {
	gin.SetMode(gin.ReleaseMode)
	
	router := gin.Default()
	
	// Load routes (Laravel-style separation)
	routes.RegisterWebRoutes(router)
	routes.RegisterAPIRoutes(router)
	
	return &Application{
		Router: router,
	}
}

func (app *Application) Run() {
	displayLogo()
	
	go func() {
		if err := app.Router.Run(":8080"); err != nil {
			fmt.Printf("\n❌ Error: %s\n", err.Error())
		}
	}()

	fmt.Printf("\n✓ Server listening on http://localhost:8080\n\n")
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\n🛑 Shutting down server...")
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

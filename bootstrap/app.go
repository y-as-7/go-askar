package bootstrap

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/app/Http/Middleware"
	"github.com/y-as-7/go-askar/config"
	"github.com/y-as-7/go-askar/routes"
)

const Version = "1.4.0"

type Application struct {
	Router *gin.Engine
	Server *http.Server
}

func NewApplication() *Application {
	// Initialize Database
	config.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)
	
	router := gin.New()
	
	// Add middleware
	router.Use(Middleware.Logger())
	router.Use(gin.Recovery())
	router.Use(Middleware.CORS())

	// Serve static files from the public directory
	router.Static("/public", "./public")

	// Load HTML templates recursively
	loadTemplates(router)
	
	// Load routes (Expressive separation)
	routes.RegisterWebRoutes(router)
	routes.RegisterAPIRoutes(router)
	
	app := &Application{
		Router: router,
	}

	app.Server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	
	return app
}

// VersionedAPI creates a new API version group
func (app *Application) VersionedAPI(version string, registerFunc func(*gin.RouterGroup)) {
	apiGroup := app.Router.Group("/api/" + version)
	registerFunc(apiGroup)
}

func loadTemplates(router *gin.Engine) {
	var templ *template.Template
	
	funcMap := template.FuncMap{
		"render": func(name string, data interface{}) (template.HTML, error) {
			var buf strings.Builder
			if err := templ.ExecuteTemplate(&buf, name, data); err != nil {
				return "", err
			}
			return template.HTML(buf.String()), nil
		},
	}

	templ = template.New("").Funcs(funcMap)
	
	// Load application views
	filepath.Walk("resources/views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})

	// Load library views (DashAskar)
	filepath.Walk("library/dash-askar/resources/views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	
	router.SetHTMLTemplate(templ)
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

              The Expressive Go Framework
              Version ` + Version + ` ✨
`
	fmt.Println(logo)
}

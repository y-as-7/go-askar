package console

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	Register(Command{
		Name:        "init",
		Description: "Initialize a new Askar project structure",
		Execute:     executeInit,
	})
}

func executeInit(args []string) error {
	fmt.Println("🚀 Initializing Askar project structure...")

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// 1. Create directory structure
	directories := []string{
		"app/Http/Controllers",
		"app/Http/Middleware",
		"app/Models",
		"app/DashAskar/Resources",
		"config",
		"routes",
		"storage/database",
		"storage/logs",
		"resources/views/layouts",
		"public/css",
		"public/js",
	}

	for _, dir := range directories {
		err := os.MkdirAll(filepath.Join(cwd, dir), 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
		fmt.Printf("✓ Created directory: %s\n", dir)
	}

	// 2. Create basic files if they don't exist
	pkgName := "test-askar-app" // Should ideally be fetched from go.mod
	if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
		if content, err := os.ReadFile("go.mod"); err == nil {
			lines := strings.Split(string(content), "\n")
			if len(lines) > 0 && strings.HasPrefix(lines[0], "module ") {
				pkgName = strings.TrimPrefix(lines[0], "module ")
			}
		}
	}

	files := map[string]string{
		".env": "APP_NAME=AskarApp\nAPP_ENV=local\nAPP_DEBUG=true\nAPP_URL=http://localhost:8080\n\nDB_DRIVER=sqlite\nDB_PATH=storage/database/app.db\n",
		"main.go": fmt.Sprintf(`package main

import (
	"%s/bootstrap"
)

func main() {
	app := bootstrap.NewApplication()
	app.Run()
}
`, pkgName),
		"routes/web.go": `package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "pages/home", gin.H{
			"title": "Welcome to Askar",
		})
	})
}
`,
		"routes/api.go": `package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
	}
}
`,
		"resources/views/pages/home.html": `<!DOCTYPE html>
<html>
<head>
    <title>{{ .title }}</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-gray-900 text-white flex items-center justify-center h-screen">
    <div class="text-center">
        <h1 class="text-6xl font-bold text-blue-500 mb-4">Askar</h1>
        <p class="text-xl text-gray-400">The Expressive Go Framework</p>
    </div>
</body>
</html>
`,
		"bootstrap/app.go": fmt.Sprintf(`package bootstrap

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/config"
	"github.com/y-as-7/go-askar/pkg/debugger"
	"%s/routes"
)

type Application struct {
	Router *gin.Engine
	Server *http.Server
}

func NewApplication() *Application {
	config.ConnectDatabase()
	router := gin.Default()
	router.Static("/public", "./public")
	
	// Add debugger middleware
	router.Use(debugger.Middleware())
	
	loadTemplates(router)
	routes.RegisterWebRoutes(router)
	routes.RegisterAPIRoutes(router)
	
	app := &Application{
		Router: router,
		Server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}
	return app
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
		"add": func(a, b int) int {
			return a + b
		},
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}
	templ = template.New("").Funcs(funcMap)
	
	// Load application views
	filepath.Walk("resources/views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			_, err = templ.ParseFiles(path)
		}
		return nil
	})

	// Load debugger views
	filepath.Walk("pkg/debugger/resources/views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			_, err = templ.ParseFiles(path)
		}
		return nil
	})

	router.SetHTMLTemplate(templ)
}

func (app *Application) Run() {
	go func() {
		fmt.Printf("\n✓ Server listening on http://localhost:8080\n\n")
		app.Server.ListenAndServe()
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
`, pkgName),
	}

	for path, content := range files {
		fullPath := filepath.Join(cwd, path)
		// Create subdirectories if needed
		os.MkdirAll(filepath.Dir(fullPath), 0755)
		
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			err = os.WriteFile(fullPath, []byte(content), 0644)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %v", path, err)
			}
			fmt.Printf("✓ Created file: %s\n", path)
		}
	}

	fmt.Println("\n✅ Askar project initialized successfully!")
	fmt.Println("\nNext steps:")
	fmt.Println("  1. Run 'go mod tidy'")
	fmt.Println("  2. Run 'go askar run' to start the server")
	
	return nil
}

package DashAskar

import (
	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/app/DashAskar/Resources/Users"
	"github.com/y-as-7/go-askar/pkg/dash-askar/core"
)

func PanelProvider() *core.DashAskar {
	panel := core.New()

	// Register Resources
	panel.AddResource(&Users.UserResource{})

	// Register Middlewares
	panel.Middleware(func(c *gin.Context) {
		// Example Custom Middleware: Set a custom header
		c.Header("X-DashAskar-Custom", "Powered by Askar")
		c.Next()
	})

	return panel
}

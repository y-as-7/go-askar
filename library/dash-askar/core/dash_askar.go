package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/app/Models"
	"github.com/y-as-7/go-askar/config"
)

type DashAskar struct {
	Prefix      string
	Resources   []Resource
	Middlewares []gin.HandlerFunc
	navItems    []NavItem
}

type NavItem struct {
	Label string
	Slug  string
	Icon  string
}

func New() *DashAskar {
	return &DashAskar{
		Prefix: "/admin",
	}
}

func (d *DashAskar) Middleware(middlewares ...gin.HandlerFunc) *DashAskar {
	d.Middlewares = append(d.Middlewares, middlewares...)
	return d
}

func (d *DashAskar) Register(router *gin.Engine) {
	admin := router.Group(d.Prefix)
	{
		admin.GET("/login", d.handleLogin)
		admin.POST("/login", d.handleDoLogin)
		admin.GET("/logout", d.handleLogout)

		// Protected Routes
		protected := admin.Group("/")
		protected.Use(d.AdminAuth())
		
		// Apply Custom Middlewares
		if len(d.Middlewares) > 0 {
			protected.Use(d.Middlewares...)
		}

		{
			protected.GET("/", d.handleDashboard)

			// Register Resource Routes
			for _, res := range d.Resources {
				slug := res.GetSlug()
				resourceGroup := protected.Group("/" + slug)
				{
					resourceGroup.GET("/", d.handleResourceIndex(res))
					resourceGroup.GET("/create", d.handleResourceCreate(res))
					resourceGroup.POST("/", d.handleResourceStore(res))
					resourceGroup.GET("/:id", d.handleResourceView(res))
					resourceGroup.GET("/:id/edit", d.handleResourceEdit(res))
					resourceGroup.PUT("/:id", d.handleResourceUpdate(res))
					resourceGroup.DELETE("/:id", d.handleResourceDelete(res))
				}

				d.navItems = append(d.navItems, NavItem{
					Label: res.GetTitle(),
					Slug:  slug,
					Icon:  res.GetIcon(),
				})
			}
		}
	}
}

func (d *DashAskar) AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("admin_token")
		if err != nil || token == "" {
			c.Redirect(http.StatusFound, d.Prefix+"/login")
			c.Abort()
			return
		}
		// In a real app, you'd validate the token/session here
		c.Next()
	}
}

func (d *DashAskar) handleLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "dash-askar/auth/login", gin.H{
		"title": "Admin Login",
	})
}

func (d *DashAskar) handleDoLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user Models.User
	if err := config.DB.Where("email = ? AND role = ?", email, "admin").First(&user).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "dash-askar/auth/login", gin.H{
			"error": "Invalid credentials or not an admin",
		})
		return
	}

	if err := user.CheckPassword(password); err != nil {
		c.HTML(http.StatusUnauthorized, "dash-askar/auth/login", gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// Set a simple cookie for demonstration
	c.SetCookie("admin_token", "dummy-admin-token", 3600, "/", "", false, true)
	c.Redirect(http.StatusFound, d.Prefix)
}

func (d *DashAskar) handleLogout(c *gin.Context) {
	c.SetCookie("admin_token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, d.Prefix+"/login")
}

func (d *DashAskar) AddResource(res Resource) {
	d.Resources = append(d.Resources, res)
}

func (d *DashAskar) handleResourceIndex(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "dash-askar/layouts/admin", gin.H{
			"title":       res.GetTitle(),
			"resource":    res,
			"nav":         d.navItems,
			"ContentName": "dash-askar/resource/index-content",
		})
	}
}

func (d *DashAskar) handleResourceView(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		infolist := &Infolist{}
		res.Infolist(infolist)

		c.HTML(200, "dash-askar/layouts/admin", gin.H{
			"title":       "View " + res.GetTitle(),
			"resource":    res,
			"id":          id,
			"schema":      infolist.Schema,
			"nav":         d.navItems,
			"ContentName": "dash-askar/resource/view-content",
		})
	}
}

func (d *DashAskar) handleResourceCreate(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		form := &Form{}
		res.Form(form)

		c.HTML(200, "dash-askar/layouts/admin", gin.H{
			"title":       "Create " + res.GetTitle(),
			"resource":    res,
			"schema":      form.Schema,
			"nav":         d.navItems,
			"ContentName": "dash-askar/resource/create-content",
		})
	}
}

func (d *DashAskar) handleResourceStore(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Store logic for " + res.GetTitle()})
	}
}

func (d *DashAskar) handleResourceEdit(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		form := &Form{}
		res.Form(form)

		c.HTML(200, "dash-askar/layouts/admin", gin.H{
			"title":       "Edit " + res.GetTitle(),
			"resource":    res,
			"schema":      form.Schema,
			"nav":         d.navItems,
			"ContentName": "dash-askar/resource/edit-content",
		})
	}
}

func (d *DashAskar) handleResourceUpdate(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Update logic for " + res.GetTitle()})
	}
}

func (d *DashAskar) handleResourceDelete(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Delete logic for " + res.GetTitle()})
	}
}

func (d *DashAskar) handleDashboard(c *gin.Context) {
	c.HTML(200, "dash-askar/layouts/admin", gin.H{
		"title":       "Admin Dashboard",
		"nav":         d.navItems,
		"ContentName": "dash-askar/dashboard-content",
	})
}

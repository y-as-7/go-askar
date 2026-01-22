package core

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/app/Models"
	"github.com/y-as-7/go-askar/config"
	"github.com/y-as-7/go-askar/pkg/debugger"
)

type DashAskar struct {
	Prefix      string
	Resources   []Resource
	Middlewares []gin.HandlerFunc
	Debug       bool
	navItems    []NavItem
}

type NavItem struct {
	Label string
	Slug  string
	Icon  string
}

type Notification struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Duration int   `json:"duration"`
}

type NotificationStore struct {
	Success []Notification
	Error   []Notification
	Warning []Notification
	Info    []Notification
}

func New() *DashAskar {
	return &DashAskar{
		Prefix: "/admin",
		Debug:  true, // Enabled by default for now
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
		protected.Use(d.MethodOverride())

		if d.Debug {
			protected.Use(debugger.Middleware())
		}

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
					resourceGroup.POST("/:id", d.handleResourceUpdate(res)) // For method override
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

func (d *DashAskar) MethodOverride() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			// Parse form data first
			c.Request.ParseForm()
			if method := c.Request.PostForm.Get("_method"); method != "" {
				c.Request.Method = method
			}
		}
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

func (d *DashAskar) SetNotification(c *gin.Context, notificationType, title, body string) {
	notifications := d.GetNotifications(c)

	notification := Notification{
		Type:     notificationType,
		Title:    title,
		Body:     body,
		Duration: 5000, // 5 seconds
	}

	switch notificationType {
	case "success":
		notifications.Success = append(notifications.Success, notification)
	case "error":
		notifications.Error = append(notifications.Error, notification)
	case "warning":
		notifications.Warning = append(notifications.Warning, notification)
	case "info":
		notifications.Info = append(notifications.Info, notification)
	}

	d.storeNotifications(c, notifications)
}

func (d *DashAskar) GetNotifications(c *gin.Context) *NotificationStore {
	// Try to get notifications from cookie
	if cookieValue, err := c.Cookie("notifications"); err == nil && cookieValue != "" {
		var store NotificationStore
		if err := json.Unmarshal([]byte(cookieValue), &store); err == nil {
			// Clear the cookie after reading
			c.SetCookie("notifications", "", -1, "/", "", false, false)
			return &store
		}
	}
	return &NotificationStore{}
}

func (d *DashAskar) storeNotifications(c *gin.Context, store *NotificationStore) {
	// Store notifications in cookie
	if data, err := json.Marshal(store); err == nil {
		c.SetCookie("notifications", string(data), 60, "/", "", false, false) // 1 minute expiry
	}
}

func (d *DashAskar) handleResourceIndex(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := debugger.GetDB(c)
		var records []map[string]interface{}
		db.Table(res.GetSlug()).Where("deleted_at IS NULL").Find(&records)

		table := &Table{}
		res.Table(table)

		c.HTML(200, "dash-askar/layouts/admin", gin.H{
			"title":         res.GetTitle(),
			"resource":      res,
			"records":       records,
			"table":         table,
			"nav":           d.navItems,
			"ContentName":   "dash-askar/resource/index-content",
			"notifications": d.GetNotifications(c),
			"AskarDebuggerData": func() interface{} {
				if dbgr := debugger.GetDebugger(c); dbgr != nil {
					return dbgr.Data
				}
				return nil
			}(),
		})
	}
}

func (d *DashAskar) handleResourceView(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := debugger.GetDB(c)
		id := c.Param("id")
		var record map[string]interface{}

		// Use Raw query instead of Table().First() to avoid GORM model requirements
		query := "SELECT * FROM " + res.GetSlug() + " WHERE id = ? AND (deleted_at IS NULL OR deleted_at = '')"
		result := db.Raw(query, id).Scan(&record)

		if result.Error != nil || result.RowsAffected == 0 {
			c.HTML(404, "dash-askar/layouts/admin", gin.H{
				"title":         "Record Not Found",
				"error":         "Record with ID " + id + " not found in table " + res.GetSlug(),
				"nav":           d.navItems,
				"ContentName":   "dash-askar/resource/view-content",
				"notifications": d.GetNotifications(c),
				"AskarDebuggerData": func() interface{} {
					if dbgr := debugger.GetDebugger(c); dbgr != nil {
						return dbgr.Data
					}
					return nil
				}(),
			})
			return
		}

		infolist := &Infolist{}
		res.Infolist(infolist)

		// Populate the schema with actual data values
		for i := range infolist.Schema {
			if value, exists := record[infolist.Schema[i].Name]; exists {
				infolist.Schema[i].Value = value
			}
		}

		c.HTML(200, "dash-askar/layouts/admin", gin.H{
			"title":         "View " + res.GetTitle(),
			"resource":      res,
			"id":            id,
			"record":        record,
			"schema":        infolist.Schema,
			"nav":           d.navItems,
			"ContentName":   "dash-askar/resource/view-content",
			"notifications": d.GetNotifications(c),
			"AskarDebuggerData": func() interface{} {
				if dbgr := debugger.GetDebugger(c); dbgr != nil {
					return dbgr.Data
				}
				return nil
			}(),
		})
	}
}

func (d *DashAskar) handleResourceCreate(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		form := &Form{}
		res.Form(form)

		c.HTML(200, "dash-askar/layouts/admin", gin.H{
			"title":         "Create " + res.GetTitle(),
			"resource":      res,
			"schema":        form.Schema,
			"nav":           d.navItems,
			"ContentName":   "dash-askar/resource/create-content",
			"notifications": d.GetNotifications(c),
			"AskarDebuggerData": func() interface{} {
				if dbgr := debugger.GetDebugger(c); dbgr != nil {
					return dbgr.Data
				}
				return nil
			}(),
		})
	}
}

func (d *DashAskar) handleResourceStore(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := debugger.GetDB(c)
		// Get form data instead of JSON
		createData := make(map[string]interface{})

		// Get form values based on the resource schema
		form := &Form{}
		res.Form(form)

		for _, field := range form.Schema {
			if value := c.PostForm(field.Name); value != "" {
				createData[field.Name] = value
			}
		}

		if len(createData) == 0 {
			d.SetNotification(c, "error", "Error", "No data provided")
			c.Redirect(http.StatusFound, d.Prefix+"/"+res.GetSlug()+"/create")
			return
		}

		result := db.Table(res.GetSlug()).Create(createData)
		if result.Error != nil {
			d.SetNotification(c, "error", "Error", "Failed to create record")
			c.Redirect(http.StatusFound, d.Prefix+"/"+res.GetSlug()+"/create")
			return
		}

		d.SetNotification(c, "success", "Success", res.GetTitle()+" created successfully")
		c.Redirect(http.StatusFound, d.Prefix+"/"+res.GetSlug())
	}
}

func (d *DashAskar) handleResourceEdit(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := debugger.GetDB(c)
		id := c.Param("id")

		// Fetch the existing record data
		var record map[string]interface{}
		query := "SELECT * FROM " + res.GetSlug() + " WHERE id = ? AND (deleted_at IS NULL OR deleted_at = '')"
		result := db.Raw(query, id).Scan(&record)

		if result.Error != nil || result.RowsAffected == 0 {
			c.HTML(404, "dash-askar/layouts/admin", gin.H{
				"title":         "Record Not Found",
				"error":         "Record with ID " + id + " not found",
				"nav":           d.navItems,
				"ContentName":   "dash-askar/resource/edit-content",
				"notifications": d.GetNotifications(c),
				"AskarDebuggerData": func() interface{} {
					if dbgr := debugger.GetDebugger(c); dbgr != nil {
						return dbgr.Data
					}
					return nil
				}(),
			})
			return
		}

		form := &Form{}
		res.Form(form)

		// Pre-populate form fields with existing data
		for i := range form.Schema {
			if value, exists := record[form.Schema[i].Name]; exists {
				form.Schema[i].Value = value
			}
		}

		c.HTML(200, "dash-askar/layouts/admin", gin.H{
			"title":         "Edit " + res.GetTitle(),
			"resource":      res,
			"id":            id,
			"record":        record,
			"schema":        form.Schema,
			"nav":           d.navItems,
			"ContentName":   "dash-askar/resource/edit-content",
			"notifications": d.GetNotifications(c),
			"AskarDebuggerData": func() interface{} {
				if dbgr := debugger.GetDebugger(c); dbgr != nil {
					return dbgr.Data
				}
				return nil
			}(),
		})
	}
}

func (d *DashAskar) handleResourceUpdate(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := debugger.GetDB(c)
		id := c.Param("id")

		// Get form data instead of JSON
		updateData := make(map[string]interface{})

		// Get form values based on the resource schema
		form := &Form{}
		res.Form(form)

		for _, field := range form.Schema {
			if value := c.PostForm(field.Name); value != "" {
				updateData[field.Name] = value
			}
		}

		if len(updateData) == 0 {
			d.SetNotification(c, "error", "Error", "No data provided")
			c.Redirect(http.StatusFound, d.Prefix+"/"+res.GetSlug()+"/"+id+"/edit")
			return
		}

		result := db.Table(res.GetSlug()).Where("id = ?", id).Updates(updateData)
		if result.Error != nil {
			d.SetNotification(c, "error", "Error", "Failed to update record")
			c.Redirect(http.StatusFound, d.Prefix+"/"+res.GetSlug()+"/"+id+"/edit")
			return
		}

		if result.RowsAffected == 0 {
			d.SetNotification(c, "error", "Error", "Record not found")
			c.Redirect(http.StatusFound, d.Prefix+"/"+res.GetSlug()+"/"+id+"/edit")
			return
		}

		d.SetNotification(c, "success", "Success", res.GetTitle()+" updated successfully")
		c.Redirect(http.StatusFound, d.Prefix+"/"+res.GetSlug())
	}
}

func (d *DashAskar) handleResourceDelete(res Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := debugger.GetDB(c)
		id := c.Param("id")

		result := db.Table(res.GetSlug()).Where("id = ?", id).Update("deleted_at", "NOW()")
		if result.Error != nil {
			d.SetNotification(c, "error", "Error", "Failed to delete record")
			c.Redirect(http.StatusFound, d.Prefix+"/"+res.GetSlug())
			return
		}

		if result.RowsAffected == 0 {
			d.SetNotification(c, "error", "Error", "Record not found")
			c.Redirect(http.StatusFound, d.Prefix+"/"+res.GetSlug())
			return
		}

		d.SetNotification(c, "success", "Success", res.GetTitle()+" deleted successfully")
		c.Redirect(http.StatusFound, d.Prefix+"/"+res.GetSlug())
	}
}

func (d *DashAskar) handleDashboard(c *gin.Context) {
	c.HTML(200, "dash-askar/layouts/admin", gin.H{
		"title":         "Admin Dashboard",
		"nav":           d.navItems,
		"ContentName":   "dash-askar/dashboard-content",
		"notifications": d.GetNotifications(c),
		"AskarDebuggerData": func() interface{} {
			if dbgr := debugger.GetDebugger(c); dbgr != nil {
				return dbgr.Data
			}
			return nil
		}(),
	})
}


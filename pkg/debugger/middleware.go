package debugger

import (
	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/config"
	"gorm.io/gorm"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize debugger for this request
		dbgr := New()
		c.Set("askardebugger", dbgr)

		// Create a session-specific DB instance with the debugger's logger
		if config.DB != nil {
			gormLogger := NewGormLogger(dbgr)
			sessionDB := config.DB.Session(&gorm.Session{
				Logger: gormLogger,
			})
			c.Set("db", sessionDB)
		}

		c.Next()

		// Finalize debugger data
		dbgr.Finish()
	}
}

func GetDebugger(c *gin.Context) *Debugger {
	if val, exists := c.Get("askardebugger"); exists {
		if dbgr, ok := val.(*Debugger); ok {
			return dbgr
		}
	}
	return nil
}

func GetDB(c *gin.Context) *gorm.DB {
	if val, exists := c.Get("db"); exists {
		if db, ok := val.(*gorm.DB); ok {
			return db
		}
	}
	return config.DB
}

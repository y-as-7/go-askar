package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/config"
	"github.com/y-as-7/go-askar/pkg/dash-askar/core"
)

type DeleteRecord struct {
	BasePage
	RecordID string
}

func NewDeleteRecord(resource core.Resource, navItems []core.NavItem) *DeleteRecord {
	return &DeleteRecord{
		BasePage: BasePage{
			Resource:   resource,
			Title:      "Delete " + resource.GetTitle(),
			Breadcrumb: "Delete " + resource.GetTitle(),
			NavItems:   navItems,
		},
	}
}

func (p *DeleteRecord) AuthorizeAccess(c *gin.Context) error {
	// Check if user can delete this record
	return nil
}

func (p *DeleteRecord) Handle(c *gin.Context) {
	p.RecordID = c.Param("id")

	if err := p.AuthorizeAccess(c); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Perform soft delete
	result := config.DB.Table(p.Resource.GetSlug()).Where("id = ?", p.RecordID).Update("deleted_at", "NOW()")
	if result.Error != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to delete record"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"success": false, "error": "Record not found"})
		return
	}

	c.JSON(200, gin.H{"success": true, "message": p.Resource.GetTitle() + " deleted successfully"})
}
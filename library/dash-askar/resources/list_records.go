package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/config"
	"github.com/y-as-7/go-askar/library/dash-askar/core"
)

type ListRecords struct {
	BasePage
	Records []map[string]interface{}
}

func NewListRecords(resource core.Resource, navItems []core.NavItem) *ListRecords {
	return &ListRecords{
		BasePage: BasePage{
			Resource:   resource,
			Title:      resource.GetTitle(),
			Breadcrumb: "List " + resource.GetTitle(),
			NavItems:   navItems,
		},
	}
}

func (p *ListRecords) AuthorizeAccess(c *gin.Context) error {
	// Check if user can list records
	return nil
}

func (p *ListRecords) Mount(c *gin.Context) error {
	if err := p.AuthorizeAccess(c); err != nil {
		return err
	}

	return p.loadRecords()
}

func (p *ListRecords) loadRecords() error {
	var records []map[string]interface{}
	result := config.DB.Table(p.Resource.GetSlug()).Where("deleted_at IS NULL").Find(&records)
	if result.Error != nil {
		return result.Error
	}

	p.Records = records
	return nil
}

func (p *ListRecords) GetTable() *Table {
	table := &Table{}
	p.Resource.Table(table)
	return table
}

func (p *ListRecords) Handle(c *gin.Context) {
	if err := p.Mount(c); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	table := p.GetTable()

	c.HTML(200, "dash-askar/layouts/admin", gin.H{
		"title":       p.GetTitle(),
		"resource":    p.Resource,
		"records":     p.Records,
		"table":       table,
		"nav":         p.NavItems,
		"ContentName": "dash-askar/resource/index-content",
	})
}

func (p *ListRecords) GetPageClasses() []string {
	return append(p.BasePage.GetPageClasses(),
		"fi-resource-list-records-page",
	)
}
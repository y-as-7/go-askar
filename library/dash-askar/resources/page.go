package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/library/dash-askar/core"
	"github.com/y-as-7/go-askar/library/dash-askar/schemas"
)

type Page interface {
	GetTitle() string
	GetBreadcrumb() string
	AuthorizeAccess(c *gin.Context) error
	Mount(c *gin.Context) error
	Handle(c *gin.Context)
	GetResource() core.Resource
	GetPageClasses() []string
}

type BasePage struct {
	Resource    core.Resource
	Title       string
	Breadcrumb  string
	NavItems    []core.NavItem
}

func (p *BasePage) GetTitle() string {
	return p.Title
}

func (p *BasePage) GetBreadcrumb() string {
	return p.Breadcrumb
}

func (p *BasePage) GetResource() core.Resource {
	return p.Resource
}

func (p *BasePage) GetPageClasses() []string {
	return []string{
		"fi-resource-page",
		"fi-resource-" + p.Resource.GetSlug(),
	}
}

func (p *BasePage) AuthorizeAccess(c *gin.Context) error {
	// Default implementation - override in specific pages
	return nil
}

func (p *BasePage) Mount(c *gin.Context) error {
	// Default implementation - override in specific pages
	return nil
}

type Schema = schemas.Schema
type Component = schemas.Component
type Form = schemas.Form
type Table = schemas.Table
type Infolist = schemas.Infolist
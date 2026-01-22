package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/config"
	"github.com/y-as-7/go-askar/pkg/dash-askar/core"
)

type CreateRecord struct {
	BasePage
	Data         map[string]interface{}
	IsCreating   bool
	PreviousUrl  string
}

func NewCreateRecord(resource core.Resource, navItems []core.NavItem) *CreateRecord {
	return &CreateRecord{
		BasePage: BasePage{
			Resource:   resource,
			Title:      "Create " + resource.GetTitle(),
			Breadcrumb: "Create " + resource.GetTitle(),
			NavItems:   navItems,
		},
		Data: make(map[string]interface{}),
	}
}

func (p *CreateRecord) AuthorizeAccess(c *gin.Context) error {
	// Check if user can create records
	return nil
}

func (p *CreateRecord) Mount(c *gin.Context) error {
	if err := p.AuthorizeAccess(c); err != nil {
		return err
	}

	p.FillForm()
	p.PreviousUrl = c.GetHeader("Referer")
	return nil
}

func (p *CreateRecord) FillForm() {
	// Initialize form with empty data
	p.Data = make(map[string]interface{})
}

func (p *CreateRecord) GetForm() *Form {
	form := &Form{}
	p.Resource.Form(form)
	return form
}

func (p *CreateRecord) Create(c *gin.Context, another bool) {
	if p.IsCreating {
		return
	}

	p.IsCreating = true
	defer func() { p.IsCreating = false }()

	if err := p.AuthorizeAccess(c); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Get form data
	createData := make(map[string]interface{})
	form := p.GetForm()

	for _, field := range form.Schema {
		if value := c.PostForm(field.Name); value != "" {
			createData[field.Name] = value
		}
	}

	if len(createData) == 0 {
		c.JSON(400, gin.H{"success": false, "error": "No data provided"})
		return
	}

	// Mutate form data before create
	createData = p.MutateFormDataBeforeCreate(createData)

	// Handle record creation
	if err := p.HandleRecordCreation(createData); err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to create record: " + err.Error()})
		return
	}

	if another {
		// Reset form for creating another record
		p.FillForm()
		c.JSON(200, gin.H{"success": true, "message": p.Resource.GetTitle() + " created successfully", "another": true})
		return
	}

	// Redirect to resource index or view
	redirectUrl := p.GetRedirectUrl()
	c.JSON(200, gin.H{"success": true, "message": p.Resource.GetTitle() + " created successfully", "redirect": redirectUrl})
}

func (p *CreateRecord) MutateFormDataBeforeCreate(data map[string]interface{}) map[string]interface{} {
	return data
}

func (p *CreateRecord) HandleRecordCreation(data map[string]interface{}) error {
	result := config.DB.Table(p.Resource.GetSlug()).Create(data)
	return result.Error
}

func (p *CreateRecord) GetRedirectUrl() string {
	return "/admin/" + p.Resource.GetSlug()
}

func (p *CreateRecord) Handle(c *gin.Context) {
	if c.Request.Method == "POST" {
		another := c.PostForm("create_another") == "1"
		p.Create(c, another)
		return
	}

	if err := p.Mount(c); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	form := p.GetForm()

	c.HTML(200, "dash-askar/layouts/admin", gin.H{
		"title":       p.GetTitle(),
		"resource":    p.Resource,
		"schema":      form.Schema,
		"nav":         p.NavItems,
		"ContentName": "dash-askar/resource/create-content",
	})
}

func (p *CreateRecord) GetPageClasses() []string {
	return append(p.BasePage.GetPageClasses(),
		"fi-resource-create-record-page",
	)
}
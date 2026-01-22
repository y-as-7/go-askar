package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/config"
	"github.com/y-as-7/go-askar/pkg/dash-askar/core"
)

type EditRecord struct {
	BasePage
	Record      map[string]interface{}
	Data        map[string]interface{}
	PreviousUrl string
	RecordID    string
}

func NewEditRecord(resource core.Resource, navItems []core.NavItem) *EditRecord {
	return &EditRecord{
		BasePage: BasePage{
			Resource:   resource,
			Title:      "Edit " + resource.GetTitle(),
			Breadcrumb: "Edit " + resource.GetTitle(),
			NavItems:   navItems,
		},
		Data: make(map[string]interface{}),
	}
}

func (p *EditRecord) AuthorizeAccess(c *gin.Context) error {
	// Check if user can edit this record
	return nil
}

func (p *EditRecord) Mount(c *gin.Context) error {
	p.RecordID = c.Param("id")

	if err := p.ResolveRecord(); err != nil {
		return err
	}

	if err := p.AuthorizeAccess(c); err != nil {
		return err
	}

	p.FillForm()
	p.PreviousUrl = c.GetHeader("Referer")
	return nil
}

func (p *EditRecord) ResolveRecord() error {
	var record map[string]interface{}
	query := "SELECT * FROM " + p.Resource.GetSlug() + " WHERE id = ? AND (deleted_at IS NULL OR deleted_at = '')"
	result := config.DB.Raw(query, p.RecordID).Scan(&record)

	if result.Error != nil || result.RowsAffected == 0 {
		return &RecordNotFoundError{ID: p.RecordID}
	}

	p.Record = record
	return nil
}

func (p *EditRecord) FillForm() {
	// Fill form with existing record data
	data := make(map[string]interface{})

	// Copy record data, applying any mutations
	for key, value := range p.Record {
		data[key] = value
	}

	p.Data = p.MutateFormDataBeforeFill(data)
}

func (p *EditRecord) MutateFormDataBeforeFill(data map[string]interface{}) map[string]interface{} {
	return data
}

func (p *EditRecord) GetForm() *Form {
	form := &Form{}
	p.Resource.Form(form)

	// Pre-populate form fields with existing data
	for i := range form.Schema {
		if value, exists := p.Data[form.Schema[i].Name]; exists {
			form.Schema[i].Value = value
		}
	}

	return form
}

func (p *EditRecord) Save(c *gin.Context) {
	if err := p.AuthorizeAccess(c); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Get form data
	updateData := make(map[string]interface{})
	form := &Form{}
	p.Resource.Form(form)

	for _, field := range form.Schema {
		if value := c.PostForm(field.Name); value != "" {
			updateData[field.Name] = value
		}
	}

	if len(updateData) == 0 {
		c.JSON(400, gin.H{"success": false, "error": "No data provided"})
		return
	}

	// Mutate form data before save
	updateData = p.MutateFormDataBeforeSave(updateData)

	// Handle record update
	if err := p.HandleRecordUpdate(updateData); err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to update record: " + err.Error()})
		return
	}

	// Refresh record data
	p.ResolveRecord()
	p.FillForm()

	redirectUrl := p.GetRedirectUrl()
	c.JSON(200, gin.H{"success": true, "message": p.Resource.GetTitle() + " updated successfully", "redirect": redirectUrl})
}

func (p *EditRecord) MutateFormDataBeforeSave(data map[string]interface{}) map[string]interface{} {
	return data
}

func (p *EditRecord) HandleRecordUpdate(data map[string]interface{}) error {
	result := config.DB.Table(p.Resource.GetSlug()).Where("id = ?", p.RecordID).Updates(data)
	return result.Error
}

func (p *EditRecord) GetRedirectUrl() string {
	return "/admin/" + p.Resource.GetSlug()
}

func (p *EditRecord) Handle(c *gin.Context) {
	if c.Request.Method == "PUT" || (c.Request.Method == "POST" && c.PostForm("_method") == "PUT") {
		p.Save(c)
		return
	}

	if err := p.Mount(c); err != nil {
		if _, ok := err.(*RecordNotFoundError); ok {
			c.HTML(404, "dash-askar/layouts/admin", gin.H{
				"title": "Record Not Found",
				"error": "Record with ID " + p.RecordID + " not found",
				"nav":   p.NavItems,
				"ContentName": "dash-askar/resource/edit-content",
			})
			return
		}
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	form := p.GetForm()

	c.HTML(200, "dash-askar/layouts/admin", gin.H{
		"title":       p.GetTitle(),
		"resource":    p.Resource,
		"id":          p.RecordID,
		"record":      p.Record,
		"schema":      form.Schema,
		"nav":         p.NavItems,
		"ContentName": "dash-askar/resource/edit-content",
	})
}

func (p *EditRecord) GetPageClasses() []string {
	return append(p.BasePage.GetPageClasses(),
		"fi-resource-edit-record-page",
		"fi-resource-record-"+p.RecordID,
	)
}

type RecordNotFoundError struct {
	ID string
}

func (e *RecordNotFoundError) Error() string {
	return "Record with ID " + e.ID + " not found"
}
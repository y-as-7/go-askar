package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/config"
	"github.com/y-as-7/go-askar/pkg/dash-askar/core"
)

type ViewRecord struct {
	BasePage
	Record   map[string]interface{}
	Data     map[string]interface{}
	RecordID string
}

func NewViewRecord(resource core.Resource, navItems []core.NavItem) *ViewRecord {
	return &ViewRecord{
		BasePage: BasePage{
			Resource:   resource,
			Title:      "View " + resource.GetTitle(),
			Breadcrumb: "View " + resource.GetTitle(),
			NavItems:   navItems,
		},
		Data: make(map[string]interface{}),
	}
}

func (p *ViewRecord) AuthorizeAccess(c *gin.Context) error {
	// Check if user can view this record
	return nil
}

func (p *ViewRecord) Mount(c *gin.Context) error {
	p.RecordID = c.Param("id")

	if err := p.ResolveRecord(); err != nil {
		return err
	}

	if err := p.AuthorizeAccess(c); err != nil {
		return err
	}

	if !p.HasInfolist() {
		p.FillForm()
	}

	return nil
}

func (p *ViewRecord) ResolveRecord() error {
	var record map[string]interface{}
	query := "SELECT * FROM " + p.Resource.GetSlug() + " WHERE id = ? AND (deleted_at IS NULL OR deleted_at = '')"
	result := config.DB.Raw(query, p.RecordID).Scan(&record)

	if result.Error != nil || result.RowsAffected == 0 {
		return &RecordNotFoundError{ID: p.RecordID}
	}

	p.Record = record
	return nil
}

func (p *ViewRecord) HasInfolist() bool {
	infolist := &Infolist{}
	p.Resource.Infolist(infolist)
	return len(infolist.Schema) > 0
}

func (p *ViewRecord) FillForm() {
	// Fill form with record data for viewing
	data := make(map[string]interface{})

	for key, value := range p.Record {
		data[key] = value
	}

	p.Data = p.MutateFormDataBeforeFill(data)
}

func (p *ViewRecord) MutateFormDataBeforeFill(data map[string]interface{}) map[string]interface{} {
	return data
}

func (p *ViewRecord) GetForm() *Form {
	form := &Form{}
	p.Resource.Form(form)

	// Populate form fields with record data (read-only)
	for i := range form.Schema {
		if value, exists := p.Data[form.Schema[i].Name]; exists {
			form.Schema[i].Value = value
		}
	}

	return form
}

func (p *ViewRecord) GetInfolist() *Infolist {
	infolist := &Infolist{}
	p.Resource.Infolist(infolist)

	// Populate infolist entries with actual data values
	for i := range infolist.Schema {
		if value, exists := p.Record[infolist.Schema[i].Name]; exists {
			infolist.Schema[i].Value = value
		}
	}

	return infolist
}

func (p *ViewRecord) GetRecordTitle() string {
	// Try to get a meaningful title from the record
	if name, ok := p.Record["name"].(string); ok && name != "" {
		return name
	}
	if title, ok := p.Record["title"].(string); ok && title != "" {
		return title
	}
	if id, ok := p.Record["id"]; ok {
		return "#" + string(rune(id.(int64)))
	}
	return "Record"
}

func (p *ViewRecord) Handle(c *gin.Context) {
	if err := p.Mount(c); err != nil {
		if _, ok := err.(*RecordNotFoundError); ok {
			c.HTML(404, "dash-askar/layouts/admin", gin.H{
				"title": "Record Not Found",
				"error": "Record with ID " + p.RecordID + " not found",
				"nav":   p.NavItems,
				"ContentName": "dash-askar/resource/view-content",
			})
			return
		}
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	var schema interface{}
	if p.HasInfolist() {
		infolist := p.GetInfolist()
		schema = infolist.Schema
	} else {
		form := p.GetForm()
		schema = form.Schema
	}

	c.HTML(200, "dash-askar/layouts/admin", gin.H{
		"title":       p.GetTitle(),
		"resource":    p.Resource,
		"id":          p.RecordID,
		"record":      p.Record,
		"schema":      schema,
		"nav":         p.NavItems,
		"ContentName": "dash-askar/resource/view-content",
	})
}

func (p *ViewRecord) GetPageClasses() []string {
	return append(p.BasePage.GetPageClasses(),
		"fi-resource-view-record-page",
		"fi-resource-record-"+p.RecordID,
	)
}